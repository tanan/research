package com.tanan.researchserver.usecase.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.*
import com.tanan.researchserver.port.article.ArticlePort
import org.springframework.stereotype.Component

@Component
class ArticleUseCase(private val articlePort: ArticlePort) {
    fun getArticleOverview(id: Id): ArticleOverview {
        return articlePort.getArticleOverview(id)
    }

    fun getArticle(id: Id): String = articlePort.getArticle(id)

    fun getLatestArticles(size: Int): String = articlePort.getLatestArticles(size)
}