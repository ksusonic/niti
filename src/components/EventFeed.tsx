import { EventCard } from './EventCard';
import type { Event } from '@/types/events';

interface EventFeedProps {
  events: Event[];
  onToggleSubscription: (eventId: string) => void;
}

export function EventFeed({ events, onToggleSubscription }: EventFeedProps) {
  return (
    <main className="min-h-screen bg-background">
      {/* Header */}
      <header className="sticky top-0 z-50 bg-black/95 backdrop-blur-lg border-b border-gray-800/50 px-4 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div>
              <h1 className="text-2xl font-bold text-foreground">Events</h1>
              <p className="text-sm text-muted-foreground">Discover amazing DJ events</p>
            </div>
          </div>
        </div>
      </header>

      {/* Events List */}
      <div className="p-4 space-y-6" role="feed">
        {events.map((event) => (
          <div key={event.id}>
            <EventCard event={event} onToggleSubscription={onToggleSubscription} />
          </div>
        ))}
      </div>
    </main>
  );
}
