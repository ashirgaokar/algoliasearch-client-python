/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * BaseIndexSettings
 *
 * @param replicas Creates [replicas](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/), which are copies of a primary index with the same records but different settings.
 * @param paginationLimitedTo Maximum number of hits accessible through pagination.
 * @param unretrievableAttributes Attributes that can't be retrieved at query time.
 * @param disableTypoToleranceOnWords Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/).
 * @param attributesToTransliterate Attributes in your index to which [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead) applies. This will ensure that words indexed in Katakana or Kanji can also be searched in Hiragana.
 * @param camelCaseAttributes Attributes on which to split [camel case](https://wikipedia.org/wiki/Camel_case) words.
 * @param decompoundedAttributes Attributes in your index to which [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding) applies.
 * @param indexLanguages Set the languages of your index, for language-specific processing steps such as [tokenization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/tokenization/) and [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
 * @param disablePrefixOnAttributes Attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search).
 * @param allowCompressionOfIntegerArray Incidates whether the engine compresses arrays with exclusively non-negative integers. When enabled, the compressed arrays may be reordered.
 * @param numericAttributesForFiltering Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters).
 * @param separatorsToIndex Controls which separators are added to an Algolia index as part of [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean). Separators are all non-letter characters except spaces and currency characters, such as $€£¥.
 * @param searchableAttributes [Attributes used for searching](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/), including determining [if matches at the beginning of a word are important (ordered) or not (unordered)](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/how-to/configuring-searchable-attributes-the-right-way/#understanding-word-position).
 * @param userData Lets you store custom data in your indices.
 * @param customNormalization A list of characters and their normalized replacements to override Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
 * @param attributeForDistinct Name of the deduplication attribute to be used with Algolia's [_distinct_ feature](https://www.algolia.com/doc/guides/managing-results/refine-results/grouping/#introducing-algolias-distinct-feature).
 */
@Serializable
public data class BaseIndexSettings(

  /** Creates [replicas](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/), which are copies of a primary index with the same records but different settings. */
  @SerialName(value = "replicas") val replicas: List<String>? = null,

  /** Maximum number of hits accessible through pagination. */
  @SerialName(value = "paginationLimitedTo") val paginationLimitedTo: Int? = null,

  /** Attributes that can't be retrieved at query time. */
  @SerialName(value = "unretrievableAttributes") val unretrievableAttributes: List<String>? = null,

  /** Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/). */
  @SerialName(value = "disableTypoToleranceOnWords") val disableTypoToleranceOnWords: List<String>? = null,

  /** Attributes in your index to which [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead) applies. This will ensure that words indexed in Katakana or Kanji can also be searched in Hiragana. */
  @SerialName(value = "attributesToTransliterate") val attributesToTransliterate: List<String>? = null,

  /** Attributes on which to split [camel case](https://wikipedia.org/wiki/Camel_case) words. */
  @SerialName(value = "camelCaseAttributes") val camelCaseAttributes: List<String>? = null,

  /** Attributes in your index to which [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding) applies. */
  @SerialName(value = "decompoundedAttributes") val decompoundedAttributes: JsonObject? = null,

  /** Set the languages of your index, for language-specific processing steps such as [tokenization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/tokenization/) and [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/). */
  @SerialName(value = "indexLanguages") val indexLanguages: List<String>? = null,

  /** Attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search). */
  @SerialName(value = "disablePrefixOnAttributes") val disablePrefixOnAttributes: List<String>? = null,

  /** Incidates whether the engine compresses arrays with exclusively non-negative integers. When enabled, the compressed arrays may be reordered.  */
  @SerialName(value = "allowCompressionOfIntegerArray") val allowCompressionOfIntegerArray: Boolean? = null,

  /** Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters). */
  @SerialName(value = "numericAttributesForFiltering") val numericAttributesForFiltering: List<String>? = null,

  /** Controls which separators are added to an Algolia index as part of [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean). Separators are all non-letter characters except spaces and currency characters, such as $€£¥. */
  @SerialName(value = "separatorsToIndex") val separatorsToIndex: String? = null,

  /** [Attributes used for searching](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/), including determining [if matches at the beginning of a word are important (ordered) or not (unordered)](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/how-to/configuring-searchable-attributes-the-right-way/#understanding-word-position).  */
  @SerialName(value = "searchableAttributes") val searchableAttributes: List<String>? = null,

  /** Lets you store custom data in your indices. */
  @SerialName(value = "userData") val userData: JsonElement? = null,

  /** A list of characters and their normalized replacements to override Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/). */
  @SerialName(value = "customNormalization") val customNormalization: Map<kotlin.String, Map<kotlin.String, String>>? = null,

  /** Name of the deduplication attribute to be used with Algolia's [_distinct_ feature](https://www.algolia.com/doc/guides/managing-results/refine-results/grouping/#introducing-algolias-distinct-feature). */
  @SerialName(value = "attributeForDistinct") val attributeForDistinct: String? = null,
)
