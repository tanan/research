
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/usecase/article_usecase.dart';
import 'package:research_front/src/views/layouts/header.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

@Component(
  selector: 'article',
  styleUrls: ['article_component.css'],
  templateUrl: 'article_component.html',
  directives: [
    MaterialIconComponent,
    materialInputDirectives,
    HeaderComponent,
  ],
  providers: [],
)

class ArticleComponent implements OnInit, OnActivate {
  final ArticleUsecase _usecase;
  final ArticleViewState _viewState;

  ArticleComponent(this._usecase, this._viewState);

  Article get article => _viewState.currentArticle;

  @override
  void ngOnInit() {
  }

  @override
  void onActivate(RouterState previous, RouterState current) {
    _usecase.findArticle(d.ArticleId(current.parameters['id']));
  } 
}