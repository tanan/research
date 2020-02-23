package com.tanan.researchserver.driver.article.contentful

import com.tanan.researchserver.driver.article.contentful.jsons.ArticleJson
import com.tanan.researchserver.driver.article.contentful.jsons.ArticlesOverviewJson
import org.springframework.boot.context.properties.ConfigurationProperties
import org.springframework.stereotype.Component
import org.springframework.web.client.RestTemplate
import java.util.*
import kotlin.collections.HashMap

@Component
class ContentfulApi(private val config: ContentfulApiConfig) {
    private val restTemplate = RestTemplate()
    private val endpoint = "${config.endPoint}/spaces/${config.spaceId}"

    fun getContent(articleId: String) =
        restTemplate.getForEntity(
                "$endpoint/entries/$articleId?access_token=${config.accessToken}",
                ArticleJson::class.java).body!!

    fun getArticlesOverview(size: Int) =
            restTemplate.getForEntity(
                    "$endpoint/entries/?select=fields.title,fields.thumbnail&content_type=articles&limit=$size&access_token=${config.accessToken}",
                    HashMap::class.java).body!!
}

@Component
@ConfigurationProperties(prefix = "contentful-api")
class ContentfulApiConfig{
    lateinit var endPoint: String
    lateinit var spaceId: String
    lateinit var accessToken: String
}
