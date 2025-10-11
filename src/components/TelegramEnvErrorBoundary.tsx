import { Component, type ReactNode, type PropsWithChildren } from 'react';
import { EnvUnsupported } from '@/components/EnvUnsupported';
import { isLaunchParamsRetrieveError } from '@telegram-apps/sdk-react';

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

  static getDerivedStateFromError(error: Error): TelegramEnvErrorBoundaryState | null {
    // Check if this is a LaunchParamsRetrieveError
    if (isLaunchParamsRetrieveError(error)) {
      console.warn('TelegramEnvErrorBoundary caught LaunchParamsRetrieveError:', error.message);
      return { hasLaunchParamsError: true };
    }
    
    // For other errors, don't update state (let them bubble up)
    return null;
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    if (isLaunchParamsRetrieveError(error)) {
      console.error('LaunchParamsRetrieveError caught in error boundary:', {
        error: error.message,
        stack: error.stack,
        componentStack: errorInfo.componentStack
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