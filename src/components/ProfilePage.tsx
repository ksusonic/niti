import { motion } from 'framer-motion';
import { Settings, Music, Bell, MapPin, Calendar, Instagram, Music2 } from 'lucide-react';
import { useState } from 'react';
import Image from 'next/image';
import type { UserProfile } from '@/types/events';

function Avatar({ src, alt, className = '' }: { src: string; alt: string; className?: string }) {
  const [error, setError] = useState(false);
  const fallback = alt.charAt(0).toUpperCase();
  
  if (!src || error) {
    return (
      <div className={`flex items-center justify-center rounded-full bg-gradient-to-br from-blue-500 to-purple-600 text-white font-bold ${className}`}>
        {fallback}
      </div>
    );
  }
  
  return (
    <Image 
      src={src} 
      alt={alt} 
      width={80}
      height={80}
      onError={() => setError(true)} 
      className={`rounded-full object-cover ${className}`}
    />
  );
}

function Badge({ children, className = '' }: { children: React.ReactNode; className?: string }) {
  return <span className={`inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-xs font-medium ${className}`}>{children}</span>;
}

function Card({ children, className = '' }: { children: React.ReactNode; className?: string }) {
  return <section className={`bg-gray-800/50 backdrop-blur-sm rounded-xl border border-gray-700/50 shadow-lg ${className}`}>{children}</section>;
}

function Switch({ checked, onCheckedChange, label }: { checked: boolean; onCheckedChange: (checked: boolean) => void; label?: string }) {
  return (
    <label className="relative inline-flex items-center cursor-pointer">
      <input
        type="checkbox"
        checked={checked}
        onChange={(e) => onCheckedChange(e.target.checked)}
        className="sr-only peer"
      />
      <div className="w-11 h-6 bg-gray-600 peer-checked:bg-blue-600 rounded-full transition-colors peer-focus:ring-2 peer-focus:ring-blue-500 peer-focus:ring-offset-2 peer-focus:ring-offset-gray-900">
        <div className="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full transition-transform peer-checked:translate-x-5" />
      </div>
      {label && <span className="ml-3 text-sm font-medium">{label}</span>}
    </label>
  );
}

interface ProfilePageProps {
  profile: UserProfile;
  onUpdateProfile: (updates: Partial<UserProfile>) => void;
}

