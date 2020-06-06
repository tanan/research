package com.tanan.researchserver.driver.history.database

import com.tanan.researchserver.driver.history.database.entity.UserHistoryLog
import org.springframework.data.jpa.repository.JpaRepository

interface UserHistoryLogDriver : JpaRepository<UserHistoryLog, String> {}