import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/driver/api/api_client.dart';
import 'package:research_front/src/driver/api/json/article.dart';


abstract class ArticlePort {
  // Future<Articles> findAll();
  Future<Article> find(ArticleId id);
  Future<Articles> findLatest();
}

class ArticleGateway implements ArticlePort {
  final ApiClient _client;

  ArticleGateway(this._client);

  @override
  Future<Article> find(ArticleId id) async =>
    (await _client.getArticle(id.id)).toArticle();

  @override
  Future<Articles> findLatest() async {
    var articles = Articles(<Article>[]);
    (await _client.getArticles()).articles.forEach((v) => {
      articles.values.add(v.toArticle())
    });
    return articles;
  }
}

extension on ArticlesJson {
  Articles toArticles() {
    articles.forEach((v) => print(v));
    return Articles(articles.map((v) => v.toArticle()));
  }
}

extension on ArticleJson {
  Article toArticle() =>
    Article(
      ArticleId(articleId),
      _toArticleOverview(),
      '');
  
  ArticleOverview _toArticleOverview() =>
    ArticleOverview(editor, articleName, lastModified, thumbnail, description);
}
