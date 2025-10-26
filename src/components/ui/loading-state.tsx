import { useId } from "react";

interface LoadingStateProps {
	message?: string;
}

export function LoadingState({ message = "Загрузка..." }: LoadingStateProps) {
	const gradientId = useId();
	const filterId = useId();

	return (
		<div className="flex flex-col items-center justify-center min-h-screen px-6">
			{/* Animated DJ Controller */}
			<div className="relative mb-8">
				<svg
					width="240"
					height="180"
					viewBox="0 0 240 180"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					className="opacity-90"
					role="img"
					aria-label="DJ Controller Loading Animation"
				>
					{/* Controller body with glow */}
					<defs>
						<linearGradient id={gradientId} x1="0%" y1="0%" x2="100%" y2="100%">
							<stop
								offset="0%"
								stopColor="rgb(168, 85, 247)"
								stopOpacity="0.2"
							/>
							<stop
								offset="100%"
								stopColor="rgb(236, 72, 153)"
								stopOpacity="0.2"
							/>
						</linearGradient>
						<filter id={filterId}>
							<feGaussianBlur stdDeviation="3" result="coloredBlur" />
							<feMerge>
								<feMergeNode in="coloredBlur" />
								<feMergeNode in="SourceGraphic" />
							</feMerge>
						</filter>
					</defs>

					<rect
						x="30"
						y="50"
						width="180"
						height="90"
						rx="12"
						fill={`url(#${gradientId})`}
						stroke="currentColor"
						strokeWidth="2"
						className="text-purple-500/50"
					/>

					{/* Left deck - spinning */}
					<g className="origin-[75px_95px]">
						<animateTransform
							attributeName="transform"
							attributeType="XML"
							type="rotate"
							from="0 75 95"
							to="360 75 95"
							dur="2s"
							repeatCount="indefinite"
						/>
						<circle
							cx="75"
							cy="95"
							r="25"
							fill="currentColor"
							className="text-purple-500/30"
						/>
						<circle
							cx="75"
							cy="95"
							r="20"
							fill="currentColor"
							stroke="currentColor"
							strokeWidth="2"
							className="text-background"
						/>
						<circle
							cx="75"
							cy="95"
							r="3"
							fill="currentColor"
							className="text-purple-500"
						/>
						<line
							x1="75"
							y1="95"
							x2="75"
							y2="75"
							stroke="currentColor"
							strokeWidth="2"
							className="text-purple-500"
							strokeLinecap="round"
						/>
					</g>

					{/* Right deck - spinning opposite direction */}
					<g className="origin-[165px_95px]">
						<animateTransform
							attributeName="transform"
							attributeType="XML"
							type="rotate"
							from="0 165 95"
							to="-360 165 95"
							dur="2s"
							repeatCount="indefinite"
						/>
						<circle
							cx="165"
							cy="95"
							r="25"
							fill="currentColor"
							className="text-pink-500/30"
						/>
						<circle
							cx="165"
							cy="95"
							r="20"
							fill="currentColor"
							stroke="currentColor"
							strokeWidth="2"
							className="text-background"
						/>
						<circle
							cx="165"
							cy="95"
							r="3"
							fill="currentColor"
							className="text-pink-500"
						/>
						<line
							x1="165"
							y1="95"
							x2="165"
							y2="75"
							stroke="currentColor"
							strokeWidth="2"
							className="text-pink-500"
							strokeLinecap="round"
						/>
					</g>

					{/* Knobs with pulsing animation */}
					<g className="animate-pulse" style={{ animationDuration: "1.5s" }}>
						<circle
							cx="55"
							cy="65"
							r="5"
							fill="currentColor"
							className="text-purple-400/80"
						/>
						<circle
							cx="95"
							cy="65"
							r="5"
							fill="currentColor"
							className="text-purple-400/80"
						/>
					</g>
					<g
						className="animate-pulse"
						style={{ animationDuration: "1.5s", animationDelay: "0.3s" }}
					>
						<circle
							cx="145"
							cy="65"
							r="5"
							fill="currentColor"
							className="text-pink-400/80"
						/>
						<circle
							cx="185"
							cy="65"
							r="5"
							fill="currentColor"
							className="text-pink-400/80"
						/>
					</g>

					{/* Animated faders */}
					<rect
						x="53"
						y="120"
						width="4"
						height="15"
						rx="2"
						fill="currentColor"
						className="text-purple-400/40"
					/>
					<rect
						x="53"
						y="120"
						width="4"
						height="6"
						rx="2"
						fill="currentColor"
						className="text-purple-400"
					>
						<animate
							attributeName="y"
							values="120;129;120"
							dur="2s"
							repeatCount="indefinite"
						/>
					</rect>

					<rect
						x="93"
						y="120"
						width="4"
						height="15"
						rx="2"
						fill="currentColor"
						className="text-purple-400/40"
					/>
					<rect
						x="93"
						y="125"
						width="4"
						height="6"
						rx="2"
						fill="currentColor"
						className="text-purple-400"
					>
						<animate
							attributeName="y"
							values="125;120;125"
							dur="2.5s"
							repeatCount="indefinite"
						/>
					</rect>

					<rect
						x="143"
						y="120"
						width="4"
						height="15"
						rx="2"
						fill="currentColor"
						className="text-pink-400/40"
					/>
					<rect
						x="143"
						y="122"
						width="4"
						height="6"
						rx="2"
						fill="currentColor"
						className="text-pink-400"
					>
						<animate
							attributeName="y"
							values="122;129;122"
							dur="1.8s"
							repeatCount="indefinite"
						/>
					</rect>

					<rect
						x="183"
						y="120"
						width="4"
						height="15"
						rx="2"
						fill="currentColor"
						className="text-pink-400/40"
					/>
					<rect
						x="183"
						y="127"
						width="4"
						height="6"
						rx="2"
						fill="currentColor"
						className="text-pink-400"
					>
						<animate
							attributeName="y"
							values="127;120;127"
							dur="2.2s"
							repeatCount="indefinite"
						/>
					</rect>

					{/* Crossfader - sliding animation */}
					<rect
						x="100"
						y="122"
						width="40"
						height="4"
						rx="2"
						fill="currentColor"
						className="text-purple-400/30"
					/>
					<rect
						x="115"
						y="120"
						width="10"
						height="8"
						rx="2"
						fill="currentColor"
						className="text-gradient-to-r from-purple-500 to-pink-500"
						filter={`url(#${filterId})`}
					>
						<animate
							attributeName="x"
							values="100;130;100"
							dur="3s"
							repeatCount="indefinite"
						/>
					</rect>

					{/* EQ visualization bars */}
					<g opacity="0.6">
						<rect
							x="50"
							y="145"
							width="3"
							height="8"
							rx="1.5"
							fill="currentColor"
							className="text-purple-400"
						>
							<animate
								attributeName="height"
								values="8;15;8"
								dur="0.6s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;138;145"
								dur="0.6s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="56"
							y="145"
							width="3"
							height="12"
							rx="1.5"
							fill="currentColor"
							className="text-purple-400"
						>
							<animate
								attributeName="height"
								values="12;8;15;12"
								dur="0.8s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;149;138;145"
								dur="0.8s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="62"
							y="145"
							width="3"
							height="10"
							rx="1.5"
							fill="currentColor"
							className="text-purple-400"
						>
							<animate
								attributeName="height"
								values="10;14;10"
								dur="0.7s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;141;145"
								dur="0.7s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="68"
							y="145"
							width="3"
							height="11"
							rx="1.5"
							fill="currentColor"
							className="text-purple-400"
						>
							<animate
								attributeName="height"
								values="11;8;13;11"
								dur="0.9s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;149;142;145"
								dur="0.9s"
								repeatCount="indefinite"
							/>
						</rect>

						<rect
							x="169"
							y="145"
							width="3"
							height="9"
							rx="1.5"
							fill="currentColor"
							className="text-pink-400"
						>
							<animate
								attributeName="height"
								values="9;14;9"
								dur="0.65s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;140;145"
								dur="0.65s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="175"
							y="145"
							width="3"
							height="13"
							rx="1.5"
							fill="currentColor"
							className="text-pink-400"
						>
							<animate
								attributeName="height"
								values="13;8;15;13"
								dur="0.75s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;150;138;145"
								dur="0.75s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="181"
							y="145"
							width="3"
							height="10"
							rx="1.5"
							fill="currentColor"
							className="text-pink-400"
						>
							<animate
								attributeName="height"
								values="10;15;10"
								dur="0.85s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;140;145"
								dur="0.85s"
								repeatCount="indefinite"
							/>
						</rect>
						<rect
							x="187"
							y="145"
							width="3"
							height="12"
							rx="1.5"
							fill="currentColor"
							className="text-pink-400"
						>
							<animate
								attributeName="height"
								values="12;9;14;12"
								dur="0.95s"
								repeatCount="indefinite"
							/>
							<animate
								attributeName="y"
								values="145;148;141;145"
								dur="0.95s"
								repeatCount="indefinite"
							/>
						</rect>
					</g>
				</svg>
			</div>

			{/* Loading text with gradient */}
			<h2 className="text-2xl font-bold mb-2 bg-gradient-to-r from-purple-500 via-pink-500 to-purple-500 bg-clip-text text-transparent animate-pulse bg-[length:200%_auto]">
				Крутим диски...
			</h2>
			<p className="text-muted-foreground text-lg">{message}</p>

			{/* Dots animation */}
			<div className="flex gap-2 mt-4">
				<div
					className="w-2 h-2 bg-purple-500 rounded-full animate-bounce"
					style={{ animationDelay: "0ms" }}
				/>
				<div
					className="w-2 h-2 bg-pink-500 rounded-full animate-bounce"
					style={{ animationDelay: "150ms" }}
				/>
				<div
					className="w-2 h-2 bg-purple-500 rounded-full animate-bounce"
					style={{ animationDelay: "300ms" }}
				/>
			</div>
		</div>
	);
}
