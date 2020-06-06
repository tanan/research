package com.tanan.researchserver.driver.history.database.entity

import javax.persistence.*

@Entity
@Table(name = "user_history_log", schema = "history")
data class UserHistoryLog(
        @Id
        @GeneratedValue(strategy = GenerationType.IDENTITY)
        var id: Int = 0,
        val userId: String,
        val actionType: Int,
        val pageId: String? = null
)
