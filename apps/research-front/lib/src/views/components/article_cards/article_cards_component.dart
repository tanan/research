import 'package:angular/angular.dart';
import 'package:angular/core.dart';
import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/views/components/article_cards/article_card_component.dart';

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

  // @Input()
  // List<Article> articles;

  List<Article> articles = [
    Article('https://media.sproutsocial.com/uploads/2017/02/10x-featured-social-media-image-size.png', 'title', 'description', 'body'),
    Article('https://media.sproutsocial.com/uploads/2017/02/10x-featured-social-media-image-size.png', 'title', 'description', 'body'),
    Article('https://media.sproutsocial.com/uploads/2017/02/10x-featured-social-media-image-size.png', 'title', 'description', 'body'),
    Article('https://media.sproutsocial.com/uploads/2017/02/10x-featured-social-media-image-size.png', 'title', 'description', 'body'),
  ];

  @override
  void ngOnInit() {
    // TODO: implement ngOnInit
  }
  
}