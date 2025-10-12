import { motion } from 'framer-motion';
import { Calendar, User } from 'lucide-react';

interface BottomNavigationProps {
  activeTab: 'events' | 'profile';
  onTabChange: (tab: 'events' | 'profile') => void;
}

export function BottomNavigation({ activeTab, onTabChange }: BottomNavigationProps) {
  const tabs = [
    {
      id: 'events' as const,
      label: 'Events',
      icon: Calendar
    },
    {
      id: 'profile' as const,
      label: 'Profile',
      icon: User
    }
  ];

  return (
    <motion.nav
      initial={{ y: 100 }}
      animate={{ y: 0 }}
      className="fixed bottom-0 left-0 right-0 z-50 bg-gray-800/95 backdrop-blur-lg border-t border-gray-700/50"
      aria-label="Main navigation"
    >
      <div className="flex items-center justify-around px-4 py-3">
        {tabs.map((tab) => {
          const Icon = tab.icon;
          const isActive = activeTab === tab.id;
          
          return (
            <motion.button
              key={tab.id}
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={() => onTabChange(tab.id)}
              aria-pressed={isActive}
              aria-label={tab.label}
              className={`flex flex-col items-center gap-1 px-6 py-2 rounded-xl transition-all duration-300 ${
                isActive 
                  ? 'bg-blue-600/20 text-white' 
                  : 'text-gray-400 hover:text-white'
              }`}
            >
              <Icon className={`h-6 w-6 ${isActive ? 'text-blue-400' : ''}`} aria-hidden="true" />
              <span className={`text-xs font-medium ${isActive ? 'text-white' : ''}`}>
                {tab.label}
              </span>
            </motion.button>
          );
        })}
      </div>
      
      {/* Gradient border effect */}
      <div className="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-blue-500/50 to-transparent" aria-hidden="true" />
    </motion.nav>
  );
}
