import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/views/state/article_viewstate.dart';

class ArticlePresenter {

  final ArticleViewState _state;

  ArticlePresenter(this._state);

  void showArticleContent(d.Article article) {
    print(article.id.id);
    _state.currentArticle = Article(article.id.id, _overview(article.articleOverview), article.content);
  }

  void showArticles(d.Articles articles) {
    _state.articles = articles.mapToList((v) => Article(v.id.id, _overview(v.articleOverview), v.content));
  }

  ArticleOverviewUnit _overview(d.ArticleOverview overview) =>
    ArticleOverviewUnit(
      overview.editor,
      overview.title,
      overview.lastModified,
      overview.thumbnail,
      overview.description
    );
}