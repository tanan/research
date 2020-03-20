import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/gateway/article_port.dart';
import 'package:research_front/src/presenter/article_presenter.dart';

class ArticleUsecase {
  final ArticlePort _port;
  final ArticlePresenter _presenter;
  
  ArticleUsecase(this._port, this._presenter);

  Future<void> findArticle(ArticleId id) async {
    await _article(id);
  }

  Future<void> _article(ArticleId id) async {
    var article = await _port.find(id);
    _presenter.setArticle(article);
  }
}