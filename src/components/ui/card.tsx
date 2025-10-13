import { cn } from "@/lib/cn";

interface CardProps {
	children: React.ReactNode;
	className?: string;
}

export function Card({ children, className }: CardProps) {
	return (
		<section
			className={cn(
				"bg-card backdrop-blur-sm rounded-xl border border-gray-800/30 shadow-lg",
				className,
			)}
		>
			{children}
		</section>
	);
}
