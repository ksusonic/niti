export interface DJ {
	id: string;
	name: string;
	avatar: string;
	time: string;
	social: {
		instagram?: string;
		soundcloud?: string;
		spotify?: string;
	};
}

export interface Event {
	id: string;
	title: string;
	description: string;
	location: string;
	videoUrl?: string;
	imageUrl: string;
	djLineup: DJ[];
	participantCount: number;
	isSubscribed: boolean;
	date: string;
	time: string;
}

export interface UserProfile {
	username: string;
	avatar: string;
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
