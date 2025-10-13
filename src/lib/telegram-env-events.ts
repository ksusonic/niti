let isEnvUnsupported = false;

export const TELEGRAM_ENV_UNSUPPORTED_EVENT = "telegram-env-unsupported";

export function getIsEnvUnsupported(): boolean {
	return isEnvUnsupported;
}

export function setEnvUnsupported(): void {
	if (!isEnvUnsupported) {
		console.warn(
			"Environment marked as unsupported - switching to fallback UI",
		);
		isEnvUnsupported = true;
		if (typeof window !== "undefined") {
			window.dispatchEvent(new CustomEvent(TELEGRAM_ENV_UNSUPPORTED_EVENT));
		}
	}
}

export function addEnvUnsupportedListener(callback: () => void): void {
	if (typeof window !== "undefined") {
		window.addEventListener(TELEGRAM_ENV_UNSUPPORTED_EVENT, callback);
	}
}

export function removeEnvUnsupportedListener(callback: () => void): void {
	if (typeof window !== "undefined") {
		window.removeEventListener(TELEGRAM_ENV_UNSUPPORTED_EVENT, callback);
	}
}
