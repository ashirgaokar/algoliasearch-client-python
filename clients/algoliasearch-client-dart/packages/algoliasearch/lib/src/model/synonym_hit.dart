// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/synonym_type.dart';

import 'package:json_annotation/json_annotation.dart';

part 'synonym_hit.g.dart';

@JsonSerializable()
final class SynonymHit {
  /// Returns a new [SynonymHit] instance.
  const SynonymHit({
    required this.objectID,
    required this.type,
    this.synonyms,
    this.input,
    this.word,
    this.corrections,
    this.placeholder,
    this.replacements,
  });

  /// Unique identifier of the synonym object to be created or updated.
  @JsonKey(name: r'objectID')
  final String objectID;

  @JsonKey(name: r'type')
  final SynonymType type;

  /// Words or phrases to be considered equivalent.
  @JsonKey(name: r'synonyms')
  final List<String>? synonyms;

  /// Word or phrase to appear in query strings (for onewaysynonym).
  @JsonKey(name: r'input')
  final String? input;

  /// Word or phrase to appear in query strings (for altcorrection1 and altcorrection2).
  @JsonKey(name: r'word')
  final String? word;

  /// Words to be matched in records.
  @JsonKey(name: r'corrections')
  final List<String>? corrections;

  /// Token to be put inside records.
  @JsonKey(name: r'placeholder')
  final String? placeholder;

  /// List of query words that will match the token.
  @JsonKey(name: r'replacements')
  final List<String>? replacements;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is SynonymHit &&
          other.objectID == objectID &&
          other.type == type &&
          other.synonyms == synonyms &&
          other.input == input &&
          other.word == word &&
          other.corrections == corrections &&
          other.placeholder == placeholder &&
          other.replacements == replacements;

  @override
  int get hashCode =>
      objectID.hashCode +
      type.hashCode +
      synonyms.hashCode +
      input.hashCode +
      word.hashCode +
      corrections.hashCode +
      placeholder.hashCode +
      replacements.hashCode;

  factory SynonymHit.fromJson(Map<String, dynamic> json) =>
      _$SynonymHitFromJson(json);

  Map<String, dynamic> toJson() => _$SynonymHitToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
