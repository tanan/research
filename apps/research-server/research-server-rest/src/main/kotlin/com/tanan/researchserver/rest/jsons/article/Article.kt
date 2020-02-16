package com.tanan.researchserver.rest.jsons.article

import com.tanan.researchserver.domain.article.Article

fun Article.toJson() = ArticleJson(
        sys.value,
        fields.toString())

data class ArticleJson(
        val sys: String,
        val fields: String)
