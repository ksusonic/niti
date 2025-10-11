import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'xelene.me',
        pathname: '/telegram.gif',
      },
    ],
  },
};

export default nextConfig;
