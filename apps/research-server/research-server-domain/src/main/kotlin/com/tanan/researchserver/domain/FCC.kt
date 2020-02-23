package com.tanan.researchserver.domain

interface FCC<T> {
    val list: List<T>

    fun <R> map(transform: (T) -> (R)): List<R> = list.map(transform)
}