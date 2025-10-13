import { cn } from "@/lib/cn";

interface SwitchProps {
	checked: boolean;
	onCheckedChange: (checked: boolean) => void;
	label?: string;
	className?: string;
}

export function Switch({
	checked,
	onCheckedChange,
	label,
	className,
}: SwitchProps) {
	return (
		<label
			className={cn(
				"relative inline-flex items-center cursor-pointer",
				className,
			)}
		>
			<input
				type="checkbox"
				checked={checked}
				onChange={(e) => onCheckedChange(e.target.checked)}
				className="sr-only peer"
			/>
			<div className="w-11 h-6 bg-gray-700 peer-checked:bg-blue-500 rounded-full transition-colors peer-focus:ring-2 peer-focus:ring-blue-500 peer-focus:ring-offset-2 peer-focus:ring-offset-black">
				<div className="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full transition-transform peer-checked:translate-x-5" />
			</div>
			{label && <span className="ml-3 text-sm font-medium">{label}</span>}
		</label>
	);
}
