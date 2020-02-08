package com.tanan.researchserver.domain.article

data class ArticleOverview(
        val name: Name,
        val description: Description,
        val category: Category,
        val url: Url
)

data class Name(val value: String)

data class Description(val value: String)

data class Category(val value: String)

data class Url(val value: String)