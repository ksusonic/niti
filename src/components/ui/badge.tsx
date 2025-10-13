import { cn } from "@/lib/cn";

interface BadgeProps {
	children: React.ReactNode;
	variant?: "default" | "primary" | "secondary";
	className?: string;
}

const variantClasses = {
	default: "bg-gray-500/20 text-gray-300",
	primary: "bg-blue-500 text-white backdrop-blur-md shadow-lg",
	secondary: "bg-secondary/50 text-muted-foreground",
};

export function Badge({
	children,
	variant = "default",
	className,
}: BadgeProps) {
	return (
		<span
			className={cn(
				"inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-xs font-medium",
				variantClasses[variant],
				className,
			)}
		>
			{children}
		</span>
	);
}
