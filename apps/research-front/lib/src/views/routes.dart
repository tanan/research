
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/views/components/dashboard/dashboard_component.template.dart' as dashboard;

class RoutePaths {
  static final dashboard = RoutePath(path: 'dashboard');
  static final article = RoutePath(path: 'article');
  static String defaultUrl() {
    return dashboard.toUrl();
  }
}

class Routes {
  static final all = <RouteDefinition>[
    RouteDefinition(
      routePath: RoutePaths.dashboard,
      component: dashboard.DashboardComponentNgFactory,
    ),
    // RouteDefinition(
    //   routePath: RoutePaths.article,
    //   component: article.ArticleComponentNgFactory,
    // ),
    RouteDefinition.redirect(
      path: '',
      redirectTo: RoutePaths.defaultUrl(),
    ),
  ];
}
