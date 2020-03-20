
class ArticleViewState {
  String articleId;
  ArticleOverviewUnit overview;
  String content;
}

class ArticleOverviewUnit {
  String editor;
  String title;
  String lastModified;
  String thumbnail;
  String description;

  ArticleOverviewUnit(this.editor, this.title, this.lastModified, this.thumbnail, this.description);
}