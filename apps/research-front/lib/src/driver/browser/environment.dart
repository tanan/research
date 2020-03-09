import 'dart:html';

import 'package:research_front/src/driver/browser/extractor.dart';

abstract class Location {
  String get path;
  Map<String, String> get parameters;
}

class HashBaseLocation implements Location {

  @override
  Map<String, String> get parameters => Extractor(_base).asMap();

  @override
  String get path => Extractor(_base).path();

  String get _base => window.location.hash;
}
