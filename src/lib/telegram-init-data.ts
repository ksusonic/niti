import { initData } from "@telegram-apps/sdk-react";
import { MOCK_INIT_DATA } from "./mocks/config";
import { isDevelopmentEnv } from "./mocks/utils";

/**
 * Gets raw init data string from Telegram SDK
 * This is used for API authentication headers
 *
 * @returns The raw initData string, mock data in development, or empty string if unavailable
 */
export function getInitData(): string {
	try {
		// Try to get raw init data from SDK
		const raw = initData.raw();
		if (raw) {
			return raw;
		}

		// Fallback to mock data in development
		if (isDevelopmentEnv()) {
			console.log("[InitData] No init data found, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	} catch (error) {
		console.error("[InitData] Error retrieving init data:", error);

		// Fallback to mock data in development on error
		if (isDevelopmentEnv()) {
			console.log("[InitData] Error occurred, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	}
}
