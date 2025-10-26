import { emitEvent, isTMA, mockTelegramEnv } from "@telegram-apps/sdk-react";
import {
	MOCK_LAUNCH_PARAMS,
	MOCK_SAFE_AREA,
	MOCK_THEME_PARAMS,
	MOCK_VIEWPORT,
} from "./lib/mocks/config";
import { areMocksEnabled } from "./lib/mocks/utils";

export async function mockEnv(): Promise<void> {
	if (!areMocksEnabled()) {
		return;
	}

	const isTma = await isTMA("complete");

	if (isTma) {
		return;
	}

	mockTelegramEnv({
		onEvent(e) {
			if (e[0] === "web_app_request_theme") {
				return emitEvent("theme_changed", {
					theme_params: MOCK_THEME_PARAMS,
				});
			}

			if (e[0] === "web_app_request_viewport") {
				return emitEvent("viewport_changed", MOCK_VIEWPORT);
			}

			if (e[0] === "web_app_request_content_safe_area") {
				return emitEvent("content_safe_area_changed", MOCK_SAFE_AREA);
			}

			if (e[0] === "web_app_request_safe_area") {
				return emitEvent("safe_area_changed", MOCK_SAFE_AREA);
			}
		},
		launchParams: MOCK_LAUNCH_PARAMS,
	});

	console.info(
		"⚠️ Mock Telegram Environment Initialized",
		"\n\nThe app is running with a mocked Telegram environment because:",
		"\n1. NEXT_PUBLIC_ENABLE_MOCKS=true is set",
		"\n2. The app is not running inside Telegram",
		"\n\nThis allows local development and testing.",
		"\n\n⚠️ IMPORTANT: This behavior is ONLY for development.",
		"\nIn production builds, mocking is automatically disabled,",
		"\nand the app will only work inside Telegram."
	);
}
