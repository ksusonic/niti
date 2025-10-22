"use client";

import { AnimatePresence, motion } from "framer-motion";
import { useCallback, useEffect, useState } from "react";
import { BottomNavigation } from "@/components/BottomNavigation";
import { EventFeed } from "@/components/EventFeed";
import { ProfilePage } from "@/components/ProfilePage";
import { ErrorState, LoadingState } from "@/components/ui";
import { TELEGRAM_INIT_DATA_HEADER } from "@/lib/constants";
import { getInitData } from "@/lib/telegram-init-data";
import type { Event, UserProfile } from "@/types/events";

// Helper function to extract user data from init data string
function parseInitDataForUser(
	initDataStr: string,
): {
	username: string;
	userId: number;
	firstName: string;
	lastName?: string;
} | null {
	try {
		const params = new URLSearchParams(initDataStr);
		const userStr = params.get("user");
		if (!userStr) return null;
		const user = JSON.parse(userStr);
		return {
			username: user.username || user.first_name || "User",
			userId: user.id,
			firstName: user.first_name,
			lastName: user.last_name,
		};
	} catch (error) {
		console.error("Error parsing init data:", error);
		return null;
	}
}

export default function Home() {
	const [activeTab, setActiveTab] = useState<"events" | "profile">("events");
	const [events, setEvents] = useState<Event[]>([]);
	const [profile, setProfile] = useState<UserProfile | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		let isMounted = true;

		async function fetchData() {
			try {
				setIsLoading(true);
				setError(null);
				const initData = getInitData();

				// Parse user data from initData
				const userData = parseInitDataForUser(initData);
				if (!userData) {
					throw new Error("Failed to parse user data from authentication");
				}

				// Fetch events
				const eventsResponse = await fetch("/api/events", {
					headers: {
						[TELEGRAM_INIT_DATA_HEADER]: initData,
					},
				});
				if (!eventsResponse.ok) {
					const errorMessage =
						eventsResponse.status === 401
							? "Authentication failed. Please restart the app."
							: eventsResponse.status === 500
								? "Server error. Please try again later."
								: `Failed to load events (${eventsResponse.status})`;
					throw new Error(errorMessage);
				}
				const eventsData = await eventsResponse.json();

				// Fetch subscribed events (upcoming)
				const subscriptionsResponse = await fetch(
					"/api/subscriptions?includePast=false",
					{
						headers: {
							[TELEGRAM_INIT_DATA_HEADER]: initData,
						},
					},
				);
				if (!subscriptionsResponse.ok) {
					console.warn(
						"Failed to fetch subscriptions, continuing without them",
					);
				}
				const subscribedEvents = subscriptionsResponse.ok
					? await subscriptionsResponse.json()
					: [];

				// Create local profile from parsed user data
				if (isMounted) {
					const localProfile: UserProfile = {
						username: userData.username,
						firstName: userData.firstName,
						lastName: userData.lastName,
						avatar: "", // Avatar would be set from Telegram if available
						isDJ: false, // Can be determined from database if needed
						subscribedEvents,
						settings: {
							notifications: true,
							preferredVenues: [],
						},
					};
					setProfile(localProfile);
					setEvents(eventsData);
				}
			} catch (error) {
				console.error("Error fetching data:", error);
				if (isMounted) {
					setError(
						error instanceof Error
							? error.message
							: "Unable to load data. Please check your connection.",
					);
				}
			} finally {
				if (isMounted) {
					setIsLoading(false);
				}
			}
		}

		fetchData();

		return () => {
			isMounted = false;
		};
	}, []);

	const handleToggleSubscription = (eventId: string) => {
		const event = events.find((e) => e.id === eventId);
		if (!event) return;

		const action = event.isSubscribed ? "unsubscribe" : "subscribe";
		const initData = getInitData();

		// Make API request
		fetch("/api/subscriptions", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				[TELEGRAM_INIT_DATA_HEADER]: initData,
			},
			body: JSON.stringify({
				eventId: parseInt(eventId, 10),
				action,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					return response.json().then((data) => {
						throw new Error(data.error || `Failed to ${action}`);
					});
				}
				return response.json();
			})
			.then((data) => {
				// Update UI with response data
				setEvents((prevEvents) =>
					prevEvents.map((e) =>
						e.id === eventId
							? {
									...e,
									isSubscribed: !e.isSubscribed,
									participantCount: data.participantCount,
								}
							: e,
					),
				);

				// Update profile subscribed events if subscribing
				if (action === "subscribe" && event) {
					setProfile((prev) => {
						if (!prev) return prev;
						return {
							...prev,
							subscribedEvents: [
								...prev.subscribedEvents,
								{
									id: event.id,
									title: event.title,
									date: event.date,
									location: event.location,
									imageUrl: event.imageUrl,
								},
							],
						};
					});
				} else if (action === "unsubscribe") {
					setProfile((prev) => {
						if (!prev) return prev;
						return {
							...prev,
							subscribedEvents: prev.subscribedEvents.filter(
								(e) => e.id !== eventId,
							),
						};
					});
				}
			})
			.catch((error) => {
				console.error(`Error ${action}ing from event:`, error);
			});
	};

	const handleUpdateProfile = (updates: Partial<UserProfile>) => {
		setProfile((prev) => {
			if (!prev) return prev;
			return { ...prev, ...updates };
		});
	};

	const handleTabChange = useCallback((tab: "events" | "profile") => {
		setActiveTab(tab);
	}, []);

	return (
		<div className="bg-background text-foreground dark pb-24">
			{isLoading ? (
				<LoadingState message="Loading events..." />
			) : error ? (
				<ErrorState error={error} />
			) : (
				<AnimatePresence mode="wait">
					<motion.div
						key={activeTab}
						className="min-h-screen page-transition"
						animate={{ opacity: 1, x: 0 }}
						exit={{ opacity: 0, x: -40 }}
						transition={{ duration: 0.3 }}
					>
						{activeTab === "events" ? (
							<EventFeed
								events={events}
								onToggleSubscription={handleToggleSubscription}
							/>
						) : profile ? (
							<ProfilePage
								profile={profile}
								onUpdateProfile={handleUpdateProfile}
							/>
						) : (
							<ErrorState error="Profile data not available" />
						)}
					</motion.div>
				</AnimatePresence>
			)}

			<div className="fixed bottom-0 left-0 right-0 z-50">
				<BottomNavigation activeTab={activeTab} onTabChange={handleTabChange} />
			</div>
		</div>
	);
}
