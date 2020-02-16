package com.tanan.researchserver.domain.article

data class Article(
        val sys: Sys,
        val fields: Fields
)

data class Sys(val value: String)

data class Fields(
        val title: String,
        val content: String
)
