package com.tanan.researchserver.driver.article.database

import com.tanan.researchserver.driver.article.database.entity.ArticleOverview
import org.springframework.data.jpa.repository.JpaRepository

interface ArticleOverviewDriver : JpaRepository<ArticleOverview, String> {}