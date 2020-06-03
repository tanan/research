import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:intl/date_symbol_data_local.dart';
import 'package:research_front/src/driver/api/api_client.dart';
import 'package:research_front/src/driver/api/error_handler.dart';
import 'package:research_front/src/driver/api/research_api_url.dart';
import 'package:research_front/src/gateway/article_port.dart';
import 'package:research_front/src/presenter/article_presenter.dart';
import 'package:research_front/src/usecase/article_usecase.dart';
import 'package:research_front/src/views/app_component.template.dart' as ng;
import 'package:research_front/src/presenter/system_presenter.dart';
import 'package:research_front/src/views/state/article_viewstate.dart';
import 'package:research_front/src/views/state/system_viewstate.dart';
import 'package:http/browser_client.dart';

import 'main.template.dart' as self;

@GenerateInjector([
  ClassProvider(HttpErrorHandler),
  ClassProvider(SystemPresenter),
  ClassProvider(SystemViewState),
  ClassProvider(ArticleUsecase),
  ClassProvider(ArticlePort, useClass: ArticleGateway),
  ClassProvider(ArticlePresenter),
  ClassProvider(ArticleViewState),
  ClassProvider(ApiClient),
  ClassProvider(BrowserClient),
  ClassProvider(ResearchApiURL),
  routerProviders,
])

final InjectorFactory injector = self.injector$Injector;

void main() async {
  await initializeDateFormatting('ja');
  runApp(ng.AppComponentNgFactory, createInjector: injector);
}
