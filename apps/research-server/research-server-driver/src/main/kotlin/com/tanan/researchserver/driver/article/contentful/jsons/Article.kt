package com.tanan.researchserver.driver.article.contentful.jsons

data class ArticleJson(
        val sys: SysJson,
        val fields: FieldsJson
)

data class SysJson(val sys: String)

data class FieldsJson(
        val title: String,
        val content: String
)
