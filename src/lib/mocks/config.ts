/**
 * Centralized mock configuration for development environment
 * Uses proper types compatible with @tma.js/sdk-react
 */

/**
 * Mock user data matching Telegram SDK User type
 * Properties are in snake_case to match the SDK
 */
export const MOCK_USER = {
	id: 1,
	username: "ryangosling",
	first_name: "Ryan",
	last_name: "Gosling",
	photo_url: "/ryan-gosling.jpg",
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
 * Magic hash value that identifies mock data during development
 */
export const MOCK_HASH_IDENTIFIER = "some-hash";

/**
 * Launch parameters for the mocked Telegram environment
 * Compatible with mockTelegramEnv from @tma.js/sdk-react
 */
export const MOCK_LAUNCH_PARAMS = new URLSearchParams([
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
