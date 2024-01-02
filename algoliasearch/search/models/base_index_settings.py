# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Any, Dict, List, Optional, Self

from pydantic import BaseModel, Field, StrictBool, StrictInt, StrictStr


class BaseIndexSettings(BaseModel):
    """
    BaseIndexSettings
    """

    replicas: Optional[List[StrictStr]] = Field(
        default=None,
        description="Creates [replicas](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/), which are copies of a primary index with the same records but different settings.",
    )
    pagination_limited_to: Optional[StrictInt] = Field(
        default=1000,
        description="Maximum number of hits accessible through pagination.",
        alias="paginationLimitedTo",
    )
    unretrievable_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes that can't be retrieved at query time.",
        alias="unretrievableAttributes",
    )
    disable_typo_tolerance_on_words: Optional[List[StrictStr]] = Field(
        default=None,
        description="Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/).",
        alias="disableTypoToleranceOnWords",
    )
    attributes_to_transliterate: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes in your index to which [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead) applies. This will ensure that words indexed in Katakana or Kanji can also be searched in Hiragana.",
        alias="attributesToTransliterate",
    )
    camel_case_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes on which to split [camel case](https://wikipedia.org/wiki/Camel_case) words.",
        alias="camelCaseAttributes",
    )
    decompounded_attributes: Optional[Dict[str, Any]] = Field(
        default=None,
        description="Attributes in your index to which [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding) applies.",
        alias="decompoundedAttributes",
    )
    index_languages: Optional[List[StrictStr]] = Field(
        default=None,
        description="Set the languages of your index, for language-specific processing steps such as [tokenization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/tokenization/) and [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).",
        alias="indexLanguages",
    )
    disable_prefix_on_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search).",
        alias="disablePrefixOnAttributes",
    )
    allow_compression_of_integer_array: Optional[StrictBool] = Field(
        default=False,
        description="Incidates whether the engine compresses arrays with exclusively non-negative integers. When enabled, the compressed arrays may be reordered. ",
        alias="allowCompressionOfIntegerArray",
    )
    numeric_attributes_for_filtering: Optional[List[StrictStr]] = Field(
        default=None,
        description="Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters).",
        alias="numericAttributesForFiltering",
    )
    separators_to_index: Optional[StrictStr] = Field(
        default="",
        description="Controls which separators are added to an Algolia index as part of [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean). Separators are all non-letter characters except spaces and currency characters, such as $€£¥.",
        alias="separatorsToIndex",
    )
    searchable_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="[Attributes used for searching](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/), including determining [if matches at the beginning of a word are important (ordered) or not (unordered)](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/how-to/configuring-searchable-attributes-the-right-way/#understanding-word-position). ",
        alias="searchableAttributes",
    )
    user_data: Optional[Any] = Field(
        default=None,
        description="Lets you store custom data in your indices.",
        alias="userData",
    )
    custom_normalization: Optional[Dict[str, Dict[str, StrictStr]]] = Field(
        default=None,
        description="A list of characters and their normalized replacements to override Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).",
        alias="customNormalization",
    )
    attribute_for_distinct: Optional[StrictStr] = Field(
        default=None,
        description="Name of the deduplication attribute to be used with Algolia's [_distinct_ feature](https://www.algolia.com/doc/guides/managing-results/refine-results/grouping/#introducing-algolias-distinct-feature).",
        alias="attributeForDistinct",
    )

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of BaseIndexSettings from a JSON string"""
        return cls.from_dict(loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={},
            exclude_none=True,
        )
        # set to None if user_data (nullable) is None
        # and model_fields_set contains the field
        if self.user_data is None and "user_data" in self.model_fields_set:
            _dict["userData"] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of BaseIndexSettings from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "replicas": obj.get("replicas"),
                "paginationLimitedTo": obj.get("paginationLimitedTo")
                if obj.get("paginationLimitedTo") is not None
                else 1000,
                "unretrievableAttributes": obj.get("unretrievableAttributes"),
                "disableTypoToleranceOnWords": obj.get("disableTypoToleranceOnWords"),
                "attributesToTransliterate": obj.get("attributesToTransliterate"),
                "camelCaseAttributes": obj.get("camelCaseAttributes"),
                "decompoundedAttributes": obj.get("decompoundedAttributes"),
                "indexLanguages": obj.get("indexLanguages"),
                "disablePrefixOnAttributes": obj.get("disablePrefixOnAttributes"),
                "allowCompressionOfIntegerArray": obj.get(
                    "allowCompressionOfIntegerArray"
                )
                if obj.get("allowCompressionOfIntegerArray") is not None
                else False,
                "numericAttributesForFiltering": obj.get(
                    "numericAttributesForFiltering"
                ),
                "separatorsToIndex": obj.get("separatorsToIndex")
                if obj.get("separatorsToIndex") is not None
                else "",
                "searchableAttributes": obj.get("searchableAttributes"),
                "userData": obj.get("userData"),
                "customNormalization": obj.get("customNormalization"),
                "attributeForDistinct": obj.get("attributeForDistinct"),
            }
        )
        return _obj