export function ProfilePage({ profile, onUpdateProfile }: ProfilePageProps) {
  const handleToggleNotifications = () => {
    onUpdateProfile({
      settings: {
        ...profile.settings,
        notifications: !profile.settings.notifications
      }
    });
  };

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <motion.header 
        initial={{ opacity: 0, y: -20 }}
        animate={{ opacity: 1, y: 0 }}
        className="sticky top-0 z-50 bg-gray-900/95 backdrop-blur-lg border-b border-gray-700/50 px-4 py-4"
      >
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold text-white">
            Profile
          </h1>
          <button 
            className="p-2 hover:bg-gray-800 rounded-lg transition-colors"
            aria-label="Settings"
          >
            <Settings className="h-6 w-6 text-gray-400" />
          </button>
        </div>
      </motion.header>

      <div className="p-4 space-y-6 pb-24">
        {/* Profile Header */}
        <motion.div
          initial={{ opacity: 0, scale: 0.95 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.5 }}
        >
          <Card className="p-6">
            <div className="flex items-center gap-4">
              <Avatar
                src={profile.avatar}
                alt={profile.username}
                className="h-20 w-20 border-2 border-blue-500/50"
              />
              
              <div className="flex-1">
                <div className="flex items-center gap-2">
                  <h2 className="text-xl font-bold text-white">{profile.username}</h2>
                  {profile.isDJ && (
                    <Badge className="bg-blue-600/20 text-blue-400 border border-blue-500/30">
                      <Music className="h-3 w-3 mr-1" />
                      DJ
                    </Badge>
                  )}
                </div>
                {profile.bio && (
                  <p className="text-gray-400 text-sm mt-1">{profile.bio}</p>
                )}
              </div>
            </div>

            {/* Social Links */}
            {profile.socialLinks && (
              <nav className="flex items-center gap-3 mt-4" aria-label="Social media links">
                {profile.socialLinks.instagram && (
                  <motion.a
                    href={profile.socialLinks.instagram}
                    target="_blank"
                    rel="noopener noreferrer"
                    whileHover={{ scale: 1.1 }}
                    whileTap={{ scale: 0.95 }}
                    className="p-2 rounded-full bg-blue-600/20 text-blue-400 hover:bg-blue-600/30 transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500"
                    aria-label="Instagram"
                  >
                    <Instagram className="h-4 w-4" />
                  </motion.a>
                )}
                {profile.socialLinks.soundcloud && (
                  <motion.a
                    href={profile.socialLinks.soundcloud}
                    target="_blank"
                    rel="noopener noreferrer"
                    whileHover={{ scale: 1.1 }}
                    whileTap={{ scale: 0.95 }}
                    className="p-2 rounded-full bg-orange-500/20 text-orange-400 hover:bg-orange-500/30 transition-colors focus:outline-none focus:ring-2 focus:ring-orange-500"
                    aria-label="SoundCloud"
                  >
                    <Music2 className="h-4 w-4" />
                  </motion.a>
                )}
                {profile.socialLinks.spotify && (
                  <motion.a
                    href={profile.socialLinks.spotify}
                    target="_blank"
                    rel="noopener noreferrer"
                    whileHover={{ scale: 1.1 }}
                    whileTap={{ scale: 0.95 }}
                    className="p-2 rounded-full bg-green-500/20 text-green-400 hover:bg-green-500/30 transition-colors focus:outline-none focus:ring-2 focus:ring-green-500"
                    aria-label="Spotify"
                  >
                    <Music className="h-4 w-4" />
                  </motion.a>
                )}
              </nav>
            )}
          </Card>
        </motion.div>

        {/* DJ Profile Section */}
        {profile.isDJ && profile.upcomingSets && (
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, delay: 0.1 }}
          >
            <Card className="p-6">
              <h3 className="text-lg font-semibold text-white/90 mb-4">Upcoming Sets</h3>
              <ul className="space-y-3" role="list">
                {profile.upcomingSets.map((set) => (
                  <li key={set.id} className="p-3 bg-gray-700/50 rounded-lg">
                    <p className="font-medium text-white">{set.event}</p>
                    <div className="flex items-center gap-4 text-sm text-gray-400 mt-1">
                      <time className="flex items-center gap-1">
                        <Calendar className="h-3 w-3" aria-hidden="true" />
                        <span>{set.date}</span>
                      </time>
                      <address className="flex items-center gap-1 not-italic">
                        <MapPin className="h-3 w-3" aria-hidden="true" />
                        <span>{set.venue}</span>
                      </address>
                    </div>
                  </li>
                ))}
              </ul>
            </Card>
          </motion.div>
        )}

        {/* Settings */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.2 }}
        >
          <Card className="p-6">
            <h3 className="text-lg font-semibold text-white/90 mb-4">Settings</h3>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <div className="flex items-center gap-3">
                  <Bell className="h-5 w-5 text-blue-400" aria-hidden="true" />
                  <div>
                    <p className="font-medium text-white">Push Notifications</p>
                    <p className="text-sm text-gray-400">Get notified about new events</p>
                  </div>
                </div>
                <Switch
                  checked={profile.settings.notifications}
                  onCheckedChange={handleToggleNotifications}
                />
              </div>
            </div>
          </Card>
        </motion.div>

        {/* Subscribed Events */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.3 }}
        >
          <Card className="p-6">
            <h3 className="text-lg font-semibold text-white/90 mb-4">My Events</h3>
            <ul className="space-y-3" role="list">
              {profile.subscribedEvents.map((event) => (
                <motion.li
                  key={event.id}
                  whileHover={{ x: 5 }}
                  className="flex items-center gap-3 p-3 bg-gray-700/50 rounded-lg hover:bg-gray-700/70 transition-colors"
                >
                  <Image
                    src={event.imageUrl}
                    alt={event.title}
                    width={48}
                    height={48}
                    className="w-12 h-12 rounded-lg object-cover border border-blue-500/30"
                  />
                  <div className="flex-1 min-w-0">
                    <p className="font-medium text-white truncate">{event.title}</p>
                    <div className="flex items-center gap-4 text-sm text-gray-400">
                      <time>{event.date}</time>
                      <address className="not-italic truncate">{event.location}</address>
                    </div>
                  </div>
                </motion.li>
              ))}
            </ul>
          </Card>
        </motion.div>
      </div>
    </div>
  );
}
