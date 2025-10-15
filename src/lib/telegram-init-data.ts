import { initData } from "@telegram-apps/sdk-react";

/**
 * Mock init data for development when environment is mocked
 * This should match the data structure from mockEnv.ts
 */
export const MOCK_INIT_DATA = new URLSearchParams([
	["auth_date", ((Date.now() / 1000) | 0).toString()],
	["hash", "some-hash"],
	["signature", "mock-signature"],
	["user", JSON.stringify({ id: 1, first_name: "Ksusonic" })],
]).toString();

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

		if (process.env.NODE_ENV === "development") {
			console.log("[InitData] No init data found, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	} catch (error) {
		console.error("[InitData] Error retrieving init data:", error);
		if (process.env.NODE_ENV === "development") {
			console.log("[InitData] Error occurred, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	}
}
