// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { HighlightResult } from './highlightResult';

export type UserHighlightResult = {
  /**
   * Show highlighted section and words matched on a query.
   */
  userID: Record<string, HighlightResult>;

  /**
   * Show highlighted section and words matched on a query.
   */
  clusterName: Record<string, HighlightResult>;
};
