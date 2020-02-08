package com.tanan.researchserver.rest.resources

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.rest.jsons.article.ArticleOverviewJson
import com.tanan.researchserver.rest.jsons.article.toJson
import com.tanan.researchserver.usecase.article.ArticleUseCase
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.router
import reactor.core.publisher.Mono

@Component
class ArticleHandler(private val articleUseCase: ArticleUseCase) {
    fun getArticleOverview(request: ServerRequest): Mono<ServerResponse> =
            request.pathVariable("id")
                    .let { articleUseCase.getArticleOverview(Id(it)) }
                    .let { ServerResponse
                            .status(200)
                            .body(Mono.just(it.toJson()), ArticleOverviewJson::class.java)
                    }
}

@Configuration
class ArticleRouter(private val handler: ArticleHandler) {
    @Bean
    fun articleRoutes() = router {
        "/v1/articles/{id}".nest {
            GET("/overview", handler::getArticleOverview)
        }
    }
}
