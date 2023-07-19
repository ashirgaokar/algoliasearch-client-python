// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.querysuggestions;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** QuerySuggestionsConfigurationResponseAllOf */
public class QuerySuggestionsConfigurationResponseAllOf {

  @JsonProperty("appId")
  private String appId;

  @JsonProperty("sourceIndicesAPIKey")
  private String sourceIndicesAPIKey;

  @JsonProperty("suggestionsIndicesAPIKey")
  private String suggestionsIndicesAPIKey;

  @JsonProperty("externalIndicesAPIKey")
  private String externalIndicesAPIKey;

  public QuerySuggestionsConfigurationResponseAllOf setAppId(String appId) {
    this.appId = appId;
    return this;
  }

  /**
   * Your Algolia application ID.
   *
   * @return appId
   */
  @javax.annotation.Nullable
  public String getAppId() {
    return appId;
  }

  public QuerySuggestionsConfigurationResponseAllOf setSourceIndicesAPIKey(String sourceIndicesAPIKey) {
    this.sourceIndicesAPIKey = sourceIndicesAPIKey;
    return this;
  }

  /**
   * API key used to read from your source index.
   *
   * @return sourceIndicesAPIKey
   */
  @javax.annotation.Nullable
  public String getSourceIndicesAPIKey() {
    return sourceIndicesAPIKey;
  }

  public QuerySuggestionsConfigurationResponseAllOf setSuggestionsIndicesAPIKey(String suggestionsIndicesAPIKey) {
    this.suggestionsIndicesAPIKey = suggestionsIndicesAPIKey;
    return this;
  }

  /**
   * API key used to write and configure your Query Suggestions index.
   *
   * @return suggestionsIndicesAPIKey
   */
  @javax.annotation.Nullable
  public String getSuggestionsIndicesAPIKey() {
    return suggestionsIndicesAPIKey;
  }

  public QuerySuggestionsConfigurationResponseAllOf setExternalIndicesAPIKey(String externalIndicesAPIKey) {
    this.externalIndicesAPIKey = externalIndicesAPIKey;
    return this;
  }

  /**
   * API key used to read from external Algolia indices.
   *
   * @return externalIndicesAPIKey
   */
  @javax.annotation.Nullable
  public String getExternalIndicesAPIKey() {
    return externalIndicesAPIKey;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    QuerySuggestionsConfigurationResponseAllOf querySuggestionsConfigurationResponseAllOf = (QuerySuggestionsConfigurationResponseAllOf) o;
    return (
      Objects.equals(this.appId, querySuggestionsConfigurationResponseAllOf.appId) &&
      Objects.equals(this.sourceIndicesAPIKey, querySuggestionsConfigurationResponseAllOf.sourceIndicesAPIKey) &&
      Objects.equals(this.suggestionsIndicesAPIKey, querySuggestionsConfigurationResponseAllOf.suggestionsIndicesAPIKey) &&
      Objects.equals(this.externalIndicesAPIKey, querySuggestionsConfigurationResponseAllOf.externalIndicesAPIKey)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(appId, sourceIndicesAPIKey, suggestionsIndicesAPIKey, externalIndicesAPIKey);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class QuerySuggestionsConfigurationResponseAllOf {\n");
    sb.append("    appId: ").append(toIndentedString(appId)).append("\n");
    sb.append("    sourceIndicesAPIKey: ").append(toIndentedString(sourceIndicesAPIKey)).append("\n");
    sb.append("    suggestionsIndicesAPIKey: ").append(toIndentedString(suggestionsIndicesAPIKey)).append("\n");
    sb.append("    externalIndicesAPIKey: ").append(toIndentedString(externalIndicesAPIKey)).append("\n");
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
