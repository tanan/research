import 'package:angular/angular.dart';

@Component(
  selector: 'article-tag',
  styleUrls: ['tag_component.css'],
  templateUrl: 'tag_component.html',
  directives: [],
  exports: [],
)
class TagComponent {
  @Input()
  String tag;
}