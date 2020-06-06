package com.tanan.researchserver.gateway.historylog

import com.tanan.researchserver.domain.action.UserAction
import com.tanan.researchserver.driver.history.database.UserHistoryLogDriver
import com.tanan.researchserver.driver.history.database.entity.UserHistoryLog
import com.tanan.researchserver.port.historylog.HistoryLogPort
import org.springframework.stereotype.Component

@Component
class HistoryLogGateway(private val userHistoryLogDriver: UserHistoryLogDriver): HistoryLogPort {
    override fun storeUserHistory(userAction: UserAction) {
        userHistoryLogDriver.save(UserHistoryLog(userId = userAction.id.value, actionType = userAction.action.id, pageId = userAction.pageId))
    }
}