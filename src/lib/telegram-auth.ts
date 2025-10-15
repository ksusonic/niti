"use server";

import { validate } from "@tma.js/init-data-node";
import { headers } from "next/headers";

/**
 * Validates Telegram auth data and returns parsed user info
 */
export async function validateTelegramAuth(): Promise<{
	valid: boolean;
	user?: Record<string, unknown>;
	error?: string;
}> {
	const headersList = await headers();
	const initData = headersList.get("x-telegram-init-data");

	if (!initData) {
		return { valid: false, error: "Missing init data" };
	}

	try {
		const botToken = process.env.TELEGRAM_BOT_TOKEN;
		if (!botToken) {
			return { valid: false, error: "Server configuration error" };
		}

		validate(initData, botToken, {
			expiresIn: 86400, // 24 hours
		});

		// Parse user data from init data
		const searchParams = new URLSearchParams(initData);
		const userStr = searchParams.get("user");
		const user = userStr ? JSON.parse(userStr) : undefined;

		return { valid: true, user };
	} catch (error) {
		console.error("Telegram auth validation failed:", error);
		return { valid: false, error: "Invalid authentication data" };
	}
}

/**
 * Gets the Telegram init data from the current request
 */
export async function getTelegramInitData(): Promise<string | null> {
	const headersList = await headers();
	return headersList.get("x-telegram-init-data");
}

/**
 * Example of using Telegram auth in an API route
 *
 * @example
 * export async function GET(request: Request) {
 *   if (!isTelegramAuthValid()) {
 *     return Response.json({ error: "Unauthorized" }, { status: 401 });
 *   }
 *
 *   const user = getParsedTelegramUser();
 *   // Use user data...
 * }
 */
