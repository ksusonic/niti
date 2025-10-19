import { NextResponse } from "next/server";
import { checkAuthHeader } from "@/lib/auth-middleware";
import { createClient } from "@/lib/supabase";

interface SubscriptionEvent {
	id: string;
	title: string;
	date: string;
	location: string;
	imageUrl: string;
	startTime: string;
}

export async function GET(request: Request) {
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

		// Parse query parameter: "includePast" to show past events
		const { searchParams } = new URL(request.url);
		const includePast = searchParams.get("includePast") === "true";

		const supabase = await createClient();

		// Fetch events user has participated in
		const { data: participants, error: participantsError } = await supabase
			.from("event_participants")
			.select("event_id, status")
			.eq("user_id", userId)
			.eq("status", "going");

		if (participantsError) {
			console.error("Error fetching subscribed events:", participantsError);
			return NextResponse.json(
				{ error: "Failed to fetch subscribed events" },
				{ status: 500 },
			);
		}

		const eventIds =
			participants
				?.map((p) => p.event_id)
				.filter((id): id is number => id != null) || [];

		if (eventIds.length === 0) {
			return NextResponse.json([]);
		}

		// Fetch event details
		const { data: events, error: eventsError } = await supabase
			.from("events")
			.select("id, title, start_time, location, banner_url")
			.in("id", eventIds);

		if (eventsError) {
			console.error("Error fetching event details:", eventsError);
			return NextResponse.json(
				{ error: "Failed to fetch event details" },
				{ status: 500 },
			);
		}

		const now = new Date();
		const filteredEvents = (events || [])
			.filter((event) => {
				const eventDate = new Date(event.start_time);
				if (includePast) {
					// Return only past events
					return eventDate < now;
				} else {
					// Return only upcoming events
					return eventDate >= now;
				}
			})
			.sort((a, b) => {
				const dateA = new Date(a.start_time).getTime();
				const dateB = new Date(b.start_time).getTime();
				return includePast ? dateB - dateA : dateA - dateB;
			});

		const formattedEvents: SubscriptionEvent[] = filteredEvents.map((event) => {
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
				date: dateStr,
				location: event.location || "",
				imageUrl: event.banner_url || "",
				startTime: timeStr,
			};
		});

		return NextResponse.json(formattedEvents);
	} catch (error) {
		console.error("Unexpected error:", error);
		return NextResponse.json(
			{ error: "Internal server error" },
			{ status: 500 },
		);
	}
}

export async function POST(request: Request) {
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

		const body = await request.json();
		const { eventId, action } = body;

		if (!eventId) {
			return NextResponse.json(
				{ error: "eventId is required" },
				{ status: 400 },
			);
		}

		if (action !== "subscribe" && action !== "unsubscribe") {
			return NextResponse.json(
				{ error: "action must be 'subscribe' or 'unsubscribe'" },
				{ status: 400 },
			);
		}

		const supabase = await createClient();

		// Ensure user exists in profiles table (lazy registration)
		const { error: upsertError } = await supabase.from("profiles").upsert(
			{
				id: userId,
				username: authResult.initData.user?.username || `user_${userId}`,
				display_name: authResult.initData.user?.firstName || "",
				role: "fan",
				created_at: new Date().toISOString(),
			},
			{ onConflict: "id" },
		);

		if (upsertError) {
			console.error("Error upserting user profile:", upsertError);
			return NextResponse.json(
				{ error: "Failed to register user" },
				{ status: 500 },
			);
		}

		if (action === "subscribe") {
			// Check if already subscribed
			const { data: existing } = await supabase
				.from("event_participants")
				.select("id")
				.eq("event_id", eventId)
				.eq("user_id", userId)
				.single();

			if (existing) {
				return NextResponse.json(
					{ error: "Already subscribed to this event" },
					{ status: 409 },
				);
			}

			// Add subscription
			const { error } = await supabase.from("event_participants").insert({
				event_id: eventId,
				user_id: userId,
				status: "going",
			});

			if (error) {
				console.error("Error subscribing to event:", error);
				return NextResponse.json(
					{ error: "Failed to subscribe to event" },
					{ status: 500 },
				);
			}

			// Get updated participant count
			const { count } = await supabase
				.from("event_participants")
				.select("*", { count: "exact", head: true })
				.eq("event_id", eventId)
				.eq("status", "going");

			return NextResponse.json(
				{
					success: true,
					message: "Successfully subscribed to event",
					participantCount: count || 0,
				},
				{ status: 201 },
			);
		} else {
			// Unsubscribe
			const { error } = await supabase
				.from("event_participants")
				.delete()
				.eq("event_id", eventId)
				.eq("user_id", userId);

			if (error) {
				console.error("Error unsubscribing from event:", error);
				return NextResponse.json(
					{ error: "Failed to unsubscribe from event" },
					{ status: 500 },
				);
			}

			// Get updated participant count
			const { count } = await supabase
				.from("event_participants")
				.select("*", { count: "exact", head: true })
				.eq("event_id", eventId)
				.eq("status", "going");

			return NextResponse.json(
				{
					success: true,
					message: "Successfully unsubscribed from event",
					participantCount: count || 0,
				},
				{ status: 200 },
			);
		}
	} catch (error) {
		console.error("Unexpected error:", error);
		return NextResponse.json(
			{ error: "Internal server error" },
			{ status: 500 },
		);
	}
}
