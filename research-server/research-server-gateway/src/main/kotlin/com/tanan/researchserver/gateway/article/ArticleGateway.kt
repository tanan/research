package com.tanan.researchserver.gateway.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.*
import com.tanan.researchserver.port.article.ArticlePort
import org.springframework.stereotype.Component

@Component
class ArticleGateway(): ArticlePort {
    override fun getArticleOverview(id: Id): ArticleOverview =
            ArticleOverview(Name(""), Description(""), Category(""), Url(""))
}