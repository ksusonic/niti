"use client";

import type { InitData } from "@tma.js/sdk-react";
import { initData, useSignal } from "@tma.js/sdk-react";

/**
 * Hook to access Telegram init data
 * Returns the current state of init data which updates when the signal changes
 */
export function useInitData(): InitData | undefined {
	return useSignal(initData.state);
}
