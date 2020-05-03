// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'document.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ElementJson _$ElementJsonFromJson(Map<String, dynamic> json) {
  return ElementJson(
    json['tag'] as String,
    (json['content'] as List)
        ?.map((e) =>
            e == null ? null : ElementJson.fromJson(e as Map<String, dynamic>))
        ?.toList(),
    json['text'] as String,
    (json['attrs'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(k, e as String),
    ),
  );
}

Map<String, dynamic> _$ElementJsonToJson(ElementJson instance) =>
    <String, dynamic>{
      'tag': instance.tag,
      'content': instance.content,
      'text': instance.text,
      'attrs': instance.attrs,
    };
