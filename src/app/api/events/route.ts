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

	try {
		const userId = authResult.initData.user?.id;
		if (!userId) {
			return NextResponse.json(
				{ error: "User ID not found in auth data" },
				{ status: 401 },
			);
		}

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

		// Fetch all participant counts and user subscriptions in a single query
		const { data: allParticipants, error: participantsError } = await supabase
			.from("event_participants")
			.select("event_id, user_id")
			.eq("status", "going");

		if (participantsError) {
			console.error("Error fetching participants:", participantsError);
			return NextResponse.json(
				{ error: "Failed to fetch participants" },
				{ status: 500 },
			);
		}

		// Build maps for participant counts and user subscriptions
		const countMap = new Map<number, number>();
		const subscribedEventIds = new Set<number>();

		allParticipants?.forEach((record) => {
			const eventId = record.event_id as number;
			// Count participants
			countMap.set(eventId, (countMap.get(eventId) || 0) + 1);
			// Check user subscription
			if (record.user_id === userId) {
				subscribedEventIds.add(eventId);
			}
		});

		// Add participant counts to events
		const eventsWithCounts = (events || []).map((event) => ({
			...event,
			participant_count: countMap.get(event.id) || 0,
		}));

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
					isSubscribed: subscribedEventIds.has(event.id),
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
