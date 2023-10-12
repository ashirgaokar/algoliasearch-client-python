/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.recommend

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*
import kotlin.jvm.JvmInline

/**
 * Precision of a geographical search (in meters), to [group results that are more or less the same distance from a central point](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/in-depth/geo-ranking-precision/).
 *
 * Implementations:
 * - [Int] - *[AroundPrecision.of]*
 * - [List<AroundPrecisionFromValueInner>] - *[AroundPrecision.of]*
 */
@Serializable(AroundPrecisionSerializer::class)
public sealed interface AroundPrecision {

  @JvmInline
  public value class IntValue(public val value: Int) : AroundPrecision

  @JvmInline
  public value class ListOfAroundPrecisionFromValueInnerValue(public val value: List<AroundPrecisionFromValueInner>) : AroundPrecision

  public companion object {

    /** [AroundPrecision] as [Int] Value. */
    public fun of(value: Int): AroundPrecision {
      return IntValue(value)
    }

    /** [AroundPrecision] as [List<AroundPrecisionFromValueInner>] Value. */
    public fun of(value: List<AroundPrecisionFromValueInner>): AroundPrecision {
      return ListOfAroundPrecisionFromValueInnerValue(value)
    }
  }
}

internal class AroundPrecisionSerializer : KSerializer<AroundPrecision> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("AroundPrecision")

  override fun serialize(encoder: Encoder, value: AroundPrecision) {
    when (value) {
      is AroundPrecision.IntValue -> Int.serializer().serialize(encoder, value.value)
      is AroundPrecision.ListOfAroundPrecisionFromValueInnerValue -> ListSerializer(AroundPrecisionFromValueInner.serializer()).serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): AroundPrecision {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize Int
    if (tree is JsonPrimitive) {
      try {
        val value = codec.json.decodeFromJsonElement(Int.serializer(), tree)
        return AroundPrecision.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize Int (error: ${e.message})")
      }
    }

    // deserialize List<AroundPrecisionFromValueInner>
    if (tree is JsonArray) {
      try {
        val value = codec.json.decodeFromJsonElement(ListSerializer(AroundPrecisionFromValueInner.serializer()), tree)
        return AroundPrecision.of(value)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize List<AroundPrecisionFromValueInner> (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
