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
    json['editor'] as String,
    json['editorIcon'] as String,
    json['articleName'] as String,
    json['lastModified'] as String,
    json['thumbnail'] as String,
    json['description'] as String,
    json['content'] as Map<String, dynamic>,
  );
}

Map<String, dynamic> _$ArticleJsonToJson(ArticleJson instance) =>
    <String, dynamic>{
      'articleId': instance.articleId,
      'editor': instance.editor,
      'editorIcon': instance.editorIcon,
      'articleName': instance.articleName,
      'lastModified': instance.lastModified,
      'thumbnail': instance.thumbnail,
      'description': instance.description,
      'content': instance.content,
    };
