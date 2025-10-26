/**
 * Centralized mock configuration for development environment
 * Uses proper types compatible with @telegram-apps/sdk-react
 */

/**
 * Mock user data matching Telegram SDK User type
 * Properties are in snake_case to match the SDK
 */
export const MOCK_USER = {
	id: 1,
	username: "ksusonic",
	first_name: "Daniil",
	last_name: "Dev",
	language_code: "en",
	is_premium: false,
	allows_write_to_pm: true,
} as const;

/**
 * Mock theme parameters for Telegram UI theming
 */
export const MOCK_THEME_PARAMS = {
	accent_text_color: "#6ab2f2",
	bg_color: "#17212b",
	button_color: "#5288c1",
	button_text_color: "#ffffff",
	destructive_text_color: "#ec3942",
	header_bg_color: "#17212b",
	hint_color: "#708499",
	link_color: "#6ab3f3",
	secondary_bg_color: "#232e3c",
	section_bg_color: "#17212b",
	section_header_text_color: "#6ab3f3",
	subtitle_text_color: "#708499",
	text_color: "#f5f5f5",
} as const;

/**
 * Mock safe area insets
 */
export const MOCK_SAFE_AREA = {
	left: 0,
	top: 0,
	right: 0,
	bottom: 0,
} as const;

/**
 * Mock viewport configuration
 */
export const MOCK_VIEWPORT = {
	height: typeof window !== "undefined" ? window.innerHeight : 800,
	width: typeof window !== "undefined" ? window.innerWidth : 600,
	is_expanded: true,
	is_state_stable: true,
} as const;

/**
 * Mock init data for development when environment is mocked
 * This creates a properly formatted init data string that matches what Telegram would send
 */
export const MOCK_INIT_DATA = new URLSearchParams([
	["auth_date", ((Date.now() / 1000) | 0).toString()],
	["hash", "some-hash"],
	["signature", "some-signature"],
	["user", JSON.stringify(MOCK_USER)],
]).toString();

/**
 * Magic hash value that identifies mock data during development
 */
export const MOCK_HASH_IDENTIFIER = "some-hash";

/**
 * Launch parameters for the mocked Telegram environment
 * Compatible with mockTelegramEnv from @telegram-apps/sdk-react
 */
export function getMockLaunchParams(): URLSearchParams {
	return new URLSearchParams([
		["tgWebAppThemeParams", JSON.stringify(MOCK_THEME_PARAMS)],
		[
			"tgWebAppData",
			new URLSearchParams([
				["auth_date", ((Date.now() / 1000) | 0).toString()],
				["hash", MOCK_HASH_IDENTIFIER],
				["signature", "some-signature"],
				["user", JSON.stringify(MOCK_USER)],
			]).toString(),
		],
		["tgWebAppVersion", "8.4"],
		["tgWebAppPlatform", "tdesktop"],
	]);
}
