import { motion } from 'framer-motion';
import { cn } from '@/lib/cn';
import { forwardRef } from 'react';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  variant?: 'default' | 'primary' | 'secondary' | 'ghost';
  size?: 'sm' | 'md' | 'lg';
  animated?: boolean;
}

const variantClasses = {
  default: 'bg-gray-800 text-white hover:bg-gray-700',
  primary: 'bg-blue-500 text-white hover:bg-blue-600',
  secondary: 'bg-white text-gray-900 hover:bg-gray-100',
  ghost: 'bg-transparent text-gray-300 hover:bg-gray-800/50',
};

const sizeClasses = {
  sm: 'px-3 py-1.5 text-sm',
  md: 'px-4 py-2 text-base',
  lg: 'px-6 py-3 text-lg',
};

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(function Button({
  children,
  variant = 'default',
  size = 'md',
  animated = true,
  className,
  onClick,
  disabled,
  type,
  ...props
}, ref) {
  const buttonClasses = cn(
    'inline-flex items-center justify-center gap-2 rounded-lg font-medium transition-all duration-200',
    'focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500',
    'disabled:opacity-50 disabled:cursor-not-allowed',
    variantClasses[variant],
    sizeClasses[size],
    className
  );

  if (animated) {
    return (
      <motion.button
        ref={ref}
        whileHover={{ scale: 1.05 }}
        whileTap={{ scale: 0.95 }}
        className={buttonClasses}
        onClick={onClick}
        disabled={disabled}
        type={type}
      >
        {children}
      </motion.button>
    );
  }

  return (
    <button ref={ref} className={buttonClasses} {...props}>
      {children}
    </button>
  );
});
