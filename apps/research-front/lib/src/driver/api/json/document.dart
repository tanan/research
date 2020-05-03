import 'package:json_annotation/json_annotation.dart';

part 'document.g.dart';

@JsonSerializable()
class ElementJson {
  final String tag;
  final List<ElementJson> content;
  final String text;
  final Map<String, String> attrs;

  ElementJson(this.tag, this.content, this.text, this.attrs);

  factory ElementJson.fromJson(Map<String, dynamic> src) => null;
  
}
