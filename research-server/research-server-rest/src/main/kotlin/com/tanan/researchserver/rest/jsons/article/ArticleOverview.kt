package com.tanan.researchserver.rest.jsons.article

import com.tanan.researchserver.domain.article.ArticleOverview

fun ArticleOverview.toJson() = ArticleOverviewJson(
        name.value,
        description.value,
        category.value,
        url.value)

data class ArticleOverviewJson(
        val name: String,
        val description: String,
        val category: String,
        val url: String)
