import type { User } from "@tma.js/sdk-react";
import { motion } from "framer-motion";
import {
	Bell,
	Calendar,
	Instagram,
	MapPin,
	Music,
	Music2,
	Settings,
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
		<div className="bg-background">
			{/* Header */}
			<header className="sticky top-0 z-50 bg-black/95 backdrop-blur-lg border-b border-gray-800/50 px-4 py-4">
				<div className="flex items-center justify-between">
					<h1 className="text-2xl font-bold text-foreground">Profile</h1>
					<button
						type="button"
						className="p-2 hover:bg-secondary rounded-lg transition-colors"
						aria-label="Settings"
					>
						<Settings className="h-6 w-6 text-gray-500" />
					</button>
				</div>
			</header>

			<div className="p-4 space-y-6">
				{/* Profile Header */}
				<div>
					<Card className="p-6">
						<div className="flex items-center gap-4">
							<Avatar
								src={user.photo_url || ""}
								alt={`${user.first_name}${user.last_name ? ` ${user.last_name}` : ""}`}
								size="lg"
								className="border-2 border-blue-500/30"
							/>

							<div className="flex-1">
								<div className="flex items-center gap-2">
									<h2 className="text-xl font-bold text-foreground">
										{user.first_name}
										{user.last_name ? ` ${user.last_name}` : ""}
									</h2>
									{profileData.isDJ && (
										<Badge
											variant="primary"
											className="bg-blue-500/20 text-blue-400 border border-blue-500/30"
										>
											<Music className="h-3 w-3" />
											DJ
										</Badge>
									)}
								</div>
								{profileData.bio && (
									<p className="text-muted-foreground text-sm mt-1">
										{profileData.bio}
									</p>
								)}
							</div>
						</div>

						{/* Social Links */}
						{profileData.socialLinks && (
							<nav
								className="flex items-center gap-3 mt-4"
								aria-label="Social media links"
							>
								{profileData.socialLinks.instagram && (
									<IconButton
										href={profileData.socialLinks.instagram}
										target="_blank"
										rel="noopener noreferrer"
										variant="blue"
										aria-label="Instagram"
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
									>
										<Music className="h-4 w-4" />
									</IconButton>
								)}
							</nav>
						)}
					</Card>
				</div>

				{/* DJ Profile Section */}
				{profileData.isDJ && profileData.upcomingSets && (
					<div>
						<Card className="p-6">
							<h3 className="text-lg font-semibold text-muted-foreground mb-4">
								Upcoming Sets
							</h3>
							<ul className="space-y-3">
								{profileData.upcomingSets.map((set) => (
									<li
										key={set.id}
										className="p-3 bg-secondary/50 rounded-lg border border-gray-800/50"
									>
										<p className="font-medium text-foreground">{set.event}</p>
										<div className="flex items-center gap-4 text-sm text-muted-foreground mt-1">
											<time className="flex items-center gap-1">
												<Calendar className="h-3 w-3" aria-hidden="true" />
												<span>{set.date}</span>
											</time>
											<address className="flex items-center gap-1 not-italic">
												<MapPin className="h-3 w-3" aria-hidden="true" />
												<span>{set.venue}</span>
											</address>
										</div>
									</li>
								))}
							</ul>
						</Card>
					</div>
				)}

				{/* Settings */}
				<div>
					<Card className="p-6">
						<h3 className="text-lg font-semibold text-muted-foreground mb-4">
							Settings
						</h3>
						<div className="space-y-4">
							<div className="flex items-center justify-between">
								<div className="flex items-center gap-3">
									<Bell className="h-5 w-5 text-blue-500" aria-hidden="true" />
									<div>
										<p className="font-medium text-foreground">
											Push Notifications
										</p>
										<p className="text-sm text-muted-foreground">
											Get notified about new events
										</p>
									</div>
								</div>
								<Switch
									checked={profileData.settings.notifications}
									onCheckedChange={handleToggleNotifications}
								/>
							</div>
						</div>
					</Card>
				</div>

				{/* Subscribed Events */}
				<div>
					<Card className="p-6">
						<h3 className="text-lg font-semibold text-muted-foreground mb-4">
							My Events
						</h3>
						<ul className="space-y-3">
							{profileData.subscribedEvents.map((event) => (
								<motion.li
									key={event.id}
									whileHover={{ x: 5 }}
									className="flex items-center gap-3 p-3 bg-secondary/50 rounded-lg border border-gray-800/50 hover:bg-secondary/70 transition-colors"
								>
									<Image
										src={event.imageUrl}
										alt={event.title}
										width={48}
										height={48}
										className="w-12 h-12 rounded-lg object-cover border border-blue-500/20"
									/>
									<div className="flex-1 min-w-0">
										<p className="font-medium text-foreground truncate">
											{event.title}
										</p>
										<div className="flex items-center gap-4 text-sm text-muted-foreground">
											<time>{event.date}</time>
											<address className="not-italic truncate">
												{event.location}
											</address>
										</div>
									</div>
								</motion.li>
							))}
						</ul>
					</Card>
				</div>
			</div>
		</div>
	);
}
