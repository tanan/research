package contentful

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"research-api/domain"
)

type ContentfulApi struct {
	URL        *url.URL
	HTTPClient *http.Client
	Header     http.Header
}

var endpoint = "https://cdn.contentful.com"

func NewContentfulApi(spaceId string) (*ContentfulApi, error) {
	if len(spaceId) == 0 {
		return nil, errors.New("invalid URL")
	}
	parsedUrl, err := url.ParseRequestURI(fmt.Sprintf("%s/spaces/%s", endpoint, spaceId))
	if err != nil {
		return nil, err
	}
	client := &ContentfulApi{
		URL:        parsedUrl,
		HTTPClient: &http.Client{},
		Header:     make(http.Header),
	}
	client.Header.Set("User-Agent", "golang")
	client.Header.Set("Content-Type", "application/json; charset=utf-8")
	client.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "9achdtJvXAIFUQqkzFcvN0_7qKv4pTxPel7YxtTgayE"))
	return client, nil
}

func decodeBody(body io.Reader, out interface{}) error {
	if out == nil {
		return nil
	}
	decoder := json.NewDecoder(body)
	return decoder.Decode(out)
}

func (c ContentfulApi) decodeResponse(code int, body io.ReadCloser, model interface{}) error {
	if code != http.StatusOK && code != http.StatusCreated {
		return fmt.Errorf("status_code: %s", code)
	}
	return decodeBody(body, model)
}

func (c *ContentfulApi) rawRequest(method, subPath string, params map[string]string) (req *http.Request, err error) {
	if method != "GET" {
		return nil, errors.New("missing requestMethod")
	}
	u := *c.URL
	u.Path = path.Join(c.URL.Path, subPath)
	var p = make(url.Values)
	for k, v := range params {
		p.Add(k, v)
	}
	u.RawQuery = p.Encode()
	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header = c.Header
	return request, err
}

func (c ContentfulApi) FindById(id domain.ArticleId) (fields map[string]interface{}, includes map[string]interface{}, err error) {
	req, err := c.rawRequest("GET", "/entries", map[string]string{
		"includes": "2",
		"sys.id":   string(id),
	})
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, nil
	}
	defer resp.Body.Close()
	var m map[string]interface{}
	err = c.decodeResponse(resp.StatusCode, resp.Body, &m)
	return c.toContentMap(m), c.toIncludesMap(m), err
}

func (c ContentfulApi) FindThumbnailUrl(id domain.ThumbnailId) (domain.ThumbnailUrl, error) {
	req, err := c.rawRequest("GET", fmt.Sprintf("/environments/master/assets/%s", id), map[string]string{})
	if err != nil {
		return "", err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var m map[string]interface{}
	err = c.decodeResponse(resp.StatusCode, resp.Body, &m)
	return c.toThumbnailUrl(m), err
}

func (c ContentfulApi) toContentMap(res map[string]interface{}) map[string]interface{} {
	if res == nil {
		return nil
	}
	fields := res["items"].([]interface{})[0].(map[string]interface{})["fields"].(map[string]interface{})
	return fields["content"].(map[string]interface{})
}

func (c ContentfulApi) toIncludesMap(res map[string]interface{}) map[string]interface{} {
	if res == nil {
		return nil
	}
	return res["includes"].(map[string]interface{})
}

func (c ContentfulApi) toThumbnailUrl(res map[string]interface{}) domain.ThumbnailUrl {
	if res == nil {
		return ""
	}
	fields := res["fields"].(map[string]interface{})
	file := fields["file"].(map[string]interface{})
	return domain.ThumbnailUrl(file["url"].(string))
}
