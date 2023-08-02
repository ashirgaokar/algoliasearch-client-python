// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:json_annotation/json_annotation.dart';

/// Type of the Destination, defines in which Algolia product the data will be stored.
@JsonEnum(valueField: 'raw')
enum DestinationType {
  /// Type of the Destination, defines in which Algolia product the data will be stored.
  search(r'search'),

  /// Type of the Destination, defines in which Algolia product the data will be stored.
  insights(r'insights'),

  /// Type of the Destination, defines in which Algolia product the data will be stored.
  flow(r'flow');

  const DestinationType(this.raw);
  final dynamic raw;

  dynamic toJson() => raw;

  static DestinationType fromJson(dynamic json) {
    for (final value in values) {
      if (value.raw == json) return value;
    }
    throw ArgumentError.value(json, "raw", "No enum value with that value");
  }

  @override
  String toString() => raw.toString();
}
