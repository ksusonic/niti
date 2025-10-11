// This file is normally used for setting up analytics and other
// services that require one-time initialization on the client.

import { retrieveLaunchParams, isTMA, isLaunchParamsRetrieveError } from '@telegram-apps/sdk-react';
import { init } from './core/init';
import { mockEnv } from './mockEnv';
import { setEnvUnsupported } from './lib';

if (typeof window !== 'undefined') {
  mockEnv().then(() => {
    // After mocking, check if we're in Telegram environment
    isTMA('complete').then((isTelegramEnv) => {
      if (!isTelegramEnv) {
        setEnvUnsupported();
        return;
      }

      try {
        const launchParams = retrieveLaunchParams();
        const { tgWebAppPlatform: platform } = launchParams;
        const debug = process.env.NODE_ENV === 'development';

        // Configure all application dependencies.
        init({
          debug,
          mockForMacOS: platform === 'macos',
        });
        console.log('Telegram SDK initialized successfully');
      } catch (e) {
        if (isLaunchParamsRetrieveError(e)) {
          console.warn('LaunchParamsRetrieveError: App running outside Telegram environment');
          setEnvUnsupported();
        } else {
          console.error('Unexpected error during launch params retrieval:', e);
        }
      }
    }).catch((error) => {
      console.warn('isTMA check failed after mocking:', error);
      setEnvUnsupported();
    });
  }).catch((mockError) => {
    console.warn('Mock environment setup failed:', mockError);
    // If mocking fails, try to check if we're in real Telegram environment
    isTMA('complete').then((isTelegramEnv) => {
      if (!isTelegramEnv) {
        setEnvUnsupported();
        return;
      }

      try {
        const launchParams = retrieveLaunchParams();
        const { tgWebAppPlatform: platform } = launchParams;
        const debug = process.env.NODE_ENV === 'development';

        init({
          debug,
          mockForMacOS: platform === 'macos',
        });
        console.log('Telegram SDK initialized successfully (fallback path)');
      } catch (e) {
        console.error('Failed to initialize after mock failure:', e);
        setEnvUnsupported();
      }
    }).catch((error) => {
      console.warn('Final isTMA check failed:', error);
      setEnvUnsupported();
    });
  });
}
