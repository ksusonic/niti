"use server";

import { parse, validate } from "@tma.js/init-data-node";
import type { InitDataParsed } from "@tma.js/sdk";
import { headers } from "next/headers";
import { NextResponse } from "next/server";
import { TELEGRAM_INIT_DATA_HEADER } from "./constants";
import { isDevelopmentEnv, isMockInitData } from "./mocks/utils";

export type AuthResult = { initData: InitDataParsed } | NextResponse;

/**
 * Checks for Telegram auth header in the request, validates it with bot token, and parses init data
 * Returns parsed init data if valid, or a NextResponse error if missing/invalid
 */
export async function checkAuthHeader(): Promise<AuthResult> {
	const headersList = await headers();
	const rawInitData = headersList.get(TELEGRAM_INIT_DATA_HEADER);

	console.debug(
		"[Auth] Raw init data received:",
		rawInitData ? "✓ present" : "✗ missing",
	);

	if (!rawInitData) {
		console.debug("[Auth] Auth failed: Missing init data header");
		return NextResponse.json(
			{ error: "Telegram authorization required" },
			{ status: 401 },
		);
	}

	try {
		// In development with mock data, skip validation
		if (isDevelopmentEnv() && isMockInitData(rawInitData)) {
			console.debug(
				"[Auth] Development mode: skipping validation for mock data",
			);
			const initData = parse(rawInitData);
			console.debug("[Auth] Parsed mock initData", {
				userId: initData.user?.id,
				firstName: initData.user?.firstName,
			});
			return { initData };
		}

		// Production or real data: validate with bot token
		const botToken = process.env.TELEGRAM_BOT_TOKEN;

		if (!botToken) {
			console.error(
				"[Auth] Failed: TELEGRAM_BOT_TOKEN environment variable is not set",
			);
			return NextResponse.json(
				{ error: "Server configuration error" },
				{ status: 500 },
			);
		}
		console.debug("[Auth] Bot token found ✓");

		validate(rawInitData, botToken, {
			expiresIn: 86400, // 24 hours
		});

		const initData = parse(rawInitData);
		console.debug("[Auth] Parsed initData", {
			userId: initData.user?.id,
			username: initData.user?.username,
		});

		return { initData };
	} catch (error) {
		console.error("[Auth] Failed to validate/parse init data:", error);

		return NextResponse.json(
			{ error: "Invalid Telegram authorization data" },
			{ status: 401 },
		);
	}
}
