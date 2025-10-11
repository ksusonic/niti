import {
  setDebug,
  mountBackButton,
  restoreInitData,
  init as initSDK,
  mountMiniAppSync,
  bindThemeParamsCssVars,
  mountViewport,
  bindViewportCssVars,
  mockTelegramEnv,
  type ThemeParams,
  themeParamsState,
  retrieveLaunchParams,
  emitEvent,
  isTMA,
} from '@telegram-apps/sdk-react';

/**
 * Initializes the application and configures its dependencies.
 */
export async function init(options: {
  debug: boolean;
  mockForMacOS: boolean;
}): Promise<void> {
  try {
    // Set @telegram-apps/sdk-react debug mode and initialize it.
    setDebug(options.debug);
    initSDK();

    // Telegram for macOS has a ton of bugs, including cases, when the client doesn't
    // even response to the "web_app_request_theme" method. It also generates an incorrect
    // event for the "web_app_request_safe_area" method.
    if (options.mockForMacOS) {
      let firstThemeSent = false;
      mockTelegramEnv({
        onEvent(event, next) {
          if (event[0] === 'web_app_request_theme') {
            let tp: ThemeParams = {};
            if (firstThemeSent) {
              tp = themeParamsState();
            } else {
              firstThemeSent = true;
              tp ||= retrieveLaunchParams().tgWebAppThemeParams;
            }
            return emitEvent('theme_changed', { theme_params: tp });
          }

          if (event[0] === 'web_app_request_safe_area') {
            return emitEvent('safe_area_changed', {
              left: 0,
              top: 0,
              right: 0,
              bottom: 0,
            });
          }

          next();
        },
      });
    }

    const isTelegramEnv = await isTMA('complete').catch((error) => {
      console.warn('Failed to detect Telegram environment in init:', error);
      return false;
    });
    
    if (isTelegramEnv) {
      try {
        mountBackButton.ifAvailable();
      } catch (error) {
        console.warn('Failed to mount back button:', error);
      }

      try {
        restoreInitData();
      } catch (error) {
        console.warn('Failed to restore init data:', error);
      }

      try {
        if (mountMiniAppSync.isAvailable()) {
          mountMiniAppSync();
          bindThemeParamsCssVars();
        }
      } catch (error) {
        console.warn('Failed to mount mini app or bind theme params:', error);
      }

      try {
        if (mountViewport.isAvailable()) {
          mountViewport().then(() => {
            bindViewportCssVars();
          }).catch((error) => {
            console.warn('Failed to bind viewport CSS vars:', error);
          });
        }
      } catch (error) {
        console.warn('Failed to mount viewport:', error);
      }
    }
  } catch (e) {
    console.error('Critical error during Telegram SDK initialization:', e);
    throw e;
  }
}
