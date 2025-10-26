"use client";

import { type PropsWithChildren, useEffect } from "react";

interface TelegramThemeProviderProps extends PropsWithChildren {
	appearance: "light" | "dark";
	platform: "ios" | "base";
}

/**
 * Custom theme provider replacing @telegram-apps/telegram-ui AppRoot
 * Provides CSS variables for theming based on Telegram Mini Apps theme
 */
export function TelegramThemeProvider({
	children,
	appearance,
	platform,
}: TelegramThemeProviderProps) {
	useEffect(() => {
		// Set theme class on document
		document.documentElement.classList.remove("light", "dark");
		document.documentElement.classList.add(appearance);

		// Set platform data attribute
		document.documentElement.setAttribute("data-platform", platform);
	}, [appearance, platform]);

	return (
		<div
			className="telegram-theme-root"
			data-appearance={appearance}
			data-platform={platform}
		>
			{children}
		</div>
	);
}
