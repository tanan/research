package com.tanan.researchserver.usecase.article

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.article.*
import com.tanan.researchserver.port.article.ArticlePort
import org.springframework.stereotype.Component

@Component
//class ArticleUseCase() {
//    fun getArticleOverview(id: Id): ArticleOverview = ArticleOverview(Name(""), Description(""), Category(""), Url(""))
//}
class ArticleUseCase(private val articlePort: ArticlePort) {
    fun getArticleOverview(id: Id): ArticleOverview = articlePort.getArticleOverview(id)
}