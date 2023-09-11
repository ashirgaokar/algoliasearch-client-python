// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** DeleteApiKeyResponse */
public class DeleteApiKeyResponse {

  @JsonProperty("deletedAt")
  private String deletedAt;

  public DeleteApiKeyResponse setDeletedAt(String deletedAt) {
    this.deletedAt = deletedAt;
    return this;
  }

  /** Timestamp of deletion in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. */
  @javax.annotation.Nonnull
  public String getDeletedAt() {
    return deletedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeleteApiKeyResponse deleteApiKeyResponse = (DeleteApiKeyResponse) o;
    return Objects.equals(this.deletedAt, deleteApiKeyResponse.deletedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(deletedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeleteApiKeyResponse {\n");
    sb.append("    deletedAt: ").append(toIndentedString(deletedAt)).append("\n");
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
