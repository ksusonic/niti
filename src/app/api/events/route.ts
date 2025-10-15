import { NextResponse } from "next/server";
import { checkAuthHeader } from "@/lib/auth-middleware";
import { createClient } from "@/lib/supabase";
import type { Database } from "@/types/supabase";

type EventRow = Database["public"]["Tables"]["events"]["Row"];
type EventLineupRow = Database["public"]["Tables"]["event_lineup"]["Row"];
type ProfileRow = Database["public"]["Tables"]["profiles"]["Row"];

type EventLineupWithProfile = EventLineupRow & {
	profiles: ProfileRow | null;
};

interface EventWithLineup extends EventRow {
	event_lineup: EventLineupWithProfile[];
	participant_count: number;
}

export async function GET() {
	const authResult = await checkAuthHeader();
	if (authResult instanceof NextResponse) {
		return authResult;
	}

	const { initData } = authResult;
	console.debug("authenticated user:", initData.user);

	try {
		const supabase = await createClient();

		// Fetch events with lineup and participant count
		const { data: events, error } = await supabase
			.from("events")
			.select(`*, event_lineup (*, profiles (*))`)
			.order("start_time", { ascending: true });

		if (error) {
			console.error("Error fetching events:", error);
			return NextResponse.json(
				{ error: "Failed to fetch events" },
				{ status: 500 },
			);
		}

		// Get participant counts for each event
		const eventsWithCounts = await Promise.all(
			(events || []).map(async (event) => {
				const { count } = await supabase
					.from("event_participants")
					.select("*", { count: "exact", head: true })
					.eq("event_id", event.id)
					.eq("status", "going");

				return {
					...event,
					participant_count: count || 0,
				};
			}),
		);

		// Transform to match frontend Event type
		const transformedEvents = (eventsWithCounts as EventWithLineup[]).map(
			(event) => {
				const startTime = new Date(event.start_time);
				const dateStr = startTime.toLocaleDateString("ru-RU", {
					month: "short",
					day: "numeric",
				});
				const timeStr = startTime.toLocaleTimeString("ru-RU", {
					hour: "numeric",
					minute: "2-digit",
					hour12: true,
				});

				return {
					id: event.id.toString(),
					title: event.title,
					description: event.description || "",
					location: event.location || "",
					imageUrl: event.banner_url || "",
					videoUrl: undefined, // TODO: Add video_url field to schema if needed
					djLineup:
						event.event_lineup?.map((lineup) => ({
							id: lineup.dj_id?.toString() || "",
							name:
								lineup.profiles?.display_name ||
								lineup.profiles?.username ||
								"",
							avatar: lineup.profiles?.avatar_url || "",
							time: `${new Date(lineup.start_time).toLocaleTimeString("en-US", {
								hour: "numeric",
								minute: "2-digit",
								hour12: true,
							})}${
								lineup.end_time
									? ` - ${new Date(lineup.end_time).toLocaleTimeString(
											"en-US",
											{
												hour: "numeric",
												minute: "2-digit",
												hour12: true,
											},
										)}`
									: ""
							}`,
							social: {
								instagram: (
									lineup.profiles?.social_links as { instagram?: string }
								)?.instagram,
								soundcloud: (
									lineup.profiles?.social_links as { soundcloud?: string }
								)?.soundcloud,
								spotify: (lineup.profiles?.social_links as { spotify?: string })
									?.spotify,
							},
						})) || [],
					participantCount: event.participant_count,
					isSubscribed: false, // TODO: Check if current user is subscribed
					date: dateStr,
					time: timeStr,
				};
			},
		);

		return NextResponse.json(transformedEvents);
	} catch (error) {
		console.error("Unexpected error:", error);
		return NextResponse.json(
			{ error: "Internal server error" },
			{ status: 500 },
		);
	}
}
