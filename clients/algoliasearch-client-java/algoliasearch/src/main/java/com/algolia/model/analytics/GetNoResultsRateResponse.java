// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetNoResultsRateResponse */
public class GetNoResultsRateResponse {

  @JsonProperty("rate")
  private Double rate;

  @JsonProperty("count")
  private Integer count;

  @JsonProperty("noResultCount")
  private Integer noResultCount;

  @JsonProperty("dates")
  private List<NoResultsRateEvent> dates = new ArrayList<>();

  public GetNoResultsRateResponse setRate(Double rate) {
    this.rate = rate;
    return this;
  }

  /**
   * [Click-through rate
   * (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
   * minimum: 0 maximum: 1
   */
  @javax.annotation.Nonnull
  public Double getRate() {
    return rate;
  }

  public GetNoResultsRateResponse setCount(Integer count) {
    this.count = count;
    return this;
  }

  /** Number of occurrences. */
  @javax.annotation.Nonnull
  public Integer getCount() {
    return count;
  }

  public GetNoResultsRateResponse setNoResultCount(Integer noResultCount) {
    this.noResultCount = noResultCount;
    return this;
  }

  /** Number of occurrences. */
  @javax.annotation.Nonnull
  public Integer getNoResultCount() {
    return noResultCount;
  }

  public GetNoResultsRateResponse setDates(List<NoResultsRateEvent> dates) {
    this.dates = dates;
    return this;
  }

  public GetNoResultsRateResponse addDates(NoResultsRateEvent datesItem) {
    this.dates.add(datesItem);
    return this;
  }

  /** Overall count of searches without results plus a daily breakdown. */
  @javax.annotation.Nonnull
  public List<NoResultsRateEvent> getDates() {
    return dates;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetNoResultsRateResponse getNoResultsRateResponse = (GetNoResultsRateResponse) o;
    return (
      Objects.equals(this.rate, getNoResultsRateResponse.rate) &&
      Objects.equals(this.count, getNoResultsRateResponse.count) &&
      Objects.equals(this.noResultCount, getNoResultsRateResponse.noResultCount) &&
      Objects.equals(this.dates, getNoResultsRateResponse.dates)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(rate, count, noResultCount, dates);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetNoResultsRateResponse {\n");
    sb.append("    rate: ").append(toIndentedString(rate)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    noResultCount: ").append(toIndentedString(noResultCount)).append("\n");
    sb.append("    dates: ").append(toIndentedString(dates)).append("\n");
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
