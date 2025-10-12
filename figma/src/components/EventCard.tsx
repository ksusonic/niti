import { motion } from 'framer-motion';
import { MapPin, Users, ExternalLink } from 'lucide-react';
import { useState } from 'react';
import Image from 'next/image';

function Avatar({ src, alt, className = '' }: { src: string; alt: string; className?: string }) {
  const [error, setError] = useState(false);
  const fallback = alt.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2);
  
  if (!src || error) {
    return (
      <div className={`flex items-center justify-center rounded-full bg-gradient-to-br from-blue-500 to-purple-600 text-white font-semibold ${className}`}>
        {fallback}
      </div>
    );
  }
  
  return (
    <Image 
      src={src} 
      alt={alt} 
      width={40}
      height={40}
      onError={() => setError(true)} 
      className={`rounded-full object-cover ${className}`}
    />
  );
}

function Badge({ children, className = '' }: { children: React.ReactNode; className?: string }) {
  return <span className={`inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-xs font-medium ${className}`}>{children}</span>;
}

function Button({ 
  children, 
  onClick, 
  className = '',
  ...props 
}: React.ButtonHTMLAttributes<HTMLButtonElement>) {
  return (
    <button
      onClick={onClick}
      className={`inline-flex items-center justify-center gap-2 px-4 py-2 rounded-lg font-medium transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}

interface DJ {
  id: string;
  name: string;
  avatar: string;
  time: string;
  social: {
    instagram?: string;
    soundcloud?: string;
    spotify?: string;
  };
}

interface Event {
  id: string;
  title: string;
  description: string;
  location: string;
  videoUrl?: string;
  imageUrl: string;
  djLineup: DJ[];
  participantCount: number;
  isSubscribed: boolean;
  date: string;
  time: string;
}

interface EventCardProps {
  event: Event;
  onToggleSubscription: (eventId: string) => void;
}

export function EventCard({ event, onToggleSubscription }: EventCardProps) {
  return (
    <motion.article
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      whileHover={{ y: -5 }}
      transition={{ duration: 0.3 }}
      className="relative bg-gray-800/50 backdrop-blur-sm rounded-xl overflow-hidden border border-gray-700/50 shadow-lg"
      aria-label={`Event: ${event.title}`}
    >
      {/* Video/Image Background */}
      <div className="relative h-64 overflow-hidden">
        {event.videoUrl ? (
          <>
            <iframe
              src={event.videoUrl}
              className="absolute inset-0 w-full h-full object-cover pointer-events-none"
              allow="autoplay; fullscreen"
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black via-black/40 to-transparent" />
          </>
        ) : (
          <>
            <div 
              className="absolute inset-0 bg-cover bg-center blur-sm scale-110 opacity-30"
              style={{ backgroundImage: `url(${event.imageUrl})` }}
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black via-black/40 to-transparent" />
          </>
        )}

        {/* Event Date/Time Badge */}
        <div className="absolute top-4 left-4 z-20">
          <Badge className="bg-blue-600/90 text-white backdrop-blur-md shadow-lg">
            <time dateTime={event.date}>
              {event.date} • {event.time}
            </time>
          </Badge>
        </div>

        {/* Event Title Overlay */}
        <header className="absolute bottom-0 left-0 right-0 z-10 p-6">
          <h3 className="text-2xl font-bold text-white mb-1">{event.title}</h3>
          <address className="flex items-center gap-2 text-sm text-white/70 not-italic">
            <MapPin className="h-4 w-4" aria-hidden="true" />
            <span>{event.location}</span>
          </address>
        </header>
      </div>

      {/* Content */}
      <div className="p-6 space-y-4">
        {/* Event Description */}
        <p className="text-gray-300 text-sm leading-relaxed">{event.description}</p>

        {/* DJ Lineup */}
        <section className="space-y-3" aria-labelledby="lineup-heading">
          <h4 id="lineup-heading" className="text-sm font-medium text-white/70">
            Lineup
          </h4>
          <ul className="space-y-2" role="list">
            {event.djLineup.map((dj) => (
              <motion.li
                key={dj.id}
                whileHover={{ x: 3 }}
                className="flex items-center justify-between bg-gray-700/30 rounded-lg p-3 border border-gray-600/30"
              >
                <div className="flex items-center gap-3 flex-1">
                  <Avatar
                    src={dj.avatar}
                    alt={dj.name}
                    className="h-10 w-10 border border-blue-500/30"
                  />
                  <div className="flex-1 min-w-0">
                    <span className="font-medium block truncate">{dj.name}</span>
                    <time className="text-xs text-gray-400">{dj.time}</time>
                  </div>
                </div>
                
                {/* Social Links */}
                <nav className="flex items-center gap-1.5" aria-label={`${dj.name} social links`}>
                  {dj.social.instagram && (
                    <motion.a
                      href={dj.social.instagram}
                      target="_blank"
                      rel="noopener noreferrer"
                      whileHover={{ scale: 1.1 }}
                      whileTap={{ scale: 0.95 }}
                      className="p-1.5 rounded-full bg-blue-600/20 text-blue-400 hover:bg-blue-600/30 transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500"
                      aria-label={`${dj.name} on Instagram`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </motion.a>
                  )}
                  {dj.social.soundcloud && (
                    <motion.a
                      href={dj.social.soundcloud}
                      target="_blank"
                      rel="noopener noreferrer"
                      whileHover={{ scale: 1.1 }}
                      whileTap={{ scale: 0.95 }}
                      className="p-1.5 rounded-full bg-orange-500/20 text-orange-400 hover:bg-orange-500/30 transition-colors focus:outline-none focus:ring-2 focus:ring-orange-500"
                      aria-label={`${dj.name} on SoundCloud`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </motion.a>
                  )}
                  {dj.social.spotify && (
                    <motion.a
                      href={dj.social.spotify}
                      target="_blank"
                      rel="noopener noreferrer"
                      whileHover={{ scale: 1.1 }}
                      whileTap={{ scale: 0.95 }}
                      className="p-1.5 rounded-full bg-green-500/20 text-green-400 hover:bg-green-500/30 transition-colors focus:outline-none focus:ring-2 focus:ring-green-500"
                      aria-label={`${dj.name} on Spotify`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </motion.a>
                  )}
                </nav>
              </motion.li>
            ))}
          </ul>
        </section>

        {/* Participants and Subscribe Button */}
        <footer className="flex items-center justify-between pt-4 border-t border-gray-600/30">
          <div className="flex items-center gap-2 text-sm text-gray-400">
            <Users className="h-4 w-4" aria-hidden="true" />
            <span>
              <strong className="font-medium text-white">{event.participantCount}</strong> going
            </span>
          </div>
          
          <motion.div
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
          >
            <Button
              onClick={() => onToggleSubscription(event.id)}
              className={
                event.isSubscribed
                  ? 'bg-white text-gray-900 hover:bg-gray-100'
                  : 'bg-blue-600 text-white hover:bg-blue-700'
              }
              aria-pressed={event.isSubscribed}
            >
              {event.isSubscribed ? 'Joined' : 'Join Event'}
            </Button>
          </motion.div>
        </footer>
      </div>
    </motion.article>
  );
}