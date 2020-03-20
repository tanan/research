import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/driver/api/api_client.dart';
import 'package:research_front/src/driver/api/json/article.dart';


abstract class ArticlePort {
  // Future<Articles> findAll();
  Future<Article> find(ArticleId id);
}

class ArticleGateway implements ArticlePort {
  final ApiClient _client;

  ArticleGateway(this._client);

  // @override
  // Future<Articles> findAll() async =>
  //   (await _client.getArticles().toArticles());

  @override
  Future<Article> find(ArticleId id) async =>
    (await _client.getArticle(id.id)).toArticle();
}

extension ArticleConverter on ArticleJson {
  Article toArticle() =>
    Article(
      ArticleId(id),
      _toArticleOverview(),
      '');
  
  ArticleOverview _toArticleOverview() =>
    ArticleOverview(articleOverview.editor, articleOverview.title, articleOverview.lastModified, articleOverview.thumbnail, articleOverview.description);
}
