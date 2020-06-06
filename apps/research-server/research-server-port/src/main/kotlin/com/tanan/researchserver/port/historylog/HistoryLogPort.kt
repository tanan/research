package com.tanan.researchserver.port.historylog

import com.tanan.researchserver.domain.action.UserAction

interface HistoryLogPort {
    fun storeUserHistory(userAction: UserAction)
}