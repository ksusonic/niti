/**
 * API client utilities with automatic Telegram authentication
 */

import { initData } from "@telegram-apps/sdk-react";
import { TELEGRAM_INIT_DATA_HEADER } from "./constants";

/**
 * Makes an authenticated API request with Telegram init data
 * Automatically includes the init data in request headers
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
 * Makes an authenticated JSON API request
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
