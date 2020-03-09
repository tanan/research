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

  Map<String, String> asMap() {
    var keyValuePairs =Optional.ofNullable(parameters())
      .map((v) => v.split('&'))
      .orElse([])
      .map((v) => v.split('='))
      .map((v) => Tuple2(v[0], v[1]));
    return { for (var v in keyValuePairs) v.item1: v.item2 };
  }
}