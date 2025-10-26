"use client";

import type { User } from "@tma.js/sdk-react";
import { AnimatePresence, motion } from "framer-motion";
import { useCallback, useEffect, useState } from "react";
import { BottomNavigation } from "@/components/BottomNavigation";
import { EventFeed } from "@/components/EventFeed";
import { ProfilePage } from "@/components/ProfilePage";
import { ErrorState, LoadingState } from "@/components/ui";
import { useInitData } from "@/hooks/useTelegramUser";
import { authenticatedFetchJson } from "@/lib/api-client";
import type { Event } from "@/types/events";

interface ProfileData {
	isDJ: boolean;
	bio?: string;
	socialLinks?: {
		instagram?: string;
		soundcloud?: string;
		spotify?: string;
	};
	upcomingSets?: Array<{
		id: string;
		event: string;
		date: string;
		venue: string;
	}>;
	subscribedEvents: Array<{
		id: string;
		title: string;
		date: string;
		location: string;
		imageUrl: string;
	}>;
	settings: {
		notifications: boolean;
		preferredVenues: string[];
	};
}

export default function Home() {
	const [activeTab, setActiveTab] = useState<"events" | "profile">("events");
	const [direction, setDirection] = useState(0);
	const [events, setEvents] = useState<Event[]>([]);
	const [user, setUser] = useState<User | null>(null);
	const [profileData, setProfileData] = useState<ProfileData | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	const initData = useInitData();

	useEffect(() => {
		let isMounted = true;

		async function fetchData() {
			try {
				setIsLoading(true);
				setError(null);

				if (!initData?.user) {
					throw new Error("Telegram init data not available");
				}

				// Fetch events with automatic authentication
				const eventsData = await authenticatedFetchJson<Event[]>("/api/events");

				// Fetch subscribed events (upcoming)
				let subscribedEvents: Event[] = [];
				try {
					subscribedEvents = await authenticatedFetchJson<Event[]>(
						"/api/subscriptions?includePast=false",
					);
				} catch (error) {
					console.warn(
						"Failed to fetch subscriptions, continuing without them:",
						error,
					);
				}

				// Create profile from typed Telegram user data
				if (isMounted) {
					const localProfileData: ProfileData = {
						isDJ: false,
						subscribedEvents,
						settings: {
							notifications: true,
							preferredVenues: [],
						},
					};
					setUser(initData.user);
					setProfileData(localProfileData);
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

		if (initData) {
			fetchData();
		} else {
			// Wait a bit for SDK to initialize
			const timeoutId = setTimeout(() => {
				if (!initData) {
					setError("Waiting for Telegram initialization...");
				}
			}, 1000);

			return () => clearTimeout(timeoutId);
		}

		return () => {
			isMounted = false;
		};
	}, [initData]);

	const handleToggleSubscription = async (eventId: string) => {
		const event = events.find((e) => e.id === eventId);
		if (!event || !initData) return;

		const action = event.isSubscribed ? "unsubscribe" : "subscribe";

		try {
			const data = await authenticatedFetchJson<{
				success: boolean;
				participantCount: number;
			}>("/api/subscriptions", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					eventId: parseInt(eventId, 10),
					action,
				}),
			});

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

			// Update profile subscribed events
			if (action === "subscribe" && event) {
				setProfileData((prev) => {
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
				setProfileData((prev) => {
					if (!prev) return prev;
					return {
						...prev,
						subscribedEvents: prev.subscribedEvents.filter(
							(e) => e.id !== eventId,
						),
					};
				});
			}
		} catch (error) {
			console.error(`Error ${action}ing from event:`, error);
		}
	};

	const handleUpdateProfileData = (updates: Partial<ProfileData>) => {
		setProfileData((prev) => {
			if (!prev) return prev;
			return { ...prev, ...updates };
		});
	};

	const handleTabChange = useCallback(
		(tab: "events" | "profile") => {
			// Determine direction: events is 0 (left), profile is 1 (right)
			const tabIndex = tab === "events" ? 0 : 1;
			const currentIndex = activeTab === "events" ? 0 : 1;
			setDirection(tabIndex - currentIndex);
			setActiveTab(tab);
		},
		[activeTab],
	);

	const slideVariants = {
		enter: (direction: number) => ({
			x: direction > 0 ? "100%" : "-100%",
			position: "absolute" as const,
		}),
		center: {
			x: 0,
			position: "relative" as const,
		},
		exit: (direction: number) => ({
			x: direction > 0 ? "-100%" : "100%",
			position: "absolute" as const,
		}),
	};

	return (
		<div className="bg-background text-foreground dark pb-24 relative overflow-hidden">
			{isLoading ? (
				<LoadingState message="Loading events..." />
			) : error ? (
				<ErrorState error={error} />
			) : (
				<AnimatePresence initial={false} custom={direction}>
					<motion.div
						key={activeTab}
						custom={direction}
						variants={slideVariants}
						initial="enter"
						animate="center"
						exit="exit"
						transition={{
							x: { type: "spring", stiffness: 300, damping: 30 },
							duration: 0.2,
						}}
						drag="x"
						dragConstraints={{ left: 0, right: 0 }}
						dragElastic={0.2}
						onDragEnd={(_, info) => {
							const swipeThreshold = 50;
							if (info.offset.x > swipeThreshold && activeTab === "profile") {
								// Swipe right → go to events (left)
								handleTabChange("events");
							} else if (
								info.offset.x < -swipeThreshold &&
								activeTab === "events"
							) {
								// Swipe left → go to profile (right)
								handleTabChange("profile");
							}
						}}
						className="min-h-screen w-full"
					>
						{activeTab === "events" ? (
							<EventFeed
								events={events}
								onToggleSubscription={handleToggleSubscription}
							/>
						) : user && profileData ? (
							<ProfilePage
								user={user}
								profileData={profileData}
								onUpdateProfileData={handleUpdateProfileData}
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
