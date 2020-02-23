package com.tanan.researchserver.gateway.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.*
import com.tanan.researchserver.driver.article.contentful.ContentfulApi
import com.tanan.researchserver.driver.article.contentful.jsons.ArticleJson
import com.tanan.researchserver.driver.article.contentful.jsons.ArticlesOverviewJson
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
            ArticleOverview(Name(entity.name), Category(entity.category), Thumbnail(""), Url(entity.url))

    override fun getArticle(id: Id): Article =
        toArticle(contentfulApi.getContent(id.toString()))

    override fun getLatestArticlesOverview(size: Int): ArticlesOverview =
        toArticlesOverview(contentfulApi.getArticlesOverview(size))

    private fun toArticlesOverview(articlesOverviewJson: ArticlesOverviewJson) =
        articlesOverviewJson.items
                .map {
                    Pair(it.fields.thumbnail.sys.id, it.fields.title)
                }.map {
                    val id = it.first
                    val title = it.second
                    articlesOverviewJson.includes
                            .map { j ->
                                val t = j.Asset.first { k -> k.sys.id === id }
                                ArticleOverview(
                                        name = Name(title),
                                        category = Category(""),
                                        thumbnail = Thumbnail(t.fields.file.url),
                                        url = Url("")
                                )
                            }
                }[0].let(::ArticlesOverview)

    private fun toArticle(articleJson: ArticleJson) =
            Article(
                    Sys(value = articleJson.sys.sys),
                    Fields(title = articleJson.fields.title, content = articleJson.fields.content)
            )

//    private fun List<ArticlesOverviewJson>.Thumbnail(id: String) = find { it.includes. { it.asset.find { it.sys.id === id } } }
}