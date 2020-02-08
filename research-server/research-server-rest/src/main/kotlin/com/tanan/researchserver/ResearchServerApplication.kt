package com.tanan.researchserver

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.ComponentScan

@SpringBootApplication
@ComponentScan(basePackages = ["research-server-domain", "research-server-port"])
class ResearchServerRestApplication

fun main(args: Array<String>) {
    runApplication<ResearchServerRestApplication>(*args)
}
