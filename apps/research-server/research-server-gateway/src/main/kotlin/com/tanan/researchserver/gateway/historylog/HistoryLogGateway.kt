package com.tanan.researchserver.gateway.historylog

import com.tanan.researchserver.domain.action.UserAction
import com.tanan.researchserver.driver.history.database.UserHistoryLogDriver
import com.tanan.researchserver.driver.history.database.entity.UserHistoryLog
import com.tanan.researchserver.port.historylog.HistoryLogPort
import org.springframework.stereotype.Component

@Component
class HistoryLogGateway(private val userHistoryLogDriver: UserHistoryLogDriver): HistoryLogPort {
    override fun storeUserHistory(userAction: UserAction) {
        userHistoryLogDriver.save(UserHistoryLog(
                dt = userAction.date,
                userId = userAction.id.value,
                actionType = userAction.action.id,
                pageId = userAction.pageId,
                session = userAction.requestInfo.session,
                ip = userAction.requestInfo.ipAddress,
                host = userAction.requestInfo.host,
                path = userAction.requestInfo.path,
                query = userAction.requestInfo.query,
                referer = userAction.requestInfo.referer,
                ua = userAction.requestInfo.userAgent
        ))
    }
}