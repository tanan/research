import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart';
import 'package:http/browser_client.dart';

import 'package:research_front/src/driver/api/error_handler.dart';
import 'package:research_front/src/driver/api/research_api_url.dart';
import 'package:research_front/src/driver/http/extensions.dart';
import 'json/article.dart';

Map<String, String> mergeMaps(List<Map<String, String>> src) {
  return src.fold({}, (p, v) => p..addAll(v));
}

class ApiClient {

  final BrowserClient _client;
  final HttpErrorHandler _handler;
  final ResearchApiURL _url;

  ApiClient(this._client, this._handler, this._url);

  Map<String, String> get _headers => mergeMaps([CONTENT_TYPE_JSON_UTF8]);

  Future<ArticleJson> getArticles() async =>
    _getAsMap(_url.latest(), (v) => ArticleJson.fromJson(v));

  Future<ArticleJson> getArticle(String id) async =>
    _getAsMap(_url.article(id), (v) => ArticleJson.fromJson(v));

  Future<T> _getAsMap<T>(String url, Converter<T> c) async {
    var response = await _client.get(url, headers: _headers);
    _handleResponse(response, _handler);
    return response.asMap(c);
  }

}

void _handleResponse(Response response, HttpErrorHandler handler) {

  if (response.statusCode == 401) {
    handler.notify(HttpError(response.statusCode));
    return;
  }

  if (response.statusCode >= 400) {
    throw HttpException(response.statusCode, response.body, response.request);
  }
}

class HttpException implements Exception {
  final int statucCode;
  final String message;
  final Request request;

  HttpException(this.statucCode, this.message, this.request);
}