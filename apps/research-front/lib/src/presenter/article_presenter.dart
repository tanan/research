import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/views/state/article_viewstate.dart';

class ArticlePresenter {

  final ArticleViewState _state;

  ArticlePresenter(this._state);

  void setArticle(d.Article article) {
    _state.articles.add(Article(article.id.id, _overview(article.articleOverview), null));
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