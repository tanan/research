
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/domain/article.dart' as d;
import 'package:research_front/src/usecase/article_usecase.dart';
import 'package:research_front/src/views/components/atoms/tag_component.dart';
import 'package:research_front/src/views/components/organisms/article/article_content_component.dart';
import 'package:research_front/src/views/components/organisms/header/header_component.dart';
import 'package:research_front/src/views/components/organisms/sidenav/sidenav_component.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';

@Component(
  selector: 'article',
  styleUrls: ['article_component.css'],
  templateUrl: 'article_component.html',
  directives: [
    MaterialIconComponent,
    materialInputDirectives,
    HeaderComponent,
    SidenavComponent,
    ArticleContentComponent,
    TagComponent,
  ],
  providers: [],
)

class ArticleComponent implements OnInit, OnActivate {
  final ArticleUsecase _usecase;
  final ArticleViewState _viewState;

  ArticleComponent(this._usecase, this._viewState);

  Article get article => _viewState.currentArticle ?? _viewState.init();

  @override
  void ngOnInit() {
  }

  @override
  void onActivate(RouterState previous, RouterState current) {
    _usecase.findArticle(d.ArticleId(current.parameters['id']));
  } 
}