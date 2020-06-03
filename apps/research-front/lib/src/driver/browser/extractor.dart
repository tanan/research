import 'package:optional/optional.dart';
import 'package:tuple/tuple.dart';

class Extractor {
  final String src;

  Extractor(this.src);

  String path() =>
    Optional.ofNullable(RegExp(r'^#(\/.+)$').firstMatch(src))
      .map((v) => v.group(1))
      .orElse(null);

  String parameters() =>
    Optional.ofNullable(RegExp(r'^#(\/.+)\?(.+)$').firstMatch(src))
      .map((v) => v.group(2))
      .orElse(null);

}