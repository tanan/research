// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'article.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

NodeJson _$NodeJsonFromJson(Map<String, dynamic> json) {
  return NodeJson(
    (json['content'] as List)
        ?.map((e) =>
            e == null ? null : NodeJson.fromJson(e as Map<String, dynamic>))
        ?.toList(),
    json['data'] as Map<String, dynamic>,
    json['nodeType'] as String,
    json['value'] as String,
  );
}

Map<String, dynamic> _$NodeJsonToJson(NodeJson instance) => <String, dynamic>{
      'content': instance.content,
      'data': instance.data,
      'nodeType': instance.nodeType,
      'value': instance.value,
    };

DataJson _$DataJsonFromJson(Map<String, dynamic> json) {
  return DataJson(
    json['target'] == null
        ? null
        : TargetJson.fromJson(json['target'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$DataJsonToJson(DataJson instance) => <String, dynamic>{
      'target': instance.target,
    };

TargetJson _$TargetJsonFromJson(Map<String, dynamic> json) {
  return TargetJson(
    json['sys'] == null
        ? null
        : SysJson.fromJson(json['sys'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$TargetJsonToJson(TargetJson instance) =>
    <String, dynamic>{
      'sys': instance.sys,
    };

SysJson _$SysJsonFromJson(Map<String, dynamic> json) {
  return SysJson(
    json['id'] as String,
  );
}

Map<String, dynamic> _$SysJsonToJson(SysJson instance) => <String, dynamic>{
      'id': instance.id,
    };

FieldsJson _$FieldsJsonFromJson(Map<String, dynamic> json) {
  return FieldsJson(
    json['body'] == null
        ? null
        : NodeJson.fromJson(json['body'] as Map<String, dynamic>),
    json['lastUpdated'] as String,
  );
}

Map<String, dynamic> _$FieldsJsonToJson(FieldsJson instance) =>
    <String, dynamic>{
      'body': instance.body,
      'lastUpdated': instance.lastUpdated,
    };

ArticleJson _$ArticleJsonFromJson(Map<String, dynamic> json) {
  return ArticleJson(
    json['id'] as String,
    json['articleOverview'] == null
        ? null
        : ArticleOverviewJson.fromJson(
            json['articleOverview'] as Map<String, dynamic>),
    json['content'] as String,
  );
}

Map<String, dynamic> _$ArticleJsonToJson(ArticleJson instance) =>
    <String, dynamic>{
      'id': instance.id,
      'articleOverview': instance.articleOverview,
      'content': instance.content,
    };

ArticleOverviewJson _$ArticleOverviewJsonFromJson(Map<String, dynamic> json) {
  return ArticleOverviewJson(
    json['editor'] as String,
    json['title'] as String,
    json['lastModified'] as String,
    json['thumbnail'] as String,
    json['description'] as String,
  );
}

Map<String, dynamic> _$ArticleOverviewJsonToJson(
        ArticleOverviewJson instance) =>
    <String, dynamic>{
      'editor': instance.editor,
      'title': instance.title,
      'lastModified': instance.lastModified,
      'thumbnail': instance.thumbnail,
      'description': instance.description,
    };

IncludesJson _$IncludesJsonFromJson(Map<String, dynamic> json) {
  return IncludesJson(
    (json['Entry'] as List)?.map((e) => e as Map<String, dynamic>)?.toList(),
  );
}

Map<String, dynamic> _$IncludesJsonToJson(IncludesJson instance) =>
    <String, dynamic>{
      'Entry': instance.Entry,
    };

EntriesJson _$EntriesJsonFromJson(Map<String, dynamic> json) {
  return EntriesJson(
    (json['items'] as List)
        ?.map((e) =>
            e == null ? null : ArticleJson.fromJson(e as Map<String, dynamic>))
        ?.toList(),
    json['includes'] == null
        ? null
        : IncludesJson.fromJson(json['includes'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$EntriesJsonToJson(EntriesJson instance) =>
    <String, dynamic>{
      'items': instance.items,
      'includes': instance.includes,
    };
