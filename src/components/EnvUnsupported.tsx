import { AppRoot, Placeholder } from "@telegram-apps/telegram-ui";
import { isColorDark, isRGB, retrieveLaunchParams } from "@tma.js/sdk-react";
import Image from "next/image";
import { useMemo } from "react";

export function EnvUnsupported() {
	const [platform, isDark] = useMemo(() => {
		try {
			const lp = retrieveLaunchParams();
			const { bg_color: bgColor } = lp.tgWebAppThemeParams;
			return [
				lp.tgWebAppPlatform,
				bgColor && isRGB(bgColor) ? isColorDark(bgColor) : false,
			];
		} catch {
			return ["android", false];
		}
	}, []);

	return (
		<AppRoot
			appearance={isDark ? "dark" : "light"}
			platform={["macos", "ios"].includes(platform) ? "ios" : "base"}
		>
			<Placeholder
				header="Упс..."
				description="Это приложение работает только внутри Telegram"
			>
				<Image
					alt="Telegram sticker"
					src="https://xelene.me/telegram.gif"
					width={144}
					height={144}
					style={{ display: "block" }}
					unoptimized
				/>
			</Placeholder>
		</AppRoot>
	);
}
