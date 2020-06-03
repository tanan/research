import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/views/routes.dart';

@Component(
  selector: 'my-app',
  styleUrls: [
    'package:angular_components/app_layout/layout.scss.css',
    'app_component.css'
  ],
  templateUrl: 'app_component.html',
  directives: [
    NgIf,
    routerDirectives,
  ],
  exports: [Routes, RoutePaths],
)
class AppComponent implements OnInit {

  // final SystemUsecase _systemUsecase;
  final Router _router;

  AppComponent(this._router);

  @override
  void ngOnInit() {
    // _systemUsecase.registerErrorHandler((error) async {
    //   _systemUsecase.rememberTargetRoute();
    //   await _router.navigate(RoutePaths.dashboard.toUrl());
    // });
  }
}
