package com.tanan.researchserver.domain.action

import com.tanan.researchserver.domain.Id

data class UserAction(
        val id: Id,
        val action: Action,
        val pageId: String?
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