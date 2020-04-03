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
  final String editor;
  final String articleName;
  final String lastModified;
  final String thumbnail;
  final String description;
  final String content;

  ArticleJson(this.articleId, this.editor, this.articleName, this.lastModified, this.thumbnail, this.description, this.content);
  factory ArticleJson.fromJson(Map<String, dynamic> json) => _$ArticleJsonFromJson(json);
}
