import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/views/components/search/search_component.dart';

@Component(
  selector: 'header',
  styleUrls: [
    'package:angular_components/app_layout/layout.scss.css',
    'header.css'
  ],
  templateUrl: 'header.html',
  directives: [
    SearchComponent,
    MaterialIconComponent,
    routerDirectives,
  ]
)

class HeaderComponent {
  // final Router _router;

  // HeaderComponent(this._router);  
}
