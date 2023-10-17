/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Run
 *
 * @param runID The run UUID.
 * @param appID
 * @param taskID The task UUID.
 * @param status
 * @param type
 * @param createdAt Date of creation (RFC3339 format).
 * @param progress
 * @param outcome
 * @param failureThreshold A percentage representing the accepted failure threshold to determine if a `run` succeeded or not.
 * @param reason Explains the result of outcome.
 * @param reasonCode
 * @param startedAt Date of start (RFC3339 format).
 * @param finishedAt Date of finish (RFC3339 format).
 */
@Serializable
public data class Run(

  /** The run UUID. */
  @SerialName(value = "runID") val runID: String,

  @SerialName(value = "appID") val appID: String,

  /** The task UUID. */
  @SerialName(value = "taskID") val taskID: String,

  @SerialName(value = "status") val status: RunStatus,

  @SerialName(value = "type") val type: RunType,

  /** Date of creation (RFC3339 format). */
  @SerialName(value = "createdAt") val createdAt: String,

  @SerialName(value = "progress") val progress: RunProgress? = null,

  @SerialName(value = "outcome") val outcome: RunOutcome? = null,

  /** A percentage representing the accepted failure threshold to determine if a `run` succeeded or not. */
  @SerialName(value = "failureThreshold") val failureThreshold: Int? = null,

  /** Explains the result of outcome. */
  @SerialName(value = "reason") val reason: String? = null,

  @SerialName(value = "reasonCode") val reasonCode: RunReasonCode? = null,

  /** Date of start (RFC3339 format). */
  @SerialName(value = "startedAt") val startedAt: String? = null,

  /** Date of finish (RFC3339 format). */
  @SerialName(value = "finishedAt") val finishedAt: String? = null,
)
