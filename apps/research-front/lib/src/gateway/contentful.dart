import 'dart:convert';

import 'package:optional/optional.dart';
import 'package:research_front/src/domain/document.dart';

class Contentful {
  final Map<String, dynamic> fields;
  final Map<String, dynamic> includes;

  Contentful(this.fields, this.includes);

  Document asDocument() =>
    Document(Element(Tag('div'), _toNodes(fields['content'] as List<dynamic>)));

  Nodes _toNodes(List<dynamic> src) =>
    Nodes(src.map((v) => _toNode(v as Map<String, dynamic>)).where((v) => v.isPresent).map((v) => v.value).toList());

  Optional<Node> _toNode(Map<String, dynamic> value) {
    var nodeType = value['nodeType'];
    switch(nodeType) {
      case 'paragraph':
        return Optional.of(Element(Tag('p'), _toNodes(value['content'] as List<dynamic>)));
      case 'text':
        return Optional.of(_getText(value));
      case 'unordered-list':
        return Optional.of(Element(Tag('ul'), _toNodes(value['content'] as List<dynamic>)));
      case 'list-item':
        return Optional.of(Element(Tag('li'), _toNodes(value['content'] as List<dynamic>)));
      case 'embedded-asset-block':
        return Optional.of(_toImage(value));
      default:
        return Optional.empty();
    }
  }

  Node _toImage(Map<String, dynamic> value) {
    var id = value['data']['target']['sys']['id'];
    var target = (includes['Asset'] as List<dynamic>).where((v) => v['sys']['id'] == id) as Map<String, dynamic>;
    return Element(Tag('img'), Nodes([]), attributes: {'src': target['fields']['url']});
  }

  Node _getText(Map<String, dynamic> value) {
    if (_isBold(value)) {
      return Element(Tag('em'), Nodes([_toTextNode(value)]));
    } else if (_isCode(value)) {
      return Element(Tag('pre'), Nodes([_toCode(value)]));
    }
    return _toTextNode(value);
  }

  bool _isCode(Map<String, dynamic> value) => value['marks'].where((v) => v['type'] == 'code').isNotEmpty;

  bool _isBold(Map<String, dynamic> value) => value['marks'].where((v) => v['type'] == 'bold').isNotEmpty;

  Element _toCode(Map<String, dynamic> value) => Element(Tag('code'), Nodes([_toTextNode(value)]));

  TextNode _toTextNode(Map<String, dynamic> value) => TextNode(_toEscapeText(value['value']));

  String _toEscapeText(String text) => HtmlEscape().convert(text);


}