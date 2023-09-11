// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetApiKeyResponse */
public class GetApiKeyResponse {

  @JsonProperty("value")
  private String value;

  @JsonProperty("createdAt")
  private Long createdAt;

  @JsonProperty("acl")
  private List<Acl> acl = null;

  @JsonProperty("description")
  private String description;

  @JsonProperty("indexes")
  private List<String> indexes;

  @JsonProperty("maxHitsPerQuery")
  private Integer maxHitsPerQuery;

  @JsonProperty("maxQueriesPerIPPerHour")
  private Integer maxQueriesPerIPPerHour;

  @JsonProperty("queryParameters")
  private String queryParameters;

  @JsonProperty("referers")
  private List<String> referers;

  @JsonProperty("validity")
  private Integer validity;

  public GetApiKeyResponse setValue(String value) {
    this.value = value;
    return this;
  }

  /** API key. */
  @javax.annotation.Nullable
  public String getValue() {
    return value;
  }

  public GetApiKeyResponse setCreatedAt(Long createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /**
   * Timestamp of creation in milliseconds in [Unix epoch
   * time](https://wikipedia.org/wiki/Unix_time).
   */
  @javax.annotation.Nonnull
  public Long getCreatedAt() {
    return createdAt;
  }

  public GetApiKeyResponse setAcl(List<Acl> acl) {
    this.acl = acl;
    return this;
  }

  public GetApiKeyResponse addAcl(Acl aclItem) {
    this.acl.add(aclItem);
    return this;
  }

  /**
   * [Permissions](https://www.algolia.com/doc/guides/security/api-keys/#access-control-list-acl)
   * associated with the key.
   */
  @javax.annotation.Nonnull
  public List<Acl> getAcl() {
    return acl;
  }

  public GetApiKeyResponse setDescription(String description) {
    this.description = description;
    return this;
  }

  /** Description of an API key for you and your team members. */
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }

  public GetApiKeyResponse setIndexes(List<String> indexes) {
    this.indexes = indexes;
    return this;
  }

  public GetApiKeyResponse addIndexes(String indexesItem) {
    if (this.indexes == null) {
      this.indexes = new ArrayList<>();
    }
    this.indexes.add(indexesItem);
    return this;
  }

  /**
   * Restricts this API key to a list of indices or index patterns. If the list is empty, all
   * indices are allowed. Specify either an exact index name or a pattern with a leading or trailing
   * wildcard character (or both). For example: - `dev_*` matches all indices starting with \"dev_\"
   * - `*_dev` matches all indices ending with \"_dev\" - `*_products_*` matches all indices
   * containing \"_products_\".
   */
  @javax.annotation.Nullable
  public List<String> getIndexes() {
    return indexes;
  }

  public GetApiKeyResponse setMaxHitsPerQuery(Integer maxHitsPerQuery) {
    this.maxHitsPerQuery = maxHitsPerQuery;
    return this;
  }

  /**
   * Maximum number of hits this API key can retrieve in one query. If zero, no limit is enforced. >
   * **Note**: Use this parameter to protect you from third-party attempts to retrieve your entire
   * content by massively querying the index.
   */
  @javax.annotation.Nullable
  public Integer getMaxHitsPerQuery() {
    return maxHitsPerQuery;
  }

  public GetApiKeyResponse setMaxQueriesPerIPPerHour(Integer maxQueriesPerIPPerHour) {
    this.maxQueriesPerIPPerHour = maxQueriesPerIPPerHour;
    return this;
  }

  /**
   * Maximum number of API calls per hour allowed from a given IP address or [user
   * token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/). Each time an API
   * call is performed with this key, a check is performed. If there were more than the specified
   * number of calls within the last hour, the API returns an error with the status code `429` (Too
   * Many Requests). > **Note**: Use this parameter to protect you from third-party attempts to
   * retrieve your entire content by massively querying the index.
   */
  @javax.annotation.Nullable
  public Integer getMaxQueriesPerIPPerHour() {
    return maxQueriesPerIPPerHour;
  }

