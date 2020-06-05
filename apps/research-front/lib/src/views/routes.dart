
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/views/components/pages/dashboard/dashboard_component.template.dart' as dashboard;
import 'package:research_front/src/views/components/pages/article/article_component.template.dart' as a;

class RoutePaths {
  static final dashboard = RoutePath(path: 'dashboard');
  static final article = RoutePath(path: 'articles/:id');
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
    RouteDefinition(
      routePath: RoutePaths.article,
      component: a.ArticleComponentNgFactory,
    ),
    RouteDefinition.redirect(
      path: '',
      redirectTo: RoutePaths.defaultUrl(),
    ),
  ];
}
