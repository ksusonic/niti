import type { User } from "@telegram-apps/sdk-react";
import { motion } from "framer-motion";
import {
	Bell,
	Calendar,
	Instagram,
	MapPin,
	Music,
	Music2,
} from "lucide-react";
import Image from "next/image";
import { Avatar, Badge, Card, IconButton, Switch } from "@/components/ui";

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

interface ProfilePageProps {
	user: User;
	profileData: ProfileData;
	onUpdateProfileData: (updates: Partial<ProfileData>) => void;
}

export function ProfilePage({
	user,
	profileData,
	onUpdateProfileData,
}: ProfilePageProps) {
	const handleToggleNotifications = () => {
		onUpdateProfileData({
			settings: {
				...profileData.settings,
				notifications: !profileData.settings.notifications,
			},
		});
	};

	return (
		<div className="bg-background min-h-screen">
			{/* Header */}
			<header className="sticky top-0 z-50 bg-gradient-to-b from-black via-black/95 to-black/90 backdrop-blur-xl border-b border-blue-500/20 px-5 py-5">
				<div className="flex items-center gap-4">
					<div className="relative">
						<Avatar
							src={user.photo_url || ""}
							alt={`${user.first_name}${user.last_name ? ` ${user.last_name}` : ""}`}
							size="lg"
							className="border-2 border-blue-500/40 shadow-lg shadow-blue-500/20"
						/>
						{profileData.isDJ && (
							<div className="absolute -bottom-1 -right-1 bg-blue-500 rounded-full p-1.5 border-2 border-black">
								<Music className="h-3 w-3 text-white" />
							</div>
						)}
					</div>
					<div className="flex-1 min-w-0">
						<h1 className="text-2xl font-bold text-white truncate tracking-tight">
							{user.first_name}
							{user.last_name ? ` ${user.last_name}` : ""}
						</h1>
						{profileData.bio ? (
							<p className="text-gray-400 text-sm mt-0.5 truncate">
								{profileData.bio}
							</p>
						) : (
							<p className="text-gray-500 text-sm mt-0.5">
								{profileData.isDJ ? "DJ" : "Party enthusiast"}
							</p>
						)}
					</div>
				</div>
			</header>

			<div className="p-5 space-y-5 pb-24">
				{/* Social Links */}
				{profileData.socialLinks && (
					<motion.div
						initial={{ opacity: 0, y: 20 }}
						animate={{ opacity: 1, y: 0 }}
						transition={{ duration: 0.3 }}
					>
						<Card className="p-5 bg-gradient-to-br from-gray-900/80 to-gray-900/40 border-gray-800/50">
							<h3 className="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-3">
								Connect
							</h3>
							<nav
								className="flex items-center gap-2"
								aria-label="Social media links"
							>
								{profileData.socialLinks.instagram && (
									<IconButton
										href={profileData.socialLinks.instagram}
										target="_blank"
										rel="noopener noreferrer"
										variant="blue"
										aria-label="Instagram"
										className="flex-1"
									>
										<Instagram className="h-4 w-4" />
									</IconButton>
								)}
								{profileData.socialLinks.soundcloud && (
									<IconButton
										href={profileData.socialLinks.soundcloud}
										target="_blank"
										rel="noopener noreferrer"
										variant="orange"
										aria-label="SoundCloud"
										className="flex-1"
									>
										<Music2 className="h-4 w-4" />
									</IconButton>
								)}
								{profileData.socialLinks.spotify && (
									<IconButton
										href={profileData.socialLinks.spotify}
										target="_blank"
										rel="noopener noreferrer"
										variant="green"
										aria-label="Spotify"
										className="flex-1"
									>
										<Music className="h-4 w-4" />
									</IconButton>
								)}
							</nav>
						</Card>
					</motion.div>
				)}

				{/* DJ Profile Section */}
				{profileData.isDJ && profileData.upcomingSets && profileData.upcomingSets.length > 0 && (
					<motion.div
						initial={{ opacity: 0, y: 20 }}
						animate={{ opacity: 1, y: 0 }}
						transition={{ duration: 0.3, delay: 0.1 }}
					>
						<Card className="p-5 bg-gradient-to-br from-blue-900/20 to-purple-900/20 border-blue-500/30">
							<div className="flex items-center gap-2 mb-4">
								<div className="p-2 bg-blue-500/20 rounded-lg">
									<Music className="h-4 w-4 text-blue-400" />
								</div>
								<h3 className="text-base font-semibold text-white">
									Upcoming Sets
								</h3>
							</div>
							<ul className="space-y-2">
								{profileData.upcomingSets.map((set) => (
									<motion.li
										key={set.id}
										whileHover={{ scale: 1.01 }}
										className="p-4 bg-black/30 backdrop-blur-sm rounded-xl border border-gray-800/50 hover:border-blue-500/30 transition-all"
									>
										<p className="font-semibold text-white mb-2">{set.event}</p>
										<div className="flex items-center gap-4 text-xs text-gray-400">
											<time className="flex items-center gap-1.5">
												<Calendar className="h-3.5 w-3.5 text-blue-400" aria-hidden="true" />
												<span>{set.date}</span>
											</time>
											<address className="flex items-center gap-1.5 not-italic">
												<MapPin className="h-3.5 w-3.5 text-blue-400" aria-hidden="true" />
												<span>{set.venue}</span>
											</address>
										</div>
									</motion.li>
								))}
							</ul>
						</Card>
					</motion.div>
				)}

				{/* Settings */}
				<motion.div
					initial={{ opacity: 0, y: 20 }}
					animate={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.3, delay: 0.2 }}
				>
					<Card className="p-5 bg-gradient-to-br from-gray-900/80 to-gray-900/40 border-gray-800/50">
						<div className="flex items-center gap-2 mb-4">
							<div className="p-2 bg-blue-500/20 rounded-lg">
								<Bell className="h-4 w-4 text-blue-400" />
							</div>
							<h3 className="text-base font-semibold text-white">
								Preferences
							</h3>
						</div>
						<div className="flex items-center justify-between p-4 bg-black/30 rounded-xl">
							<div>
								<p className="font-medium text-white text-sm">
									Push Notifications
								</p>
								<p className="text-xs text-gray-400 mt-0.5">
									Get notified about new events
								</p>
							</div>
							<Switch
								checked={profileData.settings.notifications}
								onCheckedChange={handleToggleNotifications}
							/>
						</div>
					</Card>
				</motion.div>

				{/* Subscribed Events */}
				{profileData.subscribedEvents.length > 0 && (
					<motion.div
						initial={{ opacity: 0, y: 20 }}
						animate={{ opacity: 1, y: 0 }}
						transition={{ duration: 0.3, delay: 0.3 }}
					>
						<Card className="p-5 bg-gradient-to-br from-gray-900/80 to-gray-900/40 border-gray-800/50">
							<div className="flex items-center justify-between mb-4">
								<h3 className="text-base font-semibold text-white">
									My Events
								</h3>
								<Badge variant="primary" className="bg-blue-500/20 text-blue-400 border border-blue-500/30">
									{profileData.subscribedEvents.length}
								</Badge>
							</div>
							<ul className="space-y-2">
								{profileData.subscribedEvents.map((event, index) => (
									<motion.li
										key={event.id}
										initial={{ opacity: 0, x: -20 }}
										animate={{ opacity: 1, x: 0 }}
										transition={{ duration: 0.3, delay: 0.1 * index }}
										whileHover={{ x: 4, scale: 1.01 }}
										className="group"
									>
										<div className="flex items-center gap-3 p-3 bg-black/30 rounded-xl border border-gray-800/50 hover:border-blue-500/30 hover:bg-black/40 transition-all">
											<div className="relative">
												<Image
													src={event.imageUrl}
													alt={event.title}
													width={56}
													height={56}
													className="w-14 h-14 rounded-lg object-cover border border-blue-500/20 group-hover:border-blue-500/40 transition-colors"
												/>
												<div className="absolute inset-0 bg-blue-500/0 group-hover:bg-blue-500/10 rounded-lg transition-colors" />
											</div>
											<div className="flex-1 min-w-0">
												<p className="font-semibold text-white truncate text-sm mb-1">
													{event.title}
												</p>
												<div className="flex items-center gap-3 text-xs text-gray-400">
													<time className="flex items-center gap-1">
														<Calendar className="h-3 w-3" aria-hidden="true" />
														{event.date}
													</time>
													<address className="not-italic truncate flex items-center gap-1">
														<MapPin className="h-3 w-3" aria-hidden="true" />
														{event.location}
													</address>
												</div>
											</div>
										</div>
									</motion.li>
								))}
							</ul>
						</Card>
					</motion.div>
				)}
			</div>
		</div>
	);
}
