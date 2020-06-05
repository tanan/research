import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:research_front/src/views/components/molecules/search/search_component.dart';

@Component(
  selector: 'header',
  styleUrls: [
    'header_component.css'
  ],
  templateUrl: 'header_component.html',
  directives: [
    SearchComponent,
    MaterialIconComponent,
  ]
)

class HeaderComponent {
}
