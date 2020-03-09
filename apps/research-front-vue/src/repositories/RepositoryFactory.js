import ArticleRepository from "./ArticleRepository"

const repositories = {
  articles: ArticleRepository
}

export const RepositoryFactory = {
  get: name => repositories[name]
}