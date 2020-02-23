package com.tanan.researchserver.driver.article.contentful.jsons

data class ArticlesOverviewJson(
        val sys: String,
        val total: Int,
        val skip: Int,
        val limit: Int,
        val items: List<ItemJson>,
        val includes: List<IncludeJson>
)

data class ItemJson(
        val fields: ItemFieldsJson
)

data class ItemFieldsJson(
        val title: String,
        val thumbnail: ItemFieldsThumbnailJson
)

data class ItemFieldsThumbnailJson(
        val sys: ItemFieldsThumbnailSysJson
)

data class ItemFieldsThumbnailSysJson(
        val type: String,
        val linkType: String,
        val id: String
)

data class IncludeJson(
        val Asset: List<IncludeAssetJson>
)

data class IncludeAssetJson(
        val sys: IncludeAssetSysJson,
        val fields: IncludeAssetFieldsJson
)

data class IncludeAssetSysJson(
        val space: String,
        val id: String,
        val type: String,
        val createdAt: String,
        val updatedAt: String,
        val environment: String,
        val revision: Int,
        val locale: String
)

data class IncludeAssetFieldsJson(
        val title: String,
        val description: String,
        val file: IncludeAssetFieldsFileJson,
        val fileName: String,
        val contentType: String
)

data class IncludeAssetFieldsFileJson(
        val url: String,
        val details: String
)