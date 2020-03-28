import 'package:angular/angular.dart';
import 'package:angular/core.dart';
import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/usecase/article_usecase.dart';
import 'package:research_front/src/views/components/article_cards/article_card_component.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

@Component(
  selector: 'article-cards',
  styleUrls: ['article_cards_component.css'],
  templateUrl: 'article_cards_component.html',
  directives: [
    NgFor,
    ArticleCardComponent,
  ],
  exports: [],
)
class ArticleCardsComponent implements OnInit {

  final ArticleUsecase _usecase;
  final ArticleViewState _viewState;

  ArticleCardsComponent(this._usecase, this._viewState);

  List<Article> get articles => _viewState.articles;

  @override
  void ngOnInit() {
    _usecase.findArticle(d.ArticleId('5qopcYbpSTaV7a1AIqzkSu'));
  }
  
}