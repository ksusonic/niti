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
			<header className="sticky top-0 z-50 bg-gradient-to-b from-black via-black/95 to-black/90 backdrop-blur-xl border-b border-blue-500/20">
				<div className="flex items-center justify-between px-5 py-5 max-w-7xl mx-auto">
					<div className="flex items-center gap-3">
						<div className="relative">
							<Image
								src="/logo.png"
								alt="NITI"
								width={48}
								height={48}
								className="w-12 h-12 object-contain flex-shrink-0 drop-shadow-[0_0_8px_rgba(59,130,246,0.5)]"
							/>
						</div>
						<div className="flex flex-col">
							<h1 className="text-2xl font-bold text-white !m-0 !mt-2 tracking-tight">
								События
							</h1>
							<p className="text-xs text-gray-400 !m-0">
								Лучшие тусовки Рязани
							</p>
						</div>
					</div>
				</div>
			</header>

			{/* Events List */}
			<div className="p-5 space-y-5 pb-24" role="feed">
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
