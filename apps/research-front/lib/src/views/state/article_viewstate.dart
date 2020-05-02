class ArticleViewState {
  List<Article> articles = [];

  Article currentArticle;

  Article init() => Article('', ArticleOverviewUnit('', '', '', '', ''), '');
}


class Article {
  String articleId;
  ArticleOverviewUnit overview;
  String content;

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