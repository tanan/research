package com.tanan.researchserver.rest

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.action.Action
import com.tanan.researchserver.domain.action.UserAction
import com.tanan.researchserver.usecase.historylog.HistoryLogUseCase
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.router
import reactor.core.publisher.Mono

@Component
class HistoryLogHander(private val useCase: HistoryLogUseCase) {
    fun postHistoryLog(request: ServerRequest): Mono<ServerResponse> {
        val userId = request.queryParam("u")
        val action = request.queryParam("t")
        val pageId = request.queryParam("p")
        if(userId.isEmpty || action.isEmpty) return ServerResponse.status(400).body(Mono.just("NG"), String::class.java)
        return request
                .let { useCase.storeUserHistory(UserAction(Id(userId.get()), Action.fromActionType(action.get()), pageId.get())) }
                .let { ServerResponse
                        .status(200)
                        .body(Mono.just("OK"), String::class.java)}
    }
}

@Configuration
class HistoryLogResource(private val handler: HistoryLogHander) {
    @Bean
    fun historyLogRoutes() = router {
        POST("/v1/histories", handler::postHistoryLog)
    }
}