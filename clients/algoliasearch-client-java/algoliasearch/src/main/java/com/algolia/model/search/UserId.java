// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** Unique user ID. */
public class UserId {

  @JsonProperty("userID")
  private String userID;

  @JsonProperty("clusterName")
  private String clusterName;

  @JsonProperty("nbRecords")
  private Integer nbRecords;

  @JsonProperty("dataSize")
  private Integer dataSize;

  public UserId setUserID(String userID) {
    this.userID = userID;
    return this;
  }

  /** userID of the user. */
  @javax.annotation.Nonnull
  public String getUserID() {
    return userID;
  }

  public UserId setClusterName(String clusterName) {
    this.clusterName = clusterName;
    return this;
  }

  /** Cluster to which the user is assigned. */
  @javax.annotation.Nonnull
  public String getClusterName() {
    return clusterName;
  }

  public UserId setNbRecords(Integer nbRecords) {
    this.nbRecords = nbRecords;
    return this;
  }

  /** Number of records belonging to the user. */
  @javax.annotation.Nonnull
  public Integer getNbRecords() {
    return nbRecords;
  }

  public UserId setDataSize(Integer dataSize) {
    this.dataSize = dataSize;
    return this;
  }

  /** Data size used by the user. */
  @javax.annotation.Nonnull
  public Integer getDataSize() {
    return dataSize;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserId userId = (UserId) o;
    return (
      Objects.equals(this.userID, userId.userID) &&
      Objects.equals(this.clusterName, userId.clusterName) &&
      Objects.equals(this.nbRecords, userId.nbRecords) &&
      Objects.equals(this.dataSize, userId.dataSize)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(userID, clusterName, nbRecords, dataSize);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserId {\n");
    sb.append("    userID: ").append(toIndentedString(userID)).append("\n");
    sb.append("    clusterName: ").append(toIndentedString(clusterName)).append("\n");
    sb.append("    nbRecords: ").append(toIndentedString(nbRecords)).append("\n");
    sb.append("    dataSize: ").append(toIndentedString(dataSize)).append("\n");
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
