import 'document.dart';
import 'fcc.dart';

class ArticleId {
  final String id;

  ArticleId(this.id);  
}

class Article {
  ArticleId id;
  ArticleOverview articleOverview;
  Document document;
  
    Article(this.id, this.articleOverview, this.document);
}

class ArticleOverview {
  String editor;
  String editorIcon;
  String title;
  String lastModified;
  String thumbnail;
  String description;

  ArticleOverview(this.editor, this.editorIcon, this.title, this.lastModified, this.thumbnail, this.description);
}

class Articles extends FCC<Article> {
  @override
  final List<Article> values;

  Articles(this.values);
  
}