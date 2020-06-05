import 'dart:html';

import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';

@Component(
  selector: 'search',
  styleUrls: ['search_component.css'],
  templateUrl: 'search_component.html',
  directives: [
    NgIf,
    NgFor,
    MaterialIconComponent,
    routerDirectives,
  ],
  exports: [],
)
class SearchComponent implements OnInit, OnDestroy {

  bool get focused => false;

  bool get hasList => false;

  bool get loading => false;

  String searchQuery = '';

  @ViewChild('query')
  InputElement inputElement;

  @override
  void ngOnInit() {
  }

  @override
  void ngOnDestroy() {
  }

  void onBlur() {

  }

  void onOver() {

  }

  void onFocus() {

  }

  void onClick() {

  }

}