import { createClient as createSupabaseClient } from "@supabase/supabase-js";
import type { Database } from "@/types/supabase";

/**
 * Creates a Supabase client for server-side use with service role key.
 * This client has full access to the database, bypassing RLS policies.
 *
 * ⚠️ IMPORTANT: Only use this in secure backend environments!
 * Never expose the service role key to the client/browser.
 */
export function createAdminClient() {
	const supabaseUrl = process.env.NEXT_PUBLIC_SUPABASE_URL;
	const serviceRoleKey = process.env.SUPABASE_SERVICE_ROLE_KEY;

	if (!supabaseUrl || !serviceRoleKey) {
		throw new Error(
			"Missing Supabase environment variables: NEXT_PUBLIC_SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY",
		);
	}

	// Create admin client with service role key
	// Disable session persistence since this is server-side only
	const supabase = createSupabaseClient<Database>(supabaseUrl, serviceRoleKey, {
		auth: {
			persistSession: false,
			autoRefreshToken: false,
			detectSessionInUrl: false,
		},
	});

	return supabase;
}
