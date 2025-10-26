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
		label: "События",
		icon: Calendar,
	},
	{
		id: "profile",
		label: "Профиль",
		icon: User,
	},
];

const BottomNavigationComponent = ({
	activeTab,
	onTabChange,
}: BottomNavigationProps) => {
	return (
		<nav
			className="fixed bottom-0 left-0 right-0 z-50 bg-gradient-to-t from-black via-black/98 to-black/95 backdrop-blur-2xl border-t border-blue-500/20 shadow-2xl safe-area-inset-bottom"
			aria-label="Main navigation"
		>
			<div className="relative max-w-md mx-auto">
				<div className="flex items-center justify-around px-6 py-3 safe-area-inset-bottom">
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
									"relative flex flex-col items-center gap-2 px-10 py-3 rounded-2xl transition-all duration-300",
									"active:scale-95 touch-manipulation",
									"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500/50",
									isActive
										? "text-white bg-gradient-to-br from-blue-500/30 to-purple-500/20 shadow-lg shadow-blue-500/20 border border-blue-500/30"
										: "text-gray-400 hover:text-white hover:bg-gray-800/30",
								)}
							>
								<div className="relative">
									<Icon
										className={cn(
											"relative h-6 w-6 transition-all duration-300",
											isActive &&
												"scale-110 drop-shadow-[0_0_8px_rgba(59,130,246,0.8)]",
										)}
										strokeWidth={isActive ? 2.5 : 2}
										aria-hidden="true"
									/>
									{isActive && (
										<div className="absolute inset-0 bg-blue-500/20 blur-xl rounded-full" />
									)}
								</div>

								<span
									className={cn(
										"relative text-xs font-semibold tracking-wide transition-all duration-300",
										isActive ? "text-white" : "text-gray-400",
									)}
								>
									{tab.label}
								</span>

								{isActive && (
									<div className="absolute -bottom-1 left-1/2 -translate-x-1/2 w-12 h-1 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full" />
								)}
							</button>
						);
					})}
				</div>
			</div>
		</nav>
	);
};

export const BottomNavigation = memo(BottomNavigationComponent);
