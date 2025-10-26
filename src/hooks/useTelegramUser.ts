/**
 * Typed React hooks for accessing Telegram Mini Apps user data and init data
 *
 * These hooks provide type-safe access to Telegram user information and init data
 * using the @telegram-apps/sdk-react package. They handle SDK initialization
 * and provide typed return values.
 *
 * @module useTelegramUser
 * @example
 * ```tsx
 * import { useTelegramUser, useInitDataRaw } from '@/hooks/useTelegramUser';
 *
 * function MyComponent() {
 *   const user = useTelegramUser();
 *   const initDataRaw = useInitDataRaw();
 *
 *   if (!user) return <div>Loading...</div>;
 *
 *   return (
 *     <div>
 *       <p>Hello, {user.first_name}!</p>
 *       <p>Username: {user.username}</p>
 *     </div>
 *   );
 * }
 * ```
 */

"use client";

import type { User } from "@telegram-apps/sdk-react";
import { initData, useSignal } from "@telegram-apps/sdk-react";
import { useEffect, useState } from "react";

/**
 * Hook to access typed Telegram user data from initData
 *
 * Returns the user object with proper typing from @telegram-apps/sdk-react.
 * The user object contains properties in snake_case format (e.g., first_name, last_name, photo_url)
 *
 * @returns {User | undefined} Telegram user object or undefined if not available
 *
 * @example
 * ```tsx
 * const user = useTelegramUser();
 *
 * if (user) {
 *   console.log(user.first_name); // "John"
 *   console.log(user.last_name);  // "Doe"
 *   console.log(user.username);   // "johndoe"
 *   console.log(user.photo_url);  // "https://..."
 * }
 * ```
 */
export function useTelegramUser(): User | undefined {
	const initDataState = useSignal(initData.state);

	return initDataState?.user;
}

/**
 * Hook to access full typed initData state
 *
 * Returns the complete initData state with proper typing, including user, hash,
 * auth_date, and other Telegram-specific properties
 *
 * @returns {InitData | undefined} Full init data state or undefined if not available
 *
 * @example
 * ```tsx
 * const data = useInitData();
 *
 * if (data) {
 *   console.log(data.user);      // User object
 *   console.log(data.auth_date); // Date object
 *   console.log(data.hash);      // Authentication hash
 * }
 * ```
 */
export function useInitData() {
	const initDataState = useSignal(initData.state);

	return initDataState;
}

/**
 * Hook to get raw initData string for API requests
 *
 * Returns the raw initData string that should be sent to your backend API
 * for authentication. This string includes the user data, hash, and signature.
 *
 * @returns {string} Raw initData string or empty string if not available
 *
 * @example
 * ```tsx
 * const initDataRaw = useInitDataRaw();
 *
 * // Use in API requests for authentication
 * const response = await fetch('/api/events', {
 *   headers: {
 *     'X-Telegram-Init-Data': initDataRaw,
 *   },
 * });
 * ```
 */
export function useInitDataRaw(): string {
	const [rawData, setRawData] = useState<string>("");

	useEffect(() => {
		try {
			const raw = initData.raw();
			if (raw) {
				setRawData(raw);
			}
		} catch (error) {
			console.error("[useInitDataRaw] Error getting raw init data:", error);
		}
	}, []);

	return rawData;
}
