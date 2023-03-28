// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Pagination } from './pagination';
import type { Task } from './task';

/**
 * A list of tasks with pagination details.
 */
export type ListTasksResponse = {
  tasks: Task[];

  pagination: Pagination;
};
