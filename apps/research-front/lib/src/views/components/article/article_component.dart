
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:research_front/src/views/layouts/header.dart';

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

class ArticleComponent implements OnInit, OnDestroy {
  @override
  void ngOnDestroy() {
  }

  @override
  void ngOnInit() {
  }
  
}