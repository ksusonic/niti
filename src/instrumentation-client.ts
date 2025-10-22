// This file is normally used for setting up analytics and other
// services that require one-time initialization on the client.

import {
	isLaunchParamsRetrieveError,
	isTMA,
	retrieveLaunchParams,
} from "@telegram-apps/sdk-react";
import { init } from "./core/init";
import { setEnvUnsupported } from "./lib";
import { mockEnv } from "./mockEnv";

async function checkTelegramEnvAndInitialize() {
	const isTelegramEnv = await isTMA("complete");

	if (!isTelegramEnv) {
		setEnvUnsupported();
		return;
	}

	try {
		const launchParams = retrieveLaunchParams();
		const { tgWebAppPlatform: platform } = launchParams;
		const debug = process.env.NODE_ENV === "development";

		init({
			debug,
			mockForMacOS: platform === "macos",
		});
	} catch (e) {
		if (isLaunchParamsRetrieveError(e)) {
			console.warn(
				"LaunchParamsRetrieveError: App running outside Telegram environment",
			);
			setEnvUnsupported();
		} else {
			console.error("Unexpected error during launch params retrieval:", e);
		}
	}
}

if (typeof window !== "undefined") {
	mockEnv()
		.then(() => {
			// After mocking setup, check if we're in Telegram environment
			return checkTelegramEnvAndInitialize();
		})
		.catch((error) => {
			console.warn("Mock environment setup failed:", error);
			// If mocking setup fails, try to check if we're in real Telegram environment
			return checkTelegramEnvAndInitialize();
		});
}
