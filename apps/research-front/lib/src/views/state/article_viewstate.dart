import 'package:angular/security.dart';

class ArticleViewState {
  List<Article> articles = [];

  Article currentArticle;

  Article init() => Article('', ArticleOverviewUnit('', '', '', '', ''), null);
}


class Article {
  final String articleId;
  final ArticleOverviewUnit overview;
  final SafeHtml content;

  Article(this.articleId, this.overview, this.content);
}

class ArticleOverviewUnit {
  String editor;
  String title;
  String lastModified;
  String thumbnail;
  String description;

  ArticleOverviewUnit(this.editor, this.title, this.lastModified, this.thumbnail, this.description);
}