import { LaunchParamsRetrieveError } from "@tma.js/sdk-react";
import { Component, type PropsWithChildren, type ReactNode } from "react";
import { EnvUnsupported } from "@/components/EnvUnsupported";

interface TelegramEnvErrorBoundaryState {
	hasLaunchParamsError: boolean;
}

export class TelegramEnvErrorBoundary extends Component<
	PropsWithChildren,
	TelegramEnvErrorBoundaryState
> {
	constructor(props: PropsWithChildren) {
		super(props);
		this.state = { hasLaunchParamsError: false };
	}

	static getDerivedStateFromError(
		error: Error,
	): TelegramEnvErrorBoundaryState | null {
		// Check if this is a LaunchParamsRetrieveError
		if (error instanceof LaunchParamsRetrieveError) {
			console.warn(
				"TelegramEnvErrorBoundary caught LaunchParamsRetrieveError:",
				error.message,
			);
			return { hasLaunchParamsError: true };
		}

		// For other errors, don't update state (let them bubble up)
		return null;
	}

	componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
		if (error instanceof LaunchParamsRetrieveError) {
			console.error("LaunchParamsRetrieveError caught in error boundary:", {
				error: error.message,
				stack: error.stack,
				componentStack: errorInfo.componentStack,
			});
		}
	}

	render(): ReactNode {
		if (this.state.hasLaunchParamsError) {
			return <EnvUnsupported />;
		}

		return this.props.children;
	}
}
