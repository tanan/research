

import 'package:optional/optional.dart';
import 'package:research_front/src/domain/document.dart';
import 'package:research_front/src/driver/api/json/document.dart';

class DocumentParser {
  Document parse(ElementJson json) {
    return Document(_parse(json));
  }

  Node _parse(ElementJson json) =>
    Optional.ofNullable(json.text)
      .map<Node>((v) => TextNode(v))
      .orElseGet(() => Element(Tag(json.tag), _contentToNodes(json.content), attributes: _attrs(json.tag, json.attrs)));
  
  Nodes _contentToNodes(List<ElementJson> elements) =>
    Nodes(elements.map((v) => _parse(v)).toList());

  Map<String, String> _attrs(String tag, Map<String, String> attrs) {
    if (tag != 'img') {
      return attrs;
    }
    return attrs.map((k, v) => MapEntry(k, v));
    // return attrs.map((k, v) => MapEntry(k, k == 'src' ? v : v));
  }
  
}