interface ErrorStateProps {
	error: string;
	onRetry?: () => void;
}

export function ErrorState({ error, onRetry }: ErrorStateProps) {
	const handleRetry = () => {
		if (onRetry) {
			onRetry();
		} else {
			window.location.reload();
		}
	};

	return (
		<div className="flex flex-col items-center justify-center min-h-screen px-6 text-center">
			{/* Broken DJ Controller SVG */}
			<div className="relative mb-6">
				<svg
					width="200"
					height="150"
					viewBox="0 0 200 150"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					className="opacity-80"
					role="img"
					aria-label="Broken DJ Controller"
				>
					<title>Broken DJ Controller</title>
					{/* Controller body */}
					<rect
						x="20"
						y="40"
						width="160"
						height="80"
						rx="8"
						fill="currentColor"
						className="text-muted-foreground/30"
						stroke="currentColor"
						strokeWidth="2"
					/>

					{/* Left deck */}
					<circle
						cx="60"
						cy="80"
						r="20"
						fill="currentColor"
						className="text-muted-foreground/50"
					/>
					<circle
						cx="60"
						cy="80"
						r="15"
						fill="currentColor"
						className="text-background"
					/>

					{/* Right deck */}
					<circle
						cx="140"
						cy="80"
						r="20"
						fill="currentColor"
						className="text-muted-foreground/50"
					/>
					<circle
						cx="140"
						cy="80"
						r="15"
						fill="currentColor"
						className="text-background"
					/>

					{/* Knobs */}
					<circle
						cx="50"
						cy="55"
						r="4"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<circle
						cx="70"
						cy="55"
						r="4"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<circle
						cx="130"
						cy="55"
						r="4"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<circle
						cx="150"
						cy="55"
						r="4"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>

					{/* Faders */}
					<rect
						x="48"
						y="100"
						width="4"
						height="12"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<rect
						x="68"
						y="100"
						width="4"
						height="12"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<rect
						x="128"
						y="100"
						width="4"
						height="12"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>
					<rect
						x="148"
						y="100"
						width="4"
						height="12"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/60"
					/>

					{/* Crossfader */}
					<rect
						x="80"
						y="105"
						width="40"
						height="4"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/40"
					/>
					<rect
						x="95"
						y="103"
						width="10"
						height="8"
						rx="2"
						fill="currentColor"
						className="text-muted-foreground/70"
					/>

					{/* Error crack effect */}
					<path
						d="M 100 30 L 95 50 L 105 60 L 100 80 L 110 90 L 105 110"
						stroke="#ef4444"
						strokeWidth="2"
						strokeLinecap="round"
						strokeLinejoin="round"
						className="animate-pulse"
					/>

					{/* Sparks */}
					<circle
						cx="95"
						cy="50"
						r="2"
						fill="#ef4444"
						className="animate-pulse"
					/>
					<circle
						cx="105"
						cy="60"
						r="2"
						fill="#fbbf24"
						className="animate-pulse"
						style={{ animationDelay: "0.2s" }}
					/>
					<circle
						cx="110"
						cy="90"
						r="2"
						fill="#ef4444"
						className="animate-pulse"
						style={{ animationDelay: "0.4s" }}
					/>
				</svg>
			</div>

			<h2 className="text-2xl font-bold mb-2 bg-gradient-to-r from-red-500 to-orange-500 bg-clip-text text-transparent">
				Connection Dropped!
			</h2>
			<p className="text-muted-foreground mb-2 max-w-sm">
				The beat stopped... Looks like we lost connection to the decks.
			</p>
			<p className="text-sm text-muted-foreground/80 mb-6 font-mono">{error}</p>
			<button
				type="button"
				onClick={handleRetry}
				className="px-8 py-3 bg-gradient-to-r from-purple-500 to-pink-500 text-white font-semibold rounded-lg hover:from-purple-600 hover:to-pink-600 transition-all transform hover:scale-105 shadow-lg"
			>
				🎧 Reconnect
			</button>
		</div>
	);
}
