import confetti from "canvas-confetti";
import { motion } from "framer-motion";
import { ExternalLink, MapPin, Users } from "lucide-react";
import { useEffect, useId, useRef, useState } from "react";
import { Avatar, Badge, Button, IconButton } from "@/components/ui";
import { sanitizeVideoUrl } from "@/lib/video-url-validator";
import type { Event } from "@/types/events";

const CONFETTI_COLORS = {
	gold: "#FFD700",
	orange: "#FFA500",
	pink: "#FF69B4",
	cyan: "#00CED1",
	purple: "#9370DB",
} as const;

const CONFETTI_COLOR_PALETTE = [
	CONFETTI_COLORS.gold,
	CONFETTI_COLORS.orange,
	CONFETTI_COLORS.pink,
	CONFETTI_COLORS.cyan,
	CONFETTI_COLORS.purple,
];

const CONFETTI_WARM_COLORS = [
	CONFETTI_COLORS.gold,
	CONFETTI_COLORS.orange,
	CONFETTI_COLORS.pink,
];

const CONFETTI_COOL_COLORS = [
	CONFETTI_COLORS.cyan,
	CONFETTI_COLORS.purple,
	CONFETTI_COLORS.pink,
];

const CONFETTI_CONFIG = {
	main: {
		particleCount: 100,
		spread: 70,
		ticks: 200,
	},
	side: {
		particleCount: 50,
		spread: 55,
	},
	offset: 0.1,
} as const;

const triggerConfettiCelebration = (x: number, y: number) => {
	confetti({
		particleCount: CONFETTI_CONFIG.main.particleCount,
		spread: CONFETTI_CONFIG.main.spread,
		origin: { x, y },
		colors: CONFETTI_COLOR_PALETTE,
		ticks: CONFETTI_CONFIG.main.ticks,
	});

	confetti({
		particleCount: CONFETTI_CONFIG.side.particleCount,
		angle: 60,
		spread: CONFETTI_CONFIG.side.spread,
		origin: { x: x - CONFETTI_CONFIG.offset, y },
		colors: CONFETTI_WARM_COLORS,
	});

	confetti({
		particleCount: CONFETTI_CONFIG.side.particleCount,
		angle: 120,
		spread: CONFETTI_CONFIG.side.spread,
		origin: { x: x + CONFETTI_CONFIG.offset, y },
		colors: CONFETTI_COOL_COLORS,
	});
};

interface EventCardProps {
	event: Event;
	onToggleSubscription: (eventId: string) => void;
}

