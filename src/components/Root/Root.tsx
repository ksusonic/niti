"use client";

import { AppRoot } from "@telegram-apps/telegram-ui";
import { miniApp, useLaunchParams, useSignal } from "@tma.js/sdk-react";
import { type PropsWithChildren, useEffect, useState } from "react";
import { EnvUnsupported } from "@/components/EnvUnsupported";
import { ErrorBoundary } from "@/components/ErrorBoundary";
import { ErrorPage } from "@/components/ErrorPage";
import { TelegramEnvErrorBoundary } from "@/components/TelegramEnvErrorBoundary";
import { useDidMount } from "@/hooks/useDidMount";
import {
	addEnvUnsupportedListener,
	getIsEnvUnsupported,
	removeEnvUnsupportedListener,
} from "@/lib";

import "./styles.css";

function RootInner({ children }: PropsWithChildren) {
	const lp = useLaunchParams();
	const isDark = useSignal(miniApp.isDark);

	return (
		<AppRoot
			appearance={isDark ? "dark" : "light"}
			platform={["macos", "ios"].includes(lp.tgWebAppPlatform) ? "ios" : "base"}
		>
			{children}
		</AppRoot>
	);
}

export function Root(props: PropsWithChildren) {
	// Unfortunately, Telegram Mini Apps does not allow us to use all features of
	// the Server Side Rendering. That's why we are showing loader on the server
	// side.
	const didMount = useDidMount();
	const [envUnsupported, setEnvUnsupported] = useState(false);

	useEffect(() => {
		// Check initial state
		if (getIsEnvUnsupported()) {
			setEnvUnsupported(true);
		}

		// Listen for the custom event
		const handleEnvUnsupported = () => {
			setEnvUnsupported(true);
		};

		addEnvUnsupportedListener(handleEnvUnsupported);
		return () => {
			removeEnvUnsupportedListener(handleEnvUnsupported);
		};
	}, []);

	if (!didMount) {
		return <div className="root__loading">Loading</div>;
	}

	if (envUnsupported) {
		return <EnvUnsupported />;
	}

	return (
		<ErrorBoundary fallback={ErrorPage}>
			<TelegramEnvErrorBoundary>
				<RootInner {...props} />
			</TelegramEnvErrorBoundary>
		</ErrorBoundary>
	);
}
