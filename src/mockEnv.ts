import { emitEvent, isTMA, mockTelegramEnv } from "@telegram-apps/sdk-react";
import {
	getMockLaunchParams,
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
		launchParams: getMockLaunchParams(),
	});

	console.info(
		"⚠️ As long as the current environment was not considered as the Telegram-based one, it was mocked. Take a note, that you should not do it in production and current behavior is only specific to the development process. Environment mocking is also applied only in development mode. So, after building the application, you will not see this behavior and related warning, leading to crashing the application outside Telegram.",
	);
}
