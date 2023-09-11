// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** RuleResponseMetadata */
public class RuleResponseMetadata {

  @JsonProperty("lastUpdate")
  private String lastUpdate;

  public RuleResponseMetadata setLastUpdate(String lastUpdate) {
    this.lastUpdate = lastUpdate;
    return this;
  }

  /** Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. */
  @javax.annotation.Nullable
  public String getLastUpdate() {
    return lastUpdate;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RuleResponseMetadata ruleResponseMetadata = (RuleResponseMetadata) o;
    return Objects.equals(this.lastUpdate, ruleResponseMetadata.lastUpdate);
  }

  @Override
  public int hashCode() {
    return Objects.hash(lastUpdate);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RuleResponseMetadata {\n");
    sb.append("    lastUpdate: ").append(toIndentedString(lastUpdate)).append("\n");
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
