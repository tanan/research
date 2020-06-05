
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:research_front/src/views/components/organisms/article_cards/article_cards_component.dart';
import 'package:research_front/src/views/components/organisms/header/header_component.dart';
import 'package:research_front/src/views/components/organisms/sidenav/sidenav_component.dart';

@Component(
  selector: 'dashboard',
  styleUrls: ['dashboard_component.css'],
  templateUrl: 'dashboard_component.html',
  directives: [
    MaterialIconComponent,
    materialInputDirectives,
    ArticleCardsComponent,
    HeaderComponent,
    SidenavComponent,
  ],
  providers: [],
)

class DashboardComponent {
}