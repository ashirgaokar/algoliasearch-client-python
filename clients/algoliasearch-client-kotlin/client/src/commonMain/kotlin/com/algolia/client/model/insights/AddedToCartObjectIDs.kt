/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Use this event to track when users add items to their shopping cart unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track add-to-cart events related to Algolia requests, use the \"Added to cart object IDs after search\" event.
 *
 * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
 * @param eventType
 * @param eventSubtype
 * @param index Name of the Algolia index.
 * @param objectIDs List of object identifiers for items of an Algolia index.
 * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
 * @param objectData Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.
 * @param currency If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR.
 * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
 * @param authenticatedUserToken User token for authenticated users.
 */
@Serializable
public data class AddedToCartObjectIDs(

  /** Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.  */
  @SerialName(value = "eventName") val eventName: String,

  @SerialName(value = "eventType") val eventType: ConversionEvent,

  @SerialName(value = "eventSubtype") val eventSubtype: AddToCartEvent,

  /** Name of the Algolia index. */
  @SerialName(value = "index") val index: String,

  /** List of object identifiers for items of an Algolia index. */
  @SerialName(value = "objectIDs") val objectIDs: List<String>,

  /** Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.  */
  @SerialName(value = "userToken") val userToken: String,

  /** Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.  */
  @SerialName(value = "objectData") val objectData: List<ObjectData>? = null,

  /** If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR. */
  @SerialName(value = "currency") val currency: String? = null,

  /** Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.  */
  @SerialName(value = "timestamp") val timestamp: Long? = null,

  /** User token for authenticated users. */
  @SerialName(value = "authenticatedUserToken") val authenticatedUserToken: String? = null,
) : EventsItems
