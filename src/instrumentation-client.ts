// This file is normally used for setting up analytics and other
// services that require one-time initialization on the client.

import { retrieveLaunchParams, isTMA } from '@telegram-apps/sdk-react';
import { init } from './core/init';
import { mockEnv } from './mockEnv';

let isEnvUnsupported = false;

export function getIsEnvUnsupported(): boolean {
  return isEnvUnsupported;
}

if (typeof window !== 'undefined') {
  mockEnv().then(() => {
    // After mocking, check if we're in Telegram environment
    isTMA('complete').then((isTelegramEnv) => {
      if (!isTelegramEnv) {
        isEnvUnsupported = true;
        window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
        return;
      }

      try {
        const launchParams = retrieveLaunchParams();
        const { tgWebAppPlatform: platform } = launchParams;
        const debug = process.env.NODE_ENV === 'development';

        // Configure all application dependencies.
        init({
          debug,
          eruda: debug && ['ios', 'android'].includes(platform),
          mockForMacOS: platform === 'macos',
        });
      } catch (e) {
        if (e && typeof e === 'object' && 'name' in e && e.name === 'LaunchParamsRetrieveError') {
          isEnvUnsupported = true;
          window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
        }
      }
    }).catch(() => {
      isEnvUnsupported = true;
      window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
    });
  }).catch(() => {
    // If mocking fails, try to check if we're in real Telegram environment
    isTMA('complete').then((isTelegramEnv) => {
      if (!isTelegramEnv) {
        isEnvUnsupported = true;
        window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
        return;
      }

      try {
        const launchParams = retrieveLaunchParams();
        const { tgWebAppPlatform: platform } = launchParams;
        const debug = process.env.NODE_ENV === 'development';

        init({
          debug,
          eruda: debug && ['ios', 'android'].includes(platform),
          mockForMacOS: platform === 'macos',
        });
      } catch (e) {
        isEnvUnsupported = true;
        window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
      }
    }).catch(() => {
      isEnvUnsupported = true;
      window.dispatchEvent(new CustomEvent('telegram-env-unsupported'));
    });
  });
}
