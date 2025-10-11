import { Component, type ReactNode, type PropsWithChildren } from 'react';
import { EnvUnsupported } from '@/components/EnvUnsupported';

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
    if (error.name === 'LaunchParamsRetrieveError') {
      return { hasLaunchParamsError: true };
    }
    
    // For other errors, don't update state (let them bubble up)
    return null;
  }

  componentDidCatch(error: Error, errorInfo: any) {
    if (error.name === 'LaunchParamsRetrieveError') {
      console.log('LaunchParamsRetrieveError caught:', error);
      console.log('Error info:', errorInfo);
    }
  }

  render(): ReactNode {
    if (this.state.hasLaunchParamsError) {
      return <EnvUnsupported />;
    }

    return this.props.children;
  }
}