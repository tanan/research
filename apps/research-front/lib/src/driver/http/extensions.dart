import 'dart:convert';

import 'package:http/http.dart';

const CONTENT_TYPE_JSON_UTF8 = {'content-type': 'application/json; charset=UTF-8'};

typedef Converter<T> = T Function(Map<String, dynamic> value);

extension JSONParsing on Response {
  T asMap<T>(Converter<T> c) {
    var value = json.decode(utf8.decode(bodyBytes)) as Map<String, dynamic>;
    return c(value);
  }

  List<T> asList<T>(Converter<T> c) {
    var value = json.decode(body) as List<Map<String, dynamic>>;
    return value.map((v) => c(v)).toList();
  }
}