/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * SearchQuery
 *
 * Implementations:
 * - [SearchForFacets]
 * - [SearchForHits]
 */
@Serializable(SearchQuerySerializer::class)
public sealed interface SearchQuery {

  public companion object {
  }
}

internal class SearchQuerySerializer : KSerializer<SearchQuery> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("SearchQuery")

  override fun serialize(encoder: Encoder, value: SearchQuery) {
    when (value) {
      is SearchForFacets -> SearchForFacets.serializer().serialize(encoder, value)
      is SearchForHits -> SearchForHits.serializer().serialize(encoder, value)
    }
  }

  override fun deserialize(decoder: Decoder): SearchQuery {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize SearchForFacets
    if (tree is JsonObject && tree.containsKey("facet") && tree.containsKey("type")) {
      try {
        return codec.json.decodeFromJsonElement(SearchForFacets.serializer(), tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SearchForFacets (error: ${e.message})")
      }
    }

    // deserialize SearchForHits
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement(SearchForHits.serializer(), tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SearchForHits (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
