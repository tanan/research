package com.tanan.researchserver.domain.article

import com.tanan.researchserver.domain.FCC

data class ArticlesOverview(override val list: List<ArticleOverview>): FCC<ArticleOverview>

data class ArticleOverview(
        val name: Name,
        val category: Category,
        val thumbnail: Thumbnail,
        val url: Url
)

data class Name(val value: String)

data class Category(val value: String)

data class Thumbnail(val value: String)

data class Url(val value: String)