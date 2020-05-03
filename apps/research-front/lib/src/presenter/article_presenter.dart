import 'package:angular/security.dart';
import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/domain/document.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

class ArticlePresenter {

  final ArticleViewState _state;
  final DomSanitizationService _sanitizer;

  ArticlePresenter(this._state, this._sanitizer);

  void clear() {
    _state.articles = [];
    _state.currentArticle = _state.init();
  }

  void showArticleContent(d.Article article) {
    _state.currentArticle = Article(article.id.id, _overview(article.articleOverview), _toHtml(article.document));
  }

  void showArticles(d.Articles articles) {
    _state.articles = articles.mapToList((v) => Article(v.id.id, _overview(v.articleOverview), _toHtml(v.document)));
  }

  ArticleOverviewUnit _overview(d.ArticleOverview overview) =>
    ArticleOverviewUnit(
      overview.editor,
      overview.title,
      overview.lastModified,
      overview.thumbnail,
      overview.description
    );
  
  SafeHtml _toHtml(Document document) =>
    document != null ? _sanitizer.bypassSecurityTrustHtml(document.contentAsHtml()) : null;
}