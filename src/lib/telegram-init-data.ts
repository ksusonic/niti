import { retrieveLaunchParams } from "@telegram-apps/sdk-react";

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
 * Gets init data from Telegram launch params, with fallback to mock data
 * in development when the environment is mocked
 */
export function getInitData(): string {
	try {
		const launchParams = retrieveLaunchParams();
		const initData = (launchParams.initData as { raw?: string } | undefined)
			?.raw;

		if (initData) {
			return initData;
		}

		// If no init data and in development, return mock data
		if (process.env.NODE_ENV === "development") {
			console.log("[InitData] No launch params found, using mock init data");
			return MOCK_INIT_DATA;
		}

		return "";
	} catch (error) {
		console.warn("[InitData] Error retrieving launch params:", error);

		// Fallback to mock data in development
		if (process.env.NODE_ENV === "development") {
			console.log("[InitData] Using mock init data due to error");
			return MOCK_INIT_DATA;
		}

		return "";
	}
}
