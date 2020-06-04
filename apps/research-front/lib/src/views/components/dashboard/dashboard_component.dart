
import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:research_front/src/views/components/article_cards/article_cards_component.dart';
import 'package:research_front/src/views/components/sidenav/sidenav_component.dart';
import 'package:research_front/src/views/layouts/header.dart';

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

class DashboardComponent implements OnInit, OnDestroy {
  @override
  void ngOnDestroy() {
    // TODO: implement ngOnDestroy
  }

  @override
  void ngOnInit() {
    // TODO: implement ngOnInit
  }
  
}