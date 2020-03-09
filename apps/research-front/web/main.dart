import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:research_front/src/domain/article.dart';
import 'package:research_front/src/driver/api/error_handler.dart';
import 'package:research_front/src/views/app_component.template.dart' as ng;
import 'package:research_front/src/gateway/system_port.dart';
import 'package:research_front/src/presenter/system_presenter.dart';
import 'package:research_front/src/usecase/system_usecase.dart';
import 'package:research_front/src/driver/browser/environment.dart' as env;
import 'package:research_front/src/views/state/system_state.dart';

import 'main.template.dart' as self;

@GenerateInjector([
  ClassProvider(HttpErrorHandler),
  ClassProvider(SystemUsecase),
  ClassProvider(SystemPort, useClass: SystemGateway),
  ClassProvider(SystemPresenter),
  ClassProvider(SystemState),
  ClassProvider(env.Location, useClass: env.HashBaseLocation),
  ClassProvider(Article),
  routerProvidersHash,

])

final InjectorFactory injector = self.injector$Injector;

void main() {
  runApp(ng.AppComponentNgFactory, createInjector: injector);
}