export function EventCard({ event, onToggleSubscription }: EventCardProps) {
	const safeVideoUrl = sanitizeVideoUrl(event.videoUrl);
	const buttonRef = useRef<HTMLButtonElement>(null);
	const lineupHeadingId = useId();
	const [isLoading, setIsLoading] = useState(false);
	const [previousSubscriptionState, setPreviousSubscriptionState] = useState(
		event.isSubscribed,
	);

	// Trigger confetti when subscription state changes and reset loading
	useEffect(() => {
		if (previousSubscriptionState !== event.isSubscribed) {
			// State changed, reset loading and trigger confetti if subscribed
			setIsLoading(false);
			if (
				!previousSubscriptionState &&
				event.isSubscribed &&
				buttonRef.current
			) {
				const rect = buttonRef.current.getBoundingClientRect();
				const x = (rect.left + rect.width / 2) / window.innerWidth;
				const y = (rect.top + rect.height / 2) / window.innerHeight;

				triggerConfettiCelebration(x, y);
			}
			setPreviousSubscriptionState(event.isSubscribed);
		}
	}, [event.isSubscribed, previousSubscriptionState]);

	const handleSubscribe = () => {
		setIsLoading(true);
		onToggleSubscription(event.id);
	};

	return (
		<motion.article
			initial={{ opacity: 0, y: 20 }}
			animate={{ opacity: 1, y: 0 }}
			whileHover={{ y: -8, scale: 1.01 }}
			transition={{ duration: 0.3 }}
			className="relative bg-gradient-to-br from-gray-900/80 to-gray-900/40 backdrop-blur-sm rounded-2xl overflow-hidden border border-gray-800/50 shadow-2xl hover:border-blue-500/30 hover:shadow-blue-500/10"
			aria-label={`Event: ${event.title}`}
		>
			{/* Video/Image Background */}
			<div className="relative h-72 overflow-hidden">
				{safeVideoUrl ? (
					<>
						<iframe
							src={safeVideoUrl}
							className="absolute inset-0 w-full h-full object-cover pointer-events-none"
							allow="autoplay; fullscreen"
							title={`${event.title} video`}
						/>
						<div className="absolute inset-0 bg-gradient-to-t from-black via-black/50 to-transparent" />
					</>
				) : (
					<>
						<div
							className="absolute inset-0 bg-cover bg-center blur-sm scale-110 opacity-40"
							style={{ backgroundImage: `url(${event.imageUrl})` }}
							role="img"
							aria-label={event.title}
						/>
						<div className="absolute inset-0 bg-gradient-to-t from-black via-black/50 to-transparent" />
					</>
				)}

				{/* Event Date/Time Badge */}
				<div className="absolute top-4 left-4 z-20">
					<Badge 
						variant="primary" 
						className="bg-blue-500/90 backdrop-blur-sm text-white border-blue-400/50 shadow-lg shadow-blue-500/20"
					>
						<time dateTime={event.date}>
							{event.date} • {event.time}
						</time>
					</Badge>
				</div>

				{/* Event Title Overlay */}
				<header className="absolute bottom-0 left-0 right-0 z-10 p-6 bg-gradient-to-t from-black/80 to-transparent">
					<h3 className="text-3xl font-bold text-white mb-2 drop-shadow-lg">{event.title}</h3>
					<address className="flex items-center gap-2 text-sm text-gray-300 not-italic">
						<MapPin className="h-4 w-4 text-blue-400" aria-hidden="true" />
						<span>{event.location}</span>
					</address>
				</header>
			</div>

			{/* Content */}
			<div className="p-5 space-y-5">
				{/* Event Description */}
				<p className="text-gray-400 text-sm leading-relaxed">
					{event.description}
				</p>

				{/* DJ Lineup */}
				<section className="space-y-3" aria-labelledby={lineupHeadingId}>
					<div className="flex items-center gap-2">
						<div className="p-1.5 bg-blue-500/20 rounded-lg">
							<Users className="h-4 w-4 text-blue-400" />
						</div>
						<h4
							id={lineupHeadingId}
							className="text-sm font-semibold text-white"
						>
							Lineup
						</h4>
					</div>
					<ul className="space-y-2">
						{event.djLineup.map((dj) => (
							<motion.li
								key={dj.id}
								whileHover={{ x: 4, scale: 1.01 }}
								className="flex items-center justify-between bg-black/30 backdrop-blur-sm rounded-xl p-4 border border-gray-800/50 hover:border-blue-500/30 transition-all"
							>
								<div className="flex items-center gap-3 flex-1">
									<Avatar
										src={dj.avatar}
										alt={dj.name}
										className="border-2 border-blue-500/30"
									/>
									<div className="flex-1 min-w-0">
										<span className="font-semibold block truncate text-white">
											{dj.name}
										</span>
										<time className="text-xs text-gray-400">{dj.time}</time>
									</div>
								</div>

								{/* Social Links */}
								<nav
									className="flex items-center gap-1.5"
									aria-label={`${dj.name} social links`}
								>
									{dj.social.instagram && (
										<IconButton
											href={dj.social.instagram}
											target="_blank"
											rel="noopener noreferrer"
											variant="blue"
											aria-label={`${dj.name} on Instagram`}
										>
											<ExternalLink
												className="h-3.5 w-3.5"
												aria-hidden="true"
											/>
										</IconButton>
									)}
									{dj.social.soundcloud && (
										<IconButton
											href={dj.social.soundcloud}
											target="_blank"
											rel="noopener noreferrer"
											variant="orange"
											aria-label={`${dj.name} on SoundCloud`}
										>
											<ExternalLink
												className="h-3.5 w-3.5"
												aria-hidden="true"
											/>
										</IconButton>
									)}
									{dj.social.spotify && (
										<IconButton
											href={dj.social.spotify}
											target="_blank"
											rel="noopener noreferrer"
											variant="green"
											aria-label={`${dj.name} on Spotify`}
										>
											<ExternalLink
												className="h-3.5 w-3.5"
												aria-hidden="true"
											/>
										</IconButton>
									)}
								</nav>
							</motion.li>
						))}
					</ul>
				</section>

				{/* Participants and Subscribe Button */}
				<footer className="flex items-center justify-between pt-4 border-t border-gray-800/50">
					<div className="flex items-center gap-2 text-sm">
						<div className="p-2 bg-blue-500/20 rounded-lg">
							<Users className="h-4 w-4 text-blue-400" aria-hidden="true" />
						</div>
						<span className="text-gray-400">
							<strong className="font-semibold text-white">
								{event.participantCount}
							</strong>{" "}
							going
						</span>
					</div>

					<div className="w-32">
						<Button
							ref={buttonRef}
							onClick={handleSubscribe}
							isLoading={isLoading}
							variant={event.isSubscribed ? "secondary" : "primary"}
							aria-pressed={event.isSubscribed}
							className="w-full"
						>
							{event.isSubscribed ? "Joined" : "Join Event"}
						</Button>
					</div>
				</footer>
			</div>
		</motion.article>
	);
}
