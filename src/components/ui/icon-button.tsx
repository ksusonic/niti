import { type HTMLMotionProps, motion } from "framer-motion";
import { cn } from "@/lib/cn";

interface IconButtonProps extends Omit<HTMLMotionProps<"a">, "ref"> {
	children: React.ReactNode;
	variant?: "blue" | "orange" | "green" | "default";
}

const variantClasses = {
	blue: "bg-blue-600/20 text-blue-400 hover:bg-blue-600/30 focus:ring-blue-500",
	orange:
		"bg-orange-500/20 text-orange-400 hover:bg-orange-500/30 focus:ring-orange-500",
	green:
		"bg-green-500/20 text-green-400 hover:bg-green-500/30 focus:ring-green-500",
	default:
		"bg-gray-500/20 text-gray-400 hover:bg-gray-500/30 focus:ring-gray-500",
};

export function IconButton({
	children,
	variant = "default",
	className,
	...props
}: IconButtonProps) {
	return (
		<motion.a
			whileHover={{ scale: 1.1 }}
			whileTap={{ scale: 0.95 }}
			className={cn(
				"p-1.5 rounded-full transition-colors focus:outline-none focus:ring-2",
				variantClasses[variant],
				className,
			)}
			{...props}
		>
			{children}
		</motion.a>
	);
}
