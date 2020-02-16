package com.tanan.researchserver.driver.article.database.entity

import javax.persistence.Entity
import javax.persistence.Id
import javax.persistence.Table

@Entity
@Table(name = "article_overview", schema = "research")
data class ArticleOverview(
        @Id
        val id: String,
        val name: String,
        val description: String,
        val category: String,
        val url: String
)
