package com.tanan.researchserver.port.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.Article
import com.tanan.researchserver.domain.article.ArticleOverview
import com.tanan.researchserver.domain.article.ArticlesOverview

interface ArticlePort {
    fun getArticleOverview(id: Id): ArticleOverview
    fun getArticle(id: Id): Article
    fun getLatestArticlesOverview(size: Int): ArticlesOverview
}