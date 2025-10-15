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

// Mock data
const mockProfile: UserProfile = {
	username: "RaveKid2024",
	avatar:
		"https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=150&h=150&fit=crop&crop=face",
	isDJ: true,
	bio: "Electronic music enthusiast and weekend DJ. Love deep house and techno vibes.",
	socialLinks: {
		instagram: "https://instagram.com/ravekid2024",
		soundcloud: "https://soundcloud.com/ravekid2024",
		spotify: "https://spotify.com/ravekid2024",
	},
	upcomingSets: [
		{
			id: "set1",
			event: "Late Night Sessions",
			date: "Jan 12",
			venue: "Club Voltage",
		},
		{
			id: "set2",
			event: "Weekend Warrior",
			date: "Jan 19",
			venue: "The Underground",
		},
	],
	subscribedEvents: [
		{
			id: "2",
			title: "Neon Nights Festival",
			date: "Dec 22",
			location: "Metro Convention Center",
			imageUrl:
				"https://images.unsplash.com/photo-1630497326964-62cd41a012d7?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVjdHJvbmljJTIwbXVzaWMlMjBmZXN0aXZhbHxlbnwxfHx8fDE3NTgxNTA5NTZ8MA&ixlib=rb-4.1.0&q=80&w=1080",
		},
		{
			id: "4",
			title: "Rave Revolution",
			date: "Jan 5",
			location: "Underground Tunnel System",
			imageUrl:
				"https://images.unsplash.com/photo-1465917031443-a76ab279572f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxyYXZlJTIwdW5kZXJncm91bmR8ZW58MXx8fHwxNzU4MjI1NTEzfDA&ixlib=rb-4.1.0&q=80&w=1080",
		},
	],
	settings: {
		notifications: true,
		preferredVenues: ["The Basement Club", "Metro Convention Center"],
	},
};

export default function Home() {
	const [activeTab, setActiveTab] = useState<"events" | "profile">("events");
	const [events, setEvents] = useState<Event[]>([]);
	const [profile, setProfile] = useState(mockProfile);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		let isMounted = true;

		async function fetchEvents() {
			try {
				setIsLoading(true);
				setError(null);
				const initData = getInitData();

				const response = await fetch("/api/events", {
					headers: {
						[TELEGRAM_INIT_DATA_HEADER]: initData,
					},
				});
				if (!response.ok) {
					const errorMessage =
						response.status === 401
							? "Authentication failed. Please restart the app."
							: response.status === 500
								? "Server error. Please try again later."
								: `Failed to load events (${response.status})`;
					throw new Error(errorMessage);
				}
				const data = await response.json();

				// Only update state if component is still mounted
				if (isMounted) {
					setEvents(data);
				}
			} catch (error) {
				console.error("Error fetching events:", error);
				if (isMounted) {
					setError(
						error instanceof Error
							? error.message
							: "Unable to load events. Please check your connection.",
					);
				}
			} finally {
				if (isMounted) {
					setIsLoading(false);
				}
			}
		}

		fetchEvents();

		return () => {
			isMounted = false;
		};
	}, []);

	const handleToggleSubscription = (eventId: string) => {
		setEvents((prevEvents) => {
			const updatedEvents = prevEvents.map((event) =>
				event.id === eventId
					? {
							...event,
							isSubscribed: !event.isSubscribed,
							participantCount: event.isSubscribed
								? event.participantCount - 1
								: event.participantCount + 1,
						}
					: event,
			);
			// Find the updated event after toggling
			const updatedEvent = updatedEvents.find((e) => e.id === eventId);
			if (updatedEvent) {
				if (updatedEvent.isSubscribed) {
					// Add to subscribed events
					setProfile((prev) => ({
						...prev,
						subscribedEvents: [
							...prev.subscribedEvents,
							{
								id: updatedEvent.id,
								title: updatedEvent.title,
								date: updatedEvent.date,
								location: updatedEvent.location,
								imageUrl: updatedEvent.imageUrl,
							},
						],
					}));
				} else {
					// Remove from subscribed events
					setProfile((prev) => ({
						...prev,
						subscribedEvents: prev.subscribedEvents.filter(
							(e) => e.id !== eventId,
						),
					}));
				}
			}
			return updatedEvents;
		});
	};

	const handleUpdateProfile = (updates: Partial<typeof profile>) => {
		setProfile((prev) => ({ ...prev, ...updates }));
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
						) : (
							<ProfilePage
								profile={profile}
								onUpdateProfile={handleUpdateProfile}
							/>
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
