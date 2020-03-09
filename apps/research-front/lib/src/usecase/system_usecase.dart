import 'package:research_front/src/domain/error.dart';
import 'package:research_front/src/gateway/system_port.dart';
import 'package:research_front/src/presenter/system_presenter.dart';

class SystemUsecase {
  final SystemPort _port;
  final SystemPresenter _presenter;

  SystemUsecase(this._port, this._presenter);

  void registerErrorHandler(ErrorHandler handler) {
    _port.register(handler);
  }

  void rememberTargetRoute() {
    _presenter.save(_port.getLocation());
  }
}