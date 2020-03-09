package com.tanan.researchserver.port.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.Article
import com.tanan.researchserver.domain.article.ArticleOverview

interface ArticlePort {
    fun getArticleOverview(id: Id): ArticleOverview
    fun getArticle(id: Id): String
    fun getLatestArticles(size: Int): String
}