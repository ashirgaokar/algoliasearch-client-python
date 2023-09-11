// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.algolia.exceptions.AlgoliaRuntimeException;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.io.IOException;
import java.util.List;
import java.util.logging.Logger;

/**
 * Names of facets to which automatic filtering must be applied; they must match the facet name of a
 * facet value placeholder in the query pattern.
 */
@JsonDeserialize(using = AutomaticFacetFilters.Deserializer.class)
public interface AutomaticFacetFilters {
  /** AutomaticFacetFilters as List<AutomaticFacetFilter> wrapper. */
  static AutomaticFacetFilters ofListOfAutomaticFacetFilter(List<AutomaticFacetFilter> value) {
    return new ListOfAutomaticFacetFilterWrapper(value);
  }

  /** AutomaticFacetFilters as List<String> wrapper. */
  static AutomaticFacetFilters ofListOfString(List<String> value) {
    return new ListOfStringWrapper(value);
  }

  /** AutomaticFacetFilters as List<AutomaticFacetFilter> wrapper. */
  @JsonSerialize(using = ListOfAutomaticFacetFilterWrapper.Serializer.class)
  class ListOfAutomaticFacetFilterWrapper implements AutomaticFacetFilters {

    private final List<AutomaticFacetFilter> value;

    ListOfAutomaticFacetFilterWrapper(List<AutomaticFacetFilter> value) {
      this.value = value;
    }

    public List<AutomaticFacetFilter> getValue() {
      return value;
    }

    static class Serializer extends JsonSerializer<ListOfAutomaticFacetFilterWrapper> {

      @Override
      public void serialize(ListOfAutomaticFacetFilterWrapper value, JsonGenerator gen, SerializerProvider provider) throws IOException {
        gen.writeObject(value.getValue());
      }
    }
  }

  /** AutomaticFacetFilters as List<String> wrapper. */
  @JsonSerialize(using = ListOfStringWrapper.Serializer.class)
  class ListOfStringWrapper implements AutomaticFacetFilters {

    private final List<String> value;

    ListOfStringWrapper(List<String> value) {
      this.value = value;
    }

    public List<String> getValue() {
      return value;
    }

    static class Serializer extends JsonSerializer<ListOfStringWrapper> {

      @Override
      public void serialize(ListOfStringWrapper value, JsonGenerator gen, SerializerProvider provider) throws IOException {
        gen.writeObject(value.getValue());
      }
    }
  }

  class Deserializer extends JsonDeserializer<AutomaticFacetFilters> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    @Override
    public AutomaticFacetFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize List<AutomaticFacetFilter>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(new TypeReference<List<AutomaticFacetFilter>>() {});
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf List<AutomaticFacetFilter> (error: " + e.getMessage() + ") (type: List<AutomaticFacetFilter>)"
          );
        }
      }

      // deserialize List<String>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(new TypeReference<List<String>>() {});
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf List<String> (error: " + e.getMessage() + ") (type: List<String>)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public AutomaticFacetFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "AutomaticFacetFilters cannot be null");
    }
  }
}