  public GetApiKeyResponse setQueryParameters(String queryParameters) {
    this.queryParameters = queryParameters;
    return this;
  }

  /**
   * Force some [query parameters](https://www.algolia.com/doc/api-reference/api-parameters/) to be
   * applied for each query made with this API key. It's a URL-encoded query string.
   */
  @javax.annotation.Nullable
  public String getQueryParameters() {
    return queryParameters;
  }

  public GetApiKeyResponse setReferers(List<String> referers) {
    this.referers = referers;
    return this;
  }

  public GetApiKeyResponse addReferers(String referersItem) {
    if (this.referers == null) {
      this.referers = new ArrayList<>();
    }
    this.referers.add(referersItem);
    return this;
  }

  /**
   * Restrict this API key to specific
   * [referrers](https://www.algolia.com/doc/guides/security/api-keys/in-depth/api-key-restrictions/#http-referrers).
   * If empty, all referrers are allowed. For example: - `https://algolia.com/_*` matches all
   * referrers starting with \"https://algolia.com/\" - `*.algolia.com` matches all referrers ending
   * with \".algolia.com\" - `*algolia.com*` allows everything in the domain \"algolia.com\".
   */
  @javax.annotation.Nullable
  public List<String> getReferers() {
    return referers;
  }

  public GetApiKeyResponse setValidity(Integer validity) {
    this.validity = validity;
    return this;
  }

  /**
   * Validity duration of a key (in seconds). The key will automatically be removed after this time
   * has expired. The default value of 0 never expires. Short-lived API keys are useful to grant
   * temporary access to your data. For example, in mobile apps, you can't [control when users
   * update your
   * app](https://www.algolia.com/doc/guides/security/security-best-practices/#use-secured-api-keys-in-mobile-apps).
   * So instead of encoding keys into your app as you would for a web app, you should dynamically
   * fetch them from your mobile app's backend.
   */
  @javax.annotation.Nullable
  public Integer getValidity() {
    return validity;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetApiKeyResponse getApiKeyResponse = (GetApiKeyResponse) o;
    return (
      Objects.equals(this.value, getApiKeyResponse.value) &&
      Objects.equals(this.createdAt, getApiKeyResponse.createdAt) &&
      Objects.equals(this.acl, getApiKeyResponse.acl) &&
      Objects.equals(this.description, getApiKeyResponse.description) &&
      Objects.equals(this.indexes, getApiKeyResponse.indexes) &&
      Objects.equals(this.maxHitsPerQuery, getApiKeyResponse.maxHitsPerQuery) &&
      Objects.equals(this.maxQueriesPerIPPerHour, getApiKeyResponse.maxQueriesPerIPPerHour) &&
      Objects.equals(this.queryParameters, getApiKeyResponse.queryParameters) &&
      Objects.equals(this.referers, getApiKeyResponse.referers) &&
      Objects.equals(this.validity, getApiKeyResponse.validity)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      value,
      createdAt,
      acl,
      description,
      indexes,
      maxHitsPerQuery,
      maxQueriesPerIPPerHour,
      queryParameters,
      referers,
      validity
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetApiKeyResponse {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    acl: ").append(toIndentedString(acl)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    indexes: ").append(toIndentedString(indexes)).append("\n");
    sb.append("    maxHitsPerQuery: ").append(toIndentedString(maxHitsPerQuery)).append("\n");
    sb.append("    maxQueriesPerIPPerHour: ").append(toIndentedString(maxQueriesPerIPPerHour)).append("\n");
    sb.append("    queryParameters: ").append(toIndentedString(queryParameters)).append("\n");
    sb.append("    referers: ").append(toIndentedString(referers)).append("\n");
    sb.append("    validity: ").append(toIndentedString(validity)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
