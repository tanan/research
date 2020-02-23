package com.tanan.researchserver.rest.jsons.article

import com.tanan.researchserver.domain.article.*

fun ArticlesOverview.toJson() = map { it.toJson() }.let(::ArticlesOverviewJson)

fun ArticleOverview.toJson() = ArticleOverviewJson(
        name.value,
        category.value,
        thumbnail.value,
        url.value)

data class ArticlesOverviewJson(
        val articlesOverviewJson: List<ArticleOverviewJson>
)

data class ArticleOverviewJson(
        val name: String,
        val category: String,
        val thumbnail: String,
        val url: String)
