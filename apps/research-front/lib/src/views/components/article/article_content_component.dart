
import 'package:angular/angular.dart';
import 'package:angular/security.dart';

@Component(
  selector: 'article-content',
  styleUrls: ['article_content_component.css'],
  templateUrl: 'article_content_component.html',
  directives: [],
  providers: [],
)

class ArticleContentComponent {

  @Input()
  SafeHtml content;
}