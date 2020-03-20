import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

class ArticlePresenter {

  final ArticleViewState _state;

  ArticlePresenter(this._state);

  void setArticle(Article article) {
    _state.articleId = article.id.id;
    _state.overview = _overview(article.articleOverview);
  }

  ArticleOverviewUnit _overview(ArticleOverview overview) =>
    ArticleOverviewUnit(
      overview.editor,
      overview.title,
      overview.lastModified,
      overview.thumbnail,
      overview.description
    );
}