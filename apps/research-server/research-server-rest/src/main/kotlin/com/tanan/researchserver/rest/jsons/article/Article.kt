package com.tanan.researchserver.rest.jsons.article

import com.tanan.researchserver.domain.article.Article
import com.tanan.researchserver.domain.article.LatestArticles

fun Article.toJson() = ArticleJson(
        sys.value,
        fields.toString())

data class ArticleJson(
        val sys: String,
        val fields: String)

fun LatestArticles.toJson() = LatestArticlesJson(
        data
)

data class LatestArticlesJson(
        val data: String
)