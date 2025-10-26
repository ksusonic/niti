import {
	backButton,
	emitEvent,
	init as initSDK,
	isTMA,
	miniApp,
	mockTelegramEnv,
	retrieveLaunchParams,
	setDebug,
	type ThemeParams,
	themeParams,
	viewport,
} from "@tma.js/sdk-react";

/**
 * Initializes the application and configures its dependencies.
 */
export async function init(options: {
	debug: boolean;
	mockForMacOS: boolean;
}): Promise<void> {
	try {
		setDebug(options.debug);
		initSDK();

		// Telegram for macOS has a ton of bugs, including cases, when the client doesn't
		// even response to the "web_app_request_theme" method. It also generates an incorrect
		// event for the "web_app_request_safe_area" method.
		if (options.mockForMacOS) {
			let firstThemeSent = false;
			mockTelegramEnv({
				onEvent(event, next) {
					if (event.name === "web_app_request_theme") {
						let tp: ThemeParams = {};
						if (firstThemeSent) {
							tp = themeParams.state();
						} else {
							firstThemeSent = true;
							tp ||= retrieveLaunchParams().tgWebAppThemeParams;
						}
						return emitEvent("theme_changed", { theme_params: tp });
					}

					if (event.name === "web_app_request_safe_area") {
						return emitEvent("safe_area_changed", {
							left: 0,
							top: 0,
							right: 0,
							bottom: 0,
						});
					}

					next();
				},
			});
		}

		const isTelegramEnv = await isTMA("complete").catch((error) => {
			console.warn("Failed to detect Telegram environment in init:", error);
			return false;
		});

		if (isTelegramEnv) {
			try {
				// Mount back button if available
				if (backButton.mount.isAvailable()) {
					backButton.mount();
				}
			} catch (error) {
				console.warn("Failed to mount back button:", error);
			}

			try {
				// Mount theme params first
				if (themeParams.mount.isAvailable()) {
					themeParams.mount();
				}
			} catch (error) {
				console.warn("Failed to mount theme params:", error);
			}

			try {
				// Mount mini app (requires theme params)
				if (miniApp.mount.isAvailable()) {
					miniApp.mount();
					// Bind CSS variables for theming
					themeParams.bindCssVars();
				}
			} catch (error) {
				console.warn("Failed to mount mini app or bind theme params:", error);
			}

			try {
				if (viewport.mount.isAvailable()) {
					viewport
						.mount()
						.then(() => {
							// Bind viewport CSS variables
							viewport.bindCssVars();
						})
						.catch((error: Error) => {
							console.warn("Failed to bind viewport CSS vars:", error);
						});
				}
			} catch (error) {
				console.warn("Failed to mount viewport:", error);
			}
		}
	} catch (e) {
		console.error("Critical error during Telegram SDK initialization:", e);
		throw e;
	}
}
