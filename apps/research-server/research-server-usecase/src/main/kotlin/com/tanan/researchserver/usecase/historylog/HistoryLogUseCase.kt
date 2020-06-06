package com.tanan.researchserver.usecase.historylog

import com.tanan.researchserver.domain.action.UserAction
import com.tanan.researchserver.port.historylog.HistoryLogPort
import org.springframework.stereotype.Component

@Component
class HistoryLogUseCase(private val historyLogPort: HistoryLogPort) {
    fun storeUserHistory(action: UserAction) {
        return historyLogPort.storeUserHistory(action)
    }
}