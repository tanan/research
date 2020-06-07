package com.tanan.researchserver.domain.request

import java.util.*

data class RequestInfo(
        val session : String,
        val ipAddress: String?,
        val host: String?,
        val path: String?,
        val query: String?,
        val referer: String?,
        val userAgent: String?
)