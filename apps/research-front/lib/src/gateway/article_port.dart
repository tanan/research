import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/driver/api/api_client.dart';
import 'package:research_front/src/driver/api/json/article.dart';
import 'package:research_front/src/gateway/contentful.dart';


abstract class ArticlePort {
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
    return (await _client.getArticles()).toArticles();
  }
}

extension on ArticlesJson {
  Articles toArticles() {
    return Articles(articles.map((v) => v.toArticle()).toList());
  }
}

extension on ArticleJson {
  Article toArticle() {
    return Article(
      ArticleId(articleId),
      _toArticleOverview(),
      content != null ? Contentful(content, includes).asDocument() : null);
  }

  ArticleOverview _toArticleOverview() =>
    ArticleOverview(
      overview.editor.editorName,
      overview.editor.editorIcon,
      overview.articleName,
      overview.lastModified,
      overview.thumbnail,
      overview.description
    );
}
