import 'package:equatable/equatable.dart';

import 'fcc.dart';

class Document extends Equatable {
  final Element body;

  Document(this.body);

  @override
  List<Object> get props => [body];

  @override
  String toString() {
    return asHtml();
  }

  String asHtml() {
    return '<body>${body.asHtml()}</body>';
  }

  String contentAsHtml() {
    return body.asHtml();
  }
}

class Tag extends Equatable {
  final String value;

  Tag(this.value);

  @override
  List<Object> get props => [value];
}

abstract class Node {
  String asHtml();
}

class Nodes extends FCC<Node> {
  @override
  final List<Node> values;

  Nodes(this.values);
}

class TextNode extends Equatable implements Node {
  final String value;

  TextNode(this.value);

  @override
  List<Object> get props => [value];

  @override
  String asHtml() => value;

  @override
  String toString() {
    return asHtml();
  }
}

class Element extends Equatable implements Node {
  final Tag tag;
  final Nodes children;
  final Map<String, String> attributes;

  Element(this.tag, this.children, { Map<String, String> attributes }): attributes = attributes ?? {};

  @override
  List<Object> get props => [tag, children, attributes];

  @override
  String asHtml() {
    var attr = attributes.entries.map((e) => '${e.key}=\"${e.value}\"').join(' ');
    attr = attr.isNotEmpty ? ' ${attr}' : '';
    return children.isNotEmpty
      ? '<${tag.value}${attr}>${children.map((v) => v.asHtml()).join()}</${tag.value}>'
      : '<${tag.value}${attr}>';
  }

  @override
  String toString() {
    return asHtml();
  }
}

class Elements extends FCC<Element> {
  @override
  final List<Element> values;

  Elements(this.values);
}
