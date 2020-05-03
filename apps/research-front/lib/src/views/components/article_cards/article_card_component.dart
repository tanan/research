import 'package:angular/angular.dart';
import 'package:angular/core.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/views/routes.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

@Component(
  selector: 'article-card',
  styleUrls: ['article_card_component.css'],
  templateUrl: 'article_card_component.html',
  directives: [
    NgFor,
  ],
  exports: [],
)
class ArticleCardComponent implements OnInit {

  final Router _router;

  @Input()
  Article article;

  ArticleCardComponent(this._router);

  void onClickArticle() {
    _router.navigate(RoutePaths.article.toUrl(parameters: {'id': article.articleId}));
  }

  @override
  void ngOnInit() {}
  
}