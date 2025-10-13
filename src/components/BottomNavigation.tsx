import type { LucideIcon } from "lucide-react";
import { Calendar, User } from "lucide-react";
import { memo } from "react";
import { cn } from "@/lib/cn";

interface BottomNavigationProps {
	activeTab: "events" | "profile";
	onTabChange: (tab: "events" | "profile") => void;
}

interface Tab {
	id: "events" | "profile";
	label: string;
	icon: LucideIcon;
}

const tabs: Tab[] = [
	{
		id: "events",
		label: "Events",
		icon: Calendar,
	},
	{
		id: "profile",
		label: "Profile",
		icon: User,
	},
];

const BottomNavigationComponent = ({
	activeTab,
	onTabChange,
}: BottomNavigationProps) => {
	return (
		<nav
			className="bg-card/95 backdrop-blur-xl border-t border-border/50 shadow-lg"
			aria-label="Main navigation"
		>
			<div className="relative max-w-md mx-auto">
				<div className="flex items-center justify-around px-2 py-2 safe-area-inset-bottom">
					{tabs.map((tab) => {
						const Icon = tab.icon;
						const isActive = activeTab === tab.id;

						return (
							<button
								type="button"
								key={tab.id}
								onClick={() => onTabChange(tab.id)}
								aria-pressed={isActive}
								aria-label={tab.label}
								className={cn(
									"relative flex flex-col items-center gap-1.5 px-8 py-2.5 rounded-2xl transition-all duration-200",
									"active:scale-[0.97] touch-manipulation",
									"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/50",
									isActive
										? "text-primary bg-primary/10"
										: "text-muted-foreground hover:text-foreground/80",
								)}
							>
								<Icon
									className={cn(
										"relative h-5 w-5 transition-transform duration-200",
										isActive && "scale-110",
									)}
									strokeWidth={isActive ? 2.5 : 2}
									aria-hidden="true"
								/>

								<span
									className={cn(
										"relative text-[11px] font-semibold tracking-wide",
										isActive && "text-primary",
									)}
								>
									{tab.label}
								</span>
							</button>
						);
					})}
				</div>
			</div>
		</nav>
	);
};

export const BottomNavigation = memo(BottomNavigationComponent);
