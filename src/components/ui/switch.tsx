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
		<label className={cn("inline-flex items-center cursor-pointer", className)}>
			<input
				type="checkbox"
				checked={checked}
				onChange={(e) => onCheckedChange(e.target.checked)}
				className="sr-only peer"
			/>
			<style>{`
				.switch-track {
					transition: background-color 0.2s ease;
				}
				.switch-thumb {
					transition: transform 0.2s cubic-bezier(0.4, 0.0, 0.2, 1);
				}
				input[type="checkbox"]:checked ~ .switch-track .switch-thumb {
					transform: translateX(20px);
				}
			`}</style>
			<div
				className="switch-track relative w-11 h-6 rounded-full peer-focus:ring-2 peer-focus:ring-blue-500 peer-focus:ring-offset-2 peer-focus:ring-offset-black"
				style={{
					backgroundColor: checked ? "#3b82f6" : "#374151",
				}}
			>
				<div className="switch-thumb absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full pointer-events-none" />
			</div>
			{label && <span className="ml-3 text-sm font-medium">{label}</span>}
		</label>
	);
}
