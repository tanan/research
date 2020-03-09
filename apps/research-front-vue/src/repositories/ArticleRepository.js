import Repository from "./Repository"

const resource = "/articles"

export default {
  get() {
    return Repository.get(`${resource}`)
  },
  getArticle(articleId) {
    return Repository.get(`${resource}/${articleId}/content`)
  },
  getArticlesOverview() {
    return Repository.get(`${resource}`+"?size=10")
  }
}
