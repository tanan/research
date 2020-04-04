import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/views/state/article_viewstate.dart';

class ArticlePresenter {

  final ArticleViewState _state;

  ArticlePresenter(this._state);

  void showArticles(d.Articles articles) {
    var a = <Article>[];
    articles.values.forEach((v) => a.add(Article(v.id.id, _overview(v.articleOverview), v.content)));
    _state.articles = a;
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