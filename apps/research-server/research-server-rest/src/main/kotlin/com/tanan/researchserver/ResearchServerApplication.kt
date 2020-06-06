package com.tanan.researchserver

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Configuration
import org.springframework.web.reactive.config.CorsRegistry
import org.springframework.web.reactive.config.WebFluxConfigurer

@SpringBootApplication
class ResearchServerRestApplication

fun main(args: Array<String>) {
    runApplication<ResearchServerRestApplication>(*args)
}


@Configuration
class WebConfig: WebFluxConfigurer {
    override fun addCorsMappings(registry: CorsRegistry) {
        registry.addMapping("/**")
                .allowedOrigins("http://localhost:8080")
                .allowedOrigins("http://research.weeekend.work")
    }
}