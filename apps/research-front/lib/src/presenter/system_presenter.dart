
import 'package:research_front/src/domain/system.dart';
import 'package:research_front/src/views/state/system_viewstate.dart';

class SystemPresenter {

  final SystemViewState _state;

  SystemPresenter(this._state);

  void save(Location location) {
    _state.setLocation(location.path, location.parameterAsString());
  }
}