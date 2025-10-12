import { useState } from 'react';
import Image from 'next/image';
import { cn } from '@/lib/cn';

interface AvatarProps {
  src: string;
  alt: string;
  className?: string;
  size?: 'sm' | 'md' | 'lg';
}

const sizeClasses = {
  sm: 'h-8 w-8 text-xs',
  md: 'h-10 w-10 text-sm',
  lg: 'h-12 w-12 text-base',
};

const sizeDimensions = {
  sm: 32,
  md: 40,
  lg: 48,
};

export function Avatar({ src, alt, className, size = 'md' }: AvatarProps) {
  const [error, setError] = useState(false);
  const fallback = alt
    .split(' ')
    .map((n) => n[0])
    .join('')
    .toUpperCase()
    .slice(0, 2);

  if (!src || error) {
    return (
      <div
        className={cn(
          'flex items-center justify-center rounded-full bg-gradient-to-br from-blue-500 to-purple-600 text-white font-semibold',
          sizeClasses[size],
          className
        )}
      >
        {fallback}
      </div>
    );
  }

  return (
    <Image
      src={src}
      alt={alt}
      width={sizeDimensions[size]}
      height={sizeDimensions[size]}
      onError={() => setError(true)}
      className={cn('rounded-full object-cover', sizeClasses[size], className)}
    />
  );
}
