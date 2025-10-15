import { createClient as createSupabaseClient } from "@supabase/supabase-js";
import { headers } from "next/headers";
import type { Database } from "@/types/supabase";

/**
 * Creates a Supabase client for server-side use with Telegram authentication.
 * Passes the Telegram initData from the X-Telegram-Init-Data header to Supabase.
 *
 * The initData can be validated on the Supabase side using:
 * - Edge Functions
 * - Database functions
 * - Row Level Security (RLS) policies
 *
 * For validation before reaching Supabase, use validateTelegramAuth from @/lib/telegram-auth
 */
export async function createClient() {
	const supabaseUrl = process.env.NEXT_PUBLIC_SUPABASE_URL;
	const supabaseAnonKey = process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY;

	if (!supabaseUrl || !supabaseAnonKey) {
		throw new Error(
			"Missing Supabase environment variables: NEXT_PUBLIC_SUPABASE_URL and NEXT_PUBLIC_SUPABASE_ANON_KEY",
		);
	}

	const headersList = await headers();
	const initData = headersList.get("x-telegram-init-data");

	// Create Supabase client with custom headers
	const supabase = createSupabaseClient<Database>(
		supabaseUrl,
		supabaseAnonKey,
		{
			global: {
				headers: {
					...(initData && { "x-telegram-init-data": initData }),
				},
			},
		},
	);

	return supabase;
}
