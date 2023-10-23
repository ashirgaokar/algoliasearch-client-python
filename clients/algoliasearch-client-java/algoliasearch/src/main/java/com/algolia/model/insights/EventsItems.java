// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.insights;

import com.algolia.exceptions.AlgoliaRuntimeException;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.io.IOException;
import java.util.logging.Logger;

/** EventsItems */
@JsonDeserialize(using = EventsItems.Deserializer.class)
public interface EventsItems {
  class Deserializer extends JsonDeserializer<EventsItems> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    @Override
    public EventsItems deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize AddedToCartObjectIDs
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(AddedToCartObjectIDs.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf AddedToCartObjectIDs (error: " + e.getMessage() + ") (type: AddedToCartObjectIDs)");
        }
      }

      // deserialize AddedToCartObjectIDsAfterSearch
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(AddedToCartObjectIDsAfterSearch.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf AddedToCartObjectIDsAfterSearch (error: " +
            e.getMessage() +
            ") (type: AddedToCartObjectIDsAfterSearch)"
          );
        }
      }

      // deserialize ClickedFilters
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ClickedFilters.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ClickedFilters (error: " + e.getMessage() + ") (type: ClickedFilters)");
        }
      }

      // deserialize ClickedObjectIDs
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ClickedObjectIDs.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ClickedObjectIDs (error: " + e.getMessage() + ") (type: ClickedObjectIDs)");
        }
      }

      // deserialize ClickedObjectIDsAfterSearch
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ClickedObjectIDsAfterSearch.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf ClickedObjectIDsAfterSearch (error: " + e.getMessage() + ") (type: ClickedObjectIDsAfterSearch)"
          );
        }
      }

      // deserialize ConvertedFilters
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ConvertedFilters.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ConvertedFilters (error: " + e.getMessage() + ") (type: ConvertedFilters)");
        }
      }

      // deserialize ConvertedObjectIDs
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ConvertedObjectIDs.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ConvertedObjectIDs (error: " + e.getMessage() + ") (type: ConvertedObjectIDs)");
        }
      }

      // deserialize ConvertedObjectIDsAfterSearch
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ConvertedObjectIDsAfterSearch.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf ConvertedObjectIDsAfterSearch (error: " +
            e.getMessage() +
            ") (type: ConvertedObjectIDsAfterSearch)"
          );
        }
      }

      // deserialize Identify
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(Identify.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf Identify (error: " + e.getMessage() + ") (type: Identify)");
        }
      }

      // deserialize PurchasedObjectIDs
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(PurchasedObjectIDs.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf PurchasedObjectIDs (error: " + e.getMessage() + ") (type: PurchasedObjectIDs)");
        }
      }

      // deserialize PurchasedObjectIDsAfterSearch
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(PurchasedObjectIDsAfterSearch.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf PurchasedObjectIDsAfterSearch (error: " +
            e.getMessage() +
            ") (type: PurchasedObjectIDsAfterSearch)"
          );
        }
      }

      // deserialize ViewedFilters
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ViewedFilters.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ViewedFilters (error: " + e.getMessage() + ") (type: ViewedFilters)");
        }
      }

      // deserialize ViewedObjectIDs
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          return parser.readValueAs(ViewedObjectIDs.class);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf ViewedObjectIDs (error: " + e.getMessage() + ") (type: ViewedObjectIDs)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public EventsItems getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "EventsItems cannot be null");
    }
  }
}
