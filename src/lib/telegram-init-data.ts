import { initData } from "@telegram-apps/sdk-react";
import { MOCK_INIT_DATA } from "./mocks/config";
import { isDevelopmentEnv } from "./mocks/utils";

/**
 * Gets init data from Telegram SDK (after restoreInitData() has been called),
 * with fallback to mock data in development when the environment is mocked
 */
export function getInitData(): string {
	try {
		const raw = initData.raw();
		if (raw) {
			return raw;
		}

		if (isDevelopmentEnv()) {
			console.log("[InitData] No init data found, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	} catch (error) {
		console.error("[InitData] Error retrieving init data:", error);
		if (isDevelopmentEnv()) {
			console.log("[InitData] Error occurred, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	}
}
