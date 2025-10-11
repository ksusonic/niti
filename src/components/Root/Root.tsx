'use client';

import { type PropsWithChildren, useEffect, useState } from 'react';
import {
  initData,
  miniApp,
  useLaunchParams,
  useSignal,
} from '@telegram-apps/sdk-react';
import { TonConnectUIProvider } from '@tonconnect/ui-react';
import { AppRoot } from '@telegram-apps/telegram-ui';

import { ErrorBoundary } from '@/components/ErrorBoundary';
import { ErrorPage } from '@/components/ErrorPage';
import { EnvUnsupported } from '@/components/EnvUnsupported';
import { TelegramEnvErrorBoundary } from '@/components/TelegramEnvErrorBoundary';
import { useDidMount } from '@/hooks/useDidMount';
import { getIsEnvUnsupported } from '@/instrumentation-client';

import './styles.css';

function RootInner({ children }: PropsWithChildren) {
  const lp = useLaunchParams();
  const isDark = useSignal(miniApp.isDark);
  const initDataUser = useSignal(initData.user);

  useEffect(() => {
  }, [initDataUser]);

  return (
    <TonConnectUIProvider manifestUrl="/tonconnect-manifest.json">
      <AppRoot
        appearance={isDark ? 'dark' : 'light'}
        platform={
          ['macos', 'ios'].includes(lp.tgWebAppPlatform) ? 'ios' : 'base'
        }
      >
        {children}
      </AppRoot>
    </TonConnectUIProvider>
  );
}

export function Root(props: PropsWithChildren) {
  // Unfortunately, Telegram Mini Apps does not allow us to use all features of
  // the Server Side Rendering. That's why we are showing loader on the server
  // side.
  const didMount = useDidMount();
  const [envUnsupported, setEnvUnsupported] = useState(false);

  useEffect(() => {
    // Check initial state
    if (getIsEnvUnsupported()) {
      setEnvUnsupported(true);
    }

    // Listen for the custom event
    const handleEnvUnsupported = () => {
      setEnvUnsupported(true);
    };

    if (typeof window !== 'undefined') {
      window.addEventListener('telegram-env-unsupported', handleEnvUnsupported);
      return () => {
        window.removeEventListener('telegram-env-unsupported', handleEnvUnsupported);
      };
    }
  }, []);

  if (!didMount) {
    return <div className="root__loading">Loading</div>;
  }

  if (envUnsupported) {
    return <EnvUnsupported />;
  }

  return (
    <ErrorBoundary fallback={ErrorPage}>
      <TelegramEnvErrorBoundary>
        <RootInner {...props} />
      </TelegramEnvErrorBoundary>
    </ErrorBoundary>
  );
}
