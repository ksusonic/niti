"use client";

import { initData, useSignal } from "@telegram-apps/sdk-react";

export function useInitData() {
	const initDataState = useSignal(initData.state);

	return initDataState;
}
