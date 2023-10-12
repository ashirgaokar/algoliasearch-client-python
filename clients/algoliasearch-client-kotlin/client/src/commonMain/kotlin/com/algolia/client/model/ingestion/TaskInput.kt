/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * TaskInput
 *
 * Implementations:
 * - [OnDemandDateUtilsInput]
 * - [ScheduleDateUtilsInput]
 */
@Serializable(TaskInputSerializer::class)
public sealed interface TaskInput {

  public companion object {
  }
}

internal class TaskInputSerializer : KSerializer<TaskInput> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("TaskInput")

  override fun serialize(encoder: Encoder, value: TaskInput) {
    when (value) {
      is OnDemandDateUtilsInput -> OnDemandDateUtilsInput.serializer().serialize(encoder, value)
      is ScheduleDateUtilsInput -> ScheduleDateUtilsInput.serializer().serialize(encoder, value)
    }
  }

  override fun deserialize(decoder: Decoder): TaskInput {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize OnDemandDateUtilsInput
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement(OnDemandDateUtilsInput.serializer(), tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize OnDemandDateUtilsInput (error: ${e.message})")
      }
    }

    // deserialize ScheduleDateUtilsInput
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement(ScheduleDateUtilsInput.serializer(), tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ScheduleDateUtilsInput (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
