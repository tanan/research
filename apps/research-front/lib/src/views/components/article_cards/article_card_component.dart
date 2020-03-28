import 'package:angular/angular.dart';
import 'package:angular/core.dart';
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

  @Input()
  Article article;

  @override
  void ngOnInit() {}
  
}