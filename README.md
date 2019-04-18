# app-redirector

This is a very simple server that redirects Android/iOS devices to their respective App Store/Play Store pages.

By default, it defaults to the YouTube iOS/Android apps. This is probably not useful. Use these environment variables to customize these pages:

- `PLAY_STORE_PAGE` - The page to redirect Android devices to
- `APP_STORE_PAGE` - The page to redirect iOS devices to
- `UNKNOWN_PAGE` - The page to redirect everything else to