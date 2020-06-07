package com.tanan.researchserver.driver.history.database.entity

import java.util.*
import javax.persistence.*

@Entity
@Table(name = "user_history_log", schema = "history")
data class UserHistoryLog(
        @Id
        @GeneratedValue(strategy = GenerationType.IDENTITY)
        var id: Int = 0,
        val dt: Date,
        val userId: String,
        val actionType: Int,
        val pageId: String? = null,
        val session : String,
        val ip: String?,
        val host: String?,
        val path: String?,
        val query: String?,
        val referer: String?,
        val ua: String?
)
