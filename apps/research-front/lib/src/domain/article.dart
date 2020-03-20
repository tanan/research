import 'package:equatable/equatable.dart';

import 'fcc.dart';

class ArticleId extends Equatable {
  final String id;

  ArticleId(this.id);

  @override
  List<Object> get props => [id];
  
}

class Article {
  ArticleId id;
  ArticleOverview articleOverview;
  String content;

  Article(this.id, this.articleOverview, this.content);
}

class ArticleOverview {
  String editor;
  String title;
  String lastModified;
  String thumbnail;
  String description;

  ArticleOverview(this.editor, this.title, this.lastModified, this.thumbnail, this.description);
}

class Articles extends FCC<Article> {
  @override
  final List<Article> values;

  Articles(this.values);
  
}