/**
 * Utilities for detecting and managing mock environment state
 */

/**
 * Checks if mocks are enabled via environment variable
 */
export function areMocksEnabled(): boolean {
	return process.env.NEXT_PUBLIC_ENABLE_MOCKS === "true";
}

/**
 * Checks if we're in development environment
 */
export function isDevelopmentEnv(): boolean {
	return process.env.NODE_ENV === "development";
}

/**
 * Checks if the given init data is mock data
 * Mock data is identified by the presence of the magic mock hash
 */
export function isMockInitData(rawInitData: string): boolean {
	return rawInitData.includes("hash=some-hash");
}

/**
 * Determines if mock initialization should be used
 * Returns true if we're in development and mocks are enabled
 */
export function shouldUseMocks(): boolean {
	return isDevelopmentEnv() && areMocksEnabled();
}
