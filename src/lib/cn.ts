import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

/**
 * Utility function to merge Tailwind classes
 * Handles conflicts and duplicates properly
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
