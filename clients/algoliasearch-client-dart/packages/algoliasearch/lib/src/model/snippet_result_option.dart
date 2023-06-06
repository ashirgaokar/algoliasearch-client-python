// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/match_level.dart';

import 'package:json_annotation/json_annotation.dart';

part 'snippet_result_option.g.dart';

@JsonSerializable()
final class SnippetResultOption {
  /// Returns a new [SnippetResultOption] instance.
  const SnippetResultOption({
    required this.value,
    required this.matchLevel,
  });

  /// Markup text with occurrences highlighted.
  @JsonKey(name: r'value')
  final String value;

  @JsonKey(name: r'matchLevel')
  final MatchLevel matchLevel;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is SnippetResultOption &&
          other.value == value &&
          other.matchLevel == matchLevel;

  @override
  int get hashCode => value.hashCode + matchLevel.hashCode;

  factory SnippetResultOption.fromJson(Map<String, dynamic> json) =>
      _$SnippetResultOptionFromJson(json);

  Map<String, dynamic> toJson() => _$SnippetResultOptionToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
