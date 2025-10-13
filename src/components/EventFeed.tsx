import Image from "next/image";
import type { Event } from "@/types/events";
import { EventCard } from "./EventCard";

interface EventFeedProps {
	events: Event[];
	onToggleSubscription: (eventId: string) => void;
}

export function EventFeed({ events, onToggleSubscription }: EventFeedProps) {
	return (
		<main className="min-h-screen bg-background">
			<header className="sticky top-0 z-50 bg-background/95 backdrop-blur-lg border-b border-border/50">
				<div className="flex items-center justify-between px-6 py-4 max-w-7xl mx-auto">
					<div className="flex items-center gap-3">
						<Image
							src="/logo.png"
							alt="NITI"
							width={40}
							height={40}
							className="w-10 h-10 object-contain flex-shrink-0"
						/>
						<div className="flex flex-col">
							<h1 className="text-xl font-semibold text-white !m-0 tracking-tight">
								Events
							</h1>
							<p className="text-xs text-muted-foreground !m-0">
								Discover amazing DJ events
							</p>
						</div>
					</div>
				</div>
			</header>

			{/* Events List */}
			<div className="p-4 space-y-6" role="feed">
				{events.map((event) => (
					<div key={event.id}>
						<EventCard
							event={event}
							onToggleSubscription={onToggleSubscription}
						/>
					</div>
				))}
			</div>
		</main>
	);
}
