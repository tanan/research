package com.tanan.researchserver.domain.article

data class ArticleOverview(
        val name: Name,
        val description: Description,
        val category: Category,
        val url: Url
)

data class Name(val name: String)

data class Description(val description: String)

data class Category(val category: String)

data class Url(val url: String)