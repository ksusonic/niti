/**
 * API client utilities with automatic Telegram authentication
 */

import { initData } from "@tma.js/sdk-react";
import { TELEGRAM_INIT_DATA_HEADER } from "./constants";

/**
 * Makes an authenticated API request with Telegram init data.
 * Automatically includes the init data in request headers.
 *
 * @param url - The API endpoint URL
 * @param options - Fetch request options
 * @returns Promise resolving to the Response object
 * @throws Error if Telegram init data is not available
 *
 * @example
 * ```ts
 * const response = await authenticatedFetch('/api/events');
 * const data = await response.json();
 * ```
 */
export async function authenticatedFetch(
	url: string,
	options: RequestInit = {},
): Promise<Response> {
	const raw = initData.raw();

	if (!raw) {
		throw new Error("Telegram init data not available");
	}

	const headers = {
		...options.headers,
		[TELEGRAM_INIT_DATA_HEADER]: raw,
	};

	return fetch(url, {
		...options,
		headers,
	});
}

/**
 * Makes an authenticated JSON API request.
 * Automatically includes Telegram init data and handles JSON parsing.
 *
 * @template T - The expected response type
 * @param url - The API endpoint URL
 * @param options - Fetch request options
 * @returns Promise resolving to the typed response data
 * @throws Error if request fails or init data is unavailable
 *
 * @example
 * ```ts
 * const events = await authenticatedFetchJson<Event[]>('/api/events');
 * ```
 */
export async function authenticatedFetchJson<T = unknown>(
	url: string,
	options: RequestInit = {},
): Promise<T> {
	const response = await authenticatedFetch(url, options);

	if (!response.ok) {
		const errorData = await response.json().catch(() => ({}));
		throw new Error(
			errorData.error || `Request failed with status ${response.status}`,
		);
	}

	return response.json();
}
