import 'package:json_annotation/json_annotation.dart';

part 'article.g.dart';

@JsonSerializable()
class NodeJson {
  final List<NodeJson> content;
  final Map<String, dynamic> data;
  final String nodeType;
  final String value;

  NodeJson(this.content, this.data, this.nodeType, this.value);
  factory NodeJson.fromJson(Map<String, dynamic> json) => _$NodeJsonFromJson(json);
}

@JsonSerializable()
class DataJson {
  final TargetJson target;

  DataJson(this.target);
  factory DataJson.fromJson(Map<String, dynamic> json) => _$DataJsonFromJson(json);
}

@JsonSerializable()
class TargetJson {
  final SysJson sys;

  TargetJson(this.sys);
  factory TargetJson.fromJson(Map<String, dynamic> json) => _$TargetJsonFromJson(json);
}

@JsonSerializable()
class SysJson {
  final String id;

  SysJson(this.id);
  factory SysJson.fromJson(Map<String, dynamic> json) => _$SysJsonFromJson(json);
}

@JsonSerializable()
class FieldsJson {
  final NodeJson body;
  final String lastUpdated;

  FieldsJson(this.body, this.lastUpdated);
  factory FieldsJson.fromJson(Map<String, dynamic> json) => _$FieldsJsonFromJson(json);
}

@JsonSerializable()
class ArticleJson {
  final String id;
  final ArticleOverviewJson articleOverview;
  final String content;

  ArticleJson(this.id, this.articleOverview, this.content);
  factory ArticleJson.fromJson(Map<String, dynamic> json) => _$ArticleJsonFromJson(json);
}

@JsonSerializable()
class ArticleOverviewJson {
  final String editor;
  final String title;
  final String lastModified;
  final String thumbnail;
  final String description;

  ArticleOverviewJson(this.editor, this.title, this.lastModified, this.thumbnail, this.description);
  factory ArticleOverviewJson.fromJson(Map<String, dynamic> json) => _$ArticleOverviewJsonFromJson(json);
}

@JsonSerializable()
class IncludesJson {
  final List<Map<String, dynamic>> Entry;
  
  IncludesJson(this.Entry);
  factory IncludesJson.fromJson(Map<String, dynamic> json) => _$IncludesJsonFromJson(json);
}

@JsonSerializable()
class EntriesJson {
  final List<ArticleJson> items;
  final IncludesJson includes;
  
  EntriesJson(this.items, this.includes);
  factory EntriesJson.fromJson(Map<String, dynamic> json) => _$EntriesJsonFromJson(json);
}