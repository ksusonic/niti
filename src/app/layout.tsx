import type { Metadata } from "next";
import type { PropsWithChildren } from "react";

import { Root } from "@/components/Root/Root";

import "@telegram-apps/telegram-ui/dist/styles.css";
import "normalize.css/normalize.css";
import "./globals.css";

export const metadata: Metadata = {
	title: "NITI APP",
	description: "Приложение для вечеринок команды NITI TEAM",
};

export default function RootLayout({ children }: PropsWithChildren) {
	return (
		<html lang="ru" suppressHydrationWarning>
			<body>
				<Root>{children}</Root>
				{/* CSS-only loading fallback for initial page load */}
				<style
					// biome-ignore lint/security/noDangerouslySetInnerHtml: Initial loading animation
					dangerouslySetInnerHTML={{
						__html: `
							#__next:empty::before,
							body > div:first-child:empty::before {
								content: '';
								position: fixed;
								top: 50%;
								left: 50%;
								transform: translate(-50%, -50%);
								width: 60px;
								height: 60px;
								border: 4px solid rgba(168, 85, 247, 0.3);
								border-top-color: rgb(168, 85, 247);
								border-radius: 50%;
								animation: spin 1s linear infinite;
								z-index: 9999;
							}
							
							#__next:empty::after,
							body > div:first-child:empty::after {
								content: 'Spinning Up The Decks...';
								position: fixed;
								top: 60%;
								left: 50%;
								transform: translateX(-50%);
								color: rgb(168, 85, 247);
								font-size: 1.125rem;
								font-weight: 600;
								z-index: 9999;
								animation: pulse 1.5s ease-in-out infinite;
							}
						`,
					}}
				/>
			</body>
		</html>
	);
}
