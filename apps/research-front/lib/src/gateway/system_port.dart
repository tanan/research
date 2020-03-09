import 'package:research_front/src/domain/error.dart';
import 'package:research_front/src/domain/system.dart';
import 'package:research_front/src/driver/api/error_handler.dart';
import 'package:research_front/src/driver/browser/environment.dart' as environment;

abstract class SystemPort {
  void register(ErrorHandler handler);

  Location getLocation();
}

class SystemGateway implements SystemPort {
  
  final HttpErrorHandler _handler;
  final environment.Location _location;

  SystemGateway(this._handler, this._location);

  @override
  Location getLocation() {
    return Location(_location.path, _location.parameters);
  }

  @override
  void register(ErrorHandler handler) {
    _handler.subscribe((error) => handler(Error(error.message)));
  }
  
}