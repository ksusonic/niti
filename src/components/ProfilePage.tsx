import type { User } from "@telegram-apps/sdk-react";
import { motion } from "framer-motion";
import { Bell, Calendar, Instagram, MapPin, Music } from "lucide-react";
import Image from "next/image";
import { Avatar, Badge, Card, IconButton, Switch } from "@/components/ui";

interface ProfileData {
	isDJ: boolean;
	bio?: string;
	socialLinks?: {
		telegram?: string;
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
		<div className="bg-black min-h-screen">
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
						<h1 className="text-2xl font-bold text-white !m-0 !mt-2 tracking-tight">
							{user.first_name}
							{user.last_name ? ` ${user.last_name}` : ""}
						</h1>
						{profileData.bio ? (
							<p className="text-gray-400 text-sm mt-0.5 truncate">
								{profileData.bio}
							</p>
						) : (
							<p className="text-gray-500 text-sm mt-0.5">
								{profileData.isDJ ? "Диджей" : "Гость"}
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
								Связь
							</h3>
							<nav
								className="flex items-center gap-2"
								aria-label="Social media links"
							>
								{profileData.socialLinks.telegram && (
									<IconButton
										href={profileData.socialLinks.telegram}
										target="_blank"
										rel="noopener noreferrer"
										variant="blue"
										aria-label="Telegram"
										className="flex-1"
									>
										<svg
											className="h-4 w-4"
											viewBox="0 0 24 24"
											fill="currentColor"
											aria-hidden="true"
										>
											<title>Telegram</title>
											<path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69a.2.2 0 00-.05-.18c-.06-.05-.14-.03-.21-.02-.09.02-1.49.95-4.22 2.79-.4.27-.76.41-1.08.4-.36-.01-1.04-.2-1.55-.37-.63-.2-1.12-.31-1.08-.66.02-.18.27-.36.74-.55 2.92-1.27 4.86-2.11 5.83-2.51 2.78-1.16 3.35-1.36 3.73-1.36.08 0 .27.02.39.12.1.08.13.19.14.27-.01.06.01.24 0 .38z" />
										</svg>
									</IconButton>
								)}
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
										<svg
											className="h-4 w-4"
											viewBox="0 0 24 24"
											fill="currentColor"
											aria-hidden="true"
										>
											<title>SoundCloud</title>
											<path d="M1.5 15.5h.938V11.875h-.938zm1.563 0h.937V10.938h-.937zm1.562 0h.938V12.813h-.938zm1.563 0h.937V10h-.937zm1.562 0h.938v-6.25h-.938zm1.563 0h.937V8.125h-.937zm1.562 0h.938V10h-.938zM20 13.75c0 1.031-.813 1.875-1.813 1.875H11.25V7.531c1.156.438 2.281 1.156 3.094 2.094C15.719 8.188 17.75 7.5 19.875 7.5c1.375 0 2.625.75 3.281 1.938.625 1.094.656 2.406.063 3.531C22.531 13.844 21.281 14.5 20 13.75z" />
										</svg>
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
										<svg
											className="h-4 w-4"
											viewBox="0 0 24 24"
											fill="currentColor"
											aria-hidden="true"
										>
											<title>Spotify</title>
											<path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm3.64 14.3c-.14.23-.44.3-.68.17-1.87-1.14-4.22-1.4-6.99-.77-.27.07-.55-.1-.62-.37-.07-.27.1-.55.37-.62 3.04-.69 5.64-.39 7.72.89.23.13.3.43.17.68l.03.02zm.98-2.18c-.18.29-.56.38-.85.2-2.14-1.31-5.4-1.69-7.94-.93-.33.1-.68-.09-.78-.42s.09-.68.42-.78c2.9-.87 6.54-.45 8.95 1.08.29.18.38.56.2.85zm.09-2.27c-2.57-1.52-6.81-1.66-9.26-.92-.4.12-.82-.11-.94-.51s.11-.82.51-.94c2.81-.85 7.5-.69 10.44 1.07.36.21.48.68.27 1.04-.21.36-.68.48-1.04.27l.02-.01z" />
										</svg>
									</IconButton>
								)}
							</nav>
						</Card>
					</motion.div>
				)}

				{/* DJ Profile Section */}
				{profileData.isDJ &&
					profileData.upcomingSets &&
					profileData.upcomingSets.length > 0 && (
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
										Выступления
									</h3>
								</div>
								<ul className="space-y-2">
									{profileData.upcomingSets.map((set) => (
										<motion.li
											key={set.id}
											whileHover={{ scale: 1.01 }}
											className="p-4 bg-black/30 backdrop-blur-sm rounded-xl border border-gray-800/50 hover:border-blue-500/30 transition-all"
										>
											<p className="font-semibold text-white mb-2">
												{set.event}
											</p>
											<div className="flex items-center gap-4 text-xs text-gray-400">
												<time className="flex items-center gap-1.5">
													<Calendar
														className="h-3.5 w-3.5 text-blue-400"
														aria-hidden="true"
													/>
													<span>{set.date}</span>
												</time>
												<address className="flex items-center gap-1.5 not-italic">
													<MapPin
														className="h-3.5 w-3.5 text-blue-400"
														aria-hidden="true"
													/>
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
							<h3 className="text-base font-semibold text-white">Настройки</h3>
						</div>
						<div className="flex items-center justify-between p-4 bg-black/30 rounded-xl">
							<div>
								<p className="font-medium text-white text-sm">Уведомления</p>
								<p className="text-xs text-gray-400 mt-0.5">
									Принимать сообщения о новых ивентах
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
				{profileData.subscribedEvents.length > 0 ? (
					<motion.div
						initial={{ opacity: 0, y: 20 }}
						animate={{ opacity: 1, y: 0 }}
						transition={{ duration: 0.3, delay: 0.3 }}
					>
						<Card className="p-5 bg-gradient-to-br from-gray-900/80 to-gray-900/40 border-gray-800/50">
							<div className="flex items-center justify-between mb-4">
								<h3 className="text-base font-semibold text-white">
									Мои события
								</h3>
								<Badge
									variant="primary"
									className="bg-blue-500/20 text-blue-400 border border-blue-500/30"
								>
									{profileData.subscribedEvents.length}
								</Badge>
							</div>
							<ul className="space-y-2">
								{profileData.subscribedEvents.map((event, index) => (
									<motion.li
										key={event.id}
										initial={{ opacity: 0, x: -20 }}
										animate={{ opacity: 1, x: 0 }}
										transition={{
											duration: 0.3,
											delay: Math.min(0.1 * index, 0.3),
										}}
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
				) : (
					<motion.div
						initial={{ opacity: 0, y: 20 }}
						animate={{ opacity: 1, y: 0 }}
						transition={{ duration: 0.3, delay: 0.3 }}
					>
						<Card className="p-5 bg-gradient-to-br from-gray-900/80 to-gray-900/40 border-gray-800/50">
							<div className="flex items-center justify-between mb-4">
								<h3 className="text-base font-semibold text-white">
									Мои события
								</h3>
							</div>
							<p className="text-gray-400 text-sm">У вас пока нет событий.</p>
						</Card>
					</motion.div>
				)}
			</div>
		</div>
	);
}
