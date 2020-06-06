package com.tanan.researchserver.rest

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.router
import reactor.core.publisher.Mono

@Component
class SystemHandler {
    fun ping(@Suppress("UNUSED_PARAMETER") request: ServerRequest): Mono<ServerResponse> =
            ServerResponse.status(200).body(Mono.just("pong"), String::class.java)
}

@Configuration
class SystemRouter(private val handler: SystemHandler) {
    @Bean
    fun systemRoutes() = router {
        GET("/v1/systems/ping", handler::ping)
    }
}