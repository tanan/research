package com.tanan.researchserver.gateway.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.*
import com.tanan.researchserver.driver.article.contentful.ContentfulApi
import com.tanan.researchserver.driver.article.contentful.jsons.ArticleJson
import com.tanan.researchserver.driver.article.database.ArticleOverviewDriver
import com.tanan.researchserver.port.article.ArticlePort
import org.springframework.stereotype.Component

@Component
class ArticleGateway(private val articleOverviewDriver: ArticleOverviewDriver,
                     private val contentfulApi: ContentfulApi): ArticlePort {

    override fun getArticleOverview(id: Id): ArticleOverview =
            toArticleOverview(articleOverviewDriver.findById(id.value).orElse(
                    com.tanan.researchserver.driver.article.database.entity.ArticleOverview("", "", "", "", "")
            ))

    private fun toArticleOverview(entity: com.tanan.researchserver.driver.article.database.entity.ArticleOverview) =
            ArticleOverview(Name(entity.name), Description(entity.description), Category(entity.category), Url(entity.url))

    override fun getArticle(id: Id): Article =
        toArticle(contentfulApi.getContent(id.toString()))

    private fun toArticle(articleJson: ArticleJson) =
            Article(
                    Sys(value = articleJson.sys.sys),
                    Fields(title = articleJson.fields.title, content = articleJson.fields.content)
            )
}