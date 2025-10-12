import { motion } from 'framer-motion';
import { MapPin, Users, ExternalLink } from 'lucide-react';
import { Avatar, Badge, Button, IconButton } from '@/components/ui';
import type { Event } from '@/types/events';
import { sanitizeVideoUrl } from '@/lib/video-url-validator';
import { useRef } from 'react';
import confetti from 'canvas-confetti';

interface EventCardProps {
  event: Event;
  onToggleSubscription: (eventId: string) => void;
}

export function EventCard({ event, onToggleSubscription }: EventCardProps) {
  const safeVideoUrl = sanitizeVideoUrl(event.videoUrl);
  const buttonRef = useRef<HTMLButtonElement>(null);

  const handleSubscribe = () => {
    if (!event.isSubscribed && buttonRef.current) {
      const rect = buttonRef.current.getBoundingClientRect();
      const x = (rect.left + rect.width / 2) / window.innerWidth;
      const y = (rect.top + rect.height / 2) / window.innerHeight;

      confetti({
        particleCount: 100,
        spread: 70,
        origin: { x, y },
        colors: ['#FFD700', '#FFA500', '#FF69B4', '#00CED1', '#9370DB'],
        ticks: 200,
      });

      confetti({
        particleCount: 50,
        angle: 60,
        spread: 55,
        origin: { x: x - 0.1, y },
        colors: ['#FFD700', '#FFA500', '#FF69B4'],
      });

      confetti({
        particleCount: 50,
        angle: 120,
        spread: 55,
        origin: { x: x + 0.1, y },
        colors: ['#00CED1', '#9370DB', '#FF69B4'],
      });
    }
    onToggleSubscription(event.id);
  };
  
  return (
    <motion.article
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      whileHover={{ y: -5 }}
      transition={{ duration: 0.3 }}
      className="relative bg-card backdrop-blur-sm rounded-xl overflow-hidden border border-gray-800/30 shadow-lg"
      aria-label={`Event: ${event.title}`}
    >
      {/* Video/Image Background */}
      <div className="relative h-64 overflow-hidden">
        {safeVideoUrl ? (
          <>
            <iframe
              src={safeVideoUrl}
              className="absolute inset-0 w-full h-full object-cover pointer-events-none"
              allow="autoplay; fullscreen"
              title={`${event.title} video`}
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black via-black/40 to-transparent" />
          </>
        ) : (
          <>
            <div
              className="absolute inset-0 bg-cover bg-center blur-sm scale-110 opacity-30"
              style={{ backgroundImage: `url(${event.imageUrl})` }}
              role="img"
              aria-label={event.title}
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black via-black/40 to-transparent" />
          </>
        )}

        {/* Event Date/Time Badge */}
        <div className="absolute top-4 left-4 z-20">
          <Badge variant="primary">
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
        <p className="text-muted-foreground text-sm leading-relaxed">{event.description}</p>

        {/* DJ Lineup */}
        <section className="space-y-3" aria-labelledby="lineup-heading">
          <h4 id="lineup-heading" className="text-sm font-medium text-muted-foreground">
            Lineup
          </h4>
          <ul className="space-y-2" role="list">
            {event.djLineup.map((dj) => (
              <motion.li
                key={dj.id}
                whileHover={{ x: 3 }}
                className="flex items-center justify-between bg-secondary/50 rounded-lg p-3 border border-gray-800/50"
              >
                <div className="flex items-center gap-3 flex-1">
                  <Avatar
                    src={dj.avatar}
                    alt={dj.name}
                    className="border border-blue-500/20"
                  />
                  <div className="flex-1 min-w-0">
                    <span className="font-medium block truncate text-white">{dj.name}</span>
                    <time className="text-xs text-gray-500">{dj.time}</time>
                  </div>
                </div>

                {/* Social Links */}
                <nav className="flex items-center gap-1.5" aria-label={`${dj.name} social links`}>
                  {dj.social.instagram && (
                    <IconButton
                      href={dj.social.instagram}
                      target="_blank"
                      rel="noopener noreferrer"
                      variant="blue"
                      aria-label={`${dj.name} on Instagram`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </IconButton>
                  )}
                  {dj.social.soundcloud && (
                    <IconButton
                      href={dj.social.soundcloud}
                      target="_blank"
                      rel="noopener noreferrer"
                      variant="orange"
                      aria-label={`${dj.name} on SoundCloud`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </IconButton>
                  )}
                  {dj.social.spotify && (
                    <IconButton
                      href={dj.social.spotify}
                      target="_blank"
                      rel="noopener noreferrer"
                      variant="green"
                      aria-label={`${dj.name} on Spotify`}
                    >
                      <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
                    </IconButton>
                  )}
                </nav>
              </motion.li>
            ))}
          </ul>
        </section>

        {/* Participants and Subscribe Button */}
        <footer className="flex items-center justify-between pt-4 border-t border-gray-800/50">
          <div className="flex items-center gap-2 text-sm text-gray-500">
            <Users className="h-4 w-4" aria-hidden="true" />
            <span>
              <strong className="font-medium text-white">{event.participantCount}</strong> going
            </span>
          </div>

          <Button
            ref={buttonRef}
            onClick={handleSubscribe}
            variant={event.isSubscribed ? 'secondary' : 'primary'}
            aria-pressed={event.isSubscribed}
          >
            {event.isSubscribed ? 'Joined' : 'Join Event'}
          </Button>
        </footer>
      </div>
    </motion.article>
  );
}
