package com.tanan.researchserver.domain.action

import com.tanan.researchserver.domain.Id
import com.tanan.researchserver.domain.request.RequestInfo
import java.util.*

data class UserAction(
        val id: Id,
        val date: Date,
        val action: Action,
        val pageId: String?,
        val requestInfo: RequestInfo
)

enum class Action(val id: Int, val actionType: String) {
    PAGE_VIEW(1, "pv"),
    LIKE(2, "like"),
    BOOKMARK(3, "bookmark");


    companion object {
        fun fromActionType(type: String): Action {
            return values().firstOrNull() { it.actionType == type } ?: PAGE_VIEW
        }
    }
}