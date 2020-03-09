
import 'package:research_front/src/domain/system.dart';
import 'package:research_front/src/views/state/system_state.dart';

class SystemPresenter {

  final SystemState _state;

  SystemPresenter(this._state);

  void save(Location location) {
    _state.setLocation(location.path, location.parameterAsString());
  }
}