const ALLOWED_VIDEO_DOMAINS = [
  'youtube.com',
  'www.youtube.com',
  'youtu.be',
  'vimeo.com',
  'player.vimeo.com',
  'dailymotion.com',
  'www.dailymotion.com',
  'twitch.tv',
  'www.twitch.tv',
  'player.twitch.tv',
];

export function isValidVideoUrl(url: string): boolean {
  try {
    const parsedUrl = new URL(url);
    
    if (parsedUrl.protocol !== 'https:') {
      return false;
    }
    
    return ALLOWED_VIDEO_DOMAINS.some(domain => 
      parsedUrl.hostname === domain || parsedUrl.hostname.endsWith(`.${domain}`)
    );
  } catch {
    return false;
  }
}

export function sanitizeVideoUrl(url: string | undefined): string | null {
  if (!url) {
    return null;
  }
  
  if (!isValidVideoUrl(url)) {
    console.warn(`Invalid video URL blocked: ${url}`);
    return null;
  }
  
  return url;
}
