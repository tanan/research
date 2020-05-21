import 'package:json_annotation/json_annotation.dart';

part 'article.g.dart';

@JsonSerializable()
class ArticlesJson {
  final List<ArticleJson> articles;

  ArticlesJson(this.articles);

  factory ArticlesJson.fromJson(Map<String, dynamic> json) =>
    _$ArticlesJsonFromJson(json);
}

@JsonSerializable()
class ArticleJson {
  final String articleId;
  final ArticleOverviewJson overview;
  final Map<String, dynamic> content;

  ArticleJson(this.articleId, this.overview, this.content);
  factory ArticleJson.fromJson(Map<String, dynamic> json) => _$ArticleJsonFromJson(json);
}

@JsonSerializable()
class ArticleOverviewJson {
  final EditorJson editor;
  final String articleName;
  final String lastModified;
  final String thumbnail;
  final String description;

  ArticleOverviewJson(this.editor, this.articleName, this.lastModified, this.thumbnail, this.description);
  factory ArticleOverviewJson.fromJson(Map<String, dynamic> json) => _$ArticleOverviewJsonFromJson(json);
}

@JsonSerializable()
class EditorJson {
  final String editorName;
  final String editorIcon;

  EditorJson(this.editorName, this.editorIcon);
  factory EditorJson.fromJson(Map<String, dynamic> json) => _$EditorJsonFromJson(json);
}