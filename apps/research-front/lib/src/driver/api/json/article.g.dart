// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'article.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ArticlesJson _$ArticlesJsonFromJson(Map<String, dynamic> json) {
  return ArticlesJson(
    (json['articles'] as List)
        ?.map((e) =>
            e == null ? null : ArticleJson.fromJson(e as Map<String, dynamic>))
        ?.toList(),
  );
}

Map<String, dynamic> _$ArticlesJsonToJson(ArticlesJson instance) =>
    <String, dynamic>{
      'articles': instance.articles,
    };

ArticleJson _$ArticleJsonFromJson(Map<String, dynamic> json) {
  return ArticleJson(
    json['articleId'] as String,
    json['overview'] == null
        ? null
        : ArticleOverviewJson.fromJson(
            json['overview'] as Map<String, dynamic>),
    json['content'] as Map<String, dynamic>,
    json['includes'] as Map<String, dynamic>,
  );
}

Map<String, dynamic> _$ArticleJsonToJson(ArticleJson instance) =>
    <String, dynamic>{
      'articleId': instance.articleId,
      'overview': instance.overview,
      'content': instance.content,
      'includes': instance.includes,
    };

ArticleOverviewJson _$ArticleOverviewJsonFromJson(Map<String, dynamic> json) {
  return ArticleOverviewJson(
    json['editor'] == null
        ? null
        : EditorJson.fromJson(json['editor'] as Map<String, dynamic>),
    json['articleName'] as String,
    json['lastModified'] as String,
    json['thumbnail'] as String,
    json['description'] as String,
  );
}

Map<String, dynamic> _$ArticleOverviewJsonToJson(
        ArticleOverviewJson instance) =>
    <String, dynamic>{
      'editor': instance.editor,
      'articleName': instance.articleName,
      'lastModified': instance.lastModified,
      'thumbnail': instance.thumbnail,
      'description': instance.description,
    };

EditorJson _$EditorJsonFromJson(Map<String, dynamic> json) {
  return EditorJson(
    json['editorName'] as String,
    json['editorIcon'] as String,
  );
}

Map<String, dynamic> _$EditorJsonToJson(EditorJson instance) =>
    <String, dynamic>{
      'editorName': instance.editorName,
      'editorIcon': instance.editorIcon,
    };
