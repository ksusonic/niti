'use client';

import { useState, useCallback } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { EventFeed } from '@/components/EventFeed';
import { ProfilePage } from '@/components/ProfilePage';
import { BottomNavigation } from '@/components/BottomNavigation';
import type { Event, UserProfile } from '@/types/events';

// Mock data
const mockEvents: Event[] = [
  {
    id: '1',
    title: 'Underground Beats',
    description: 'A night of deep house and techno in the heart of the city. Experience the underground scene with top DJs spinning the latest tracks.',
    location: 'The Basement Club, Downtown',
    videoUrl: 'https://player.vimeo.com/video/76979871?background=1&autoplay=1&loop=1&byline=0&title=0',
    imageUrl: 'https://images.unsplash.com/photo-1643236990197-9c95f22d3c20?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxESiUyMG1peGluZyUyMG11c2ljfGVufDF8fHx8MTc1ODIyNTQ5NXww&ixlib=rb-4.1.0&q=80&w=1080',
    djLineup: [
      {
        id: 'dj1',
        name: 'DJ Nexus',
        avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&h=150&fit=crop&crop=face',
        time: '10:00 PM - 11:30 PM',
        social: {
          instagram: 'https://instagram.com/djnexus',
          soundcloud: 'https://soundcloud.com/djnexus'
        }
      },
      {
        id: 'dj2',
        name: 'Luna Eclipse',
        avatar: 'https://images.unsplash.com/photo-1494790108755-2616b5e000ec?w=150&h=150&fit=crop&crop=face',
        time: '11:30 PM - 1:00 AM',
        social: {
          instagram: 'https://instagram.com/lunaeclipse',
          spotify: 'https://spotify.com/lunaeclipse'
        }
      }
    ],
    participantCount: 127,
    isSubscribed: false,
    date: 'Dec 15',
    time: '10:00 PM'
  },
  {
    id: '2',
    title: 'Neon Nights Festival',
    description: 'Multi-stage electronic music festival featuring the biggest names in EDM, progressive house, and trance.',
    location: 'Metro Convention Center',
    imageUrl: 'https://images.unsplash.com/photo-1630497326964-62cd41a012d7?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVjdHJvbmljJTIwbXVzaWMlMjBmZXN0aXZhbHxlbnwxfHx8fDE3NTgxNTA5NTZ8MA&ixlib=rb-4.1.0&q=80&w=1080',
    djLineup: [
      {
        id: 'dj3',
        name: 'Electric Storm',
        avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=150&h=150&fit=crop&crop=face',
        time: '6:00 PM - 7:30 PM',
        social: {
          instagram: 'https://instagram.com/electricstorm',
          soundcloud: 'https://soundcloud.com/electricstorm'
        }
      },
      {
        id: 'dj4',
        name: 'Synthia Wave',
        avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=150&h=150&fit=crop&crop=face',
        time: '7:30 PM - 9:00 PM',
        social: {
          instagram: 'https://instagram.com/synthiawave',
          spotify: 'https://spotify.com/synthiawave'
        }
      },
      {
        id: 'dj5',
        name: 'Bass Prophet',
        avatar: 'https://images.unsplash.com/photo-1599566150163-29194dcaad36?w=150&h=150&fit=crop&crop=face',
        time: '9:00 PM - 11:00 PM',
        social: {
          soundcloud: 'https://soundcloud.com/bassprophet'
        }
      }
    ],
    participantCount: 2341,
    isSubscribed: true,
    date: 'Dec 22',
    time: '6:00 PM'
  },
  {
    id: '3',
    title: 'Warehouse Sessions',
    description: 'Raw and unfiltered techno in an authentic warehouse setting. Pure underground vibes.',
    location: 'Industrial District Warehouse',
    imageUrl: 'https://images.unsplash.com/photo-1709239511642-1677410c306a?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxjbHViJTIwcGFydHklMjBsaWdodHN8ZW58MXx8fHwxNzU4MjI1NTA1fDA&ixlib=rb-4.1.0&q=80&w=1080',
    djLineup: [
      {
        id: 'dj6',
        name: 'Concrete Jungle',
        avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=150&h=150&fit=crop&crop=face',
        time: '11:00 PM - 2:00 AM',
        social: {
          instagram: 'https://instagram.com/concretejungle',
          soundcloud: 'https://soundcloud.com/concretejungle'
        }
      }
    ],
    participantCount: 89,
    isSubscribed: false,
    date: 'Dec 28',
    time: '11:00 PM'
  },
  {
    id: '4',
    title: 'Rave Revolution',
    description: 'Old school meets new school in this epic rave experience. Featuring classic breaks, drum & bass, and hardcore.',
    location: 'Underground Tunnel System',
    videoUrl: 'https://player.vimeo.com/video/271392047?background=1&autoplay=1&loop=1&byline=0&title=0',
    imageUrl: 'https://images.unsplash.com/photo-1465917031443-a76ab279572f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxyYXZlJTIwdW5kZXJncm91bmR8ZW58MXx8fHwxNzU4MjI1NTEzfDA&ixlib=rb-4.1.0&q=80&w=1080',
    djLineup: [
      {
        id: 'dj7',
        name: 'Quantum Beats',
        avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=150&h=150&fit=crop&crop=face',
        time: '9:00 PM - 10:30 PM',
        social: {
          instagram: 'https://instagram.com/quantumbeats',
          soundcloud: 'https://soundcloud.com/quantumbeats',
          spotify: 'https://spotify.com/quantumbeats'
        }
      },
      {
        id: 'dj8',
        name: 'Hardcore Hero',
        avatar: 'https://images.unsplash.com/photo-1463453091185-61582044d556?w=150&h=150&fit=crop&crop=face',
        time: '10:30 PM - 12:00 AM',
        social: {
          soundcloud: 'https://soundcloud.com/hardcorehero'
        }
      }
    ],
    participantCount: 456,
    isSubscribed: true,
    date: 'Jan 5',
    time: '9:00 PM'
  }
];

const mockProfile: UserProfile = {
  username: 'RaveKid2024',
  avatar: 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=150&h=150&fit=crop&crop=face',
  isDJ: true,
  bio: 'Electronic music enthusiast and weekend DJ. Love deep house and techno vibes.',
  socialLinks: {
    instagram: 'https://instagram.com/ravekid2024',
    soundcloud: 'https://soundcloud.com/ravekid2024',
    spotify: 'https://spotify.com/ravekid2024'
  },
  upcomingSets: [
    {
      id: 'set1',
      event: 'Late Night Sessions',
      date: 'Jan 12',
      venue: 'Club Voltage'
    },
    {
      id: 'set2',
      event: 'Weekend Warrior',
      date: 'Jan 19',
      venue: 'The Underground'
    }
  ],
  subscribedEvents: [
    {
      id: '2',
      title: 'Neon Nights Festival',
      date: 'Dec 22',
      location: 'Metro Convention Center',
      imageUrl: 'https://images.unsplash.com/photo-1630497326964-62cd41a012d7?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVjdHJvbmljJTIwbXVzaWMlMjBmZXN0aXZhbHxlbnwxfHx8fDE3NTgxNTA5NTZ8MA&ixlib=rb-4.1.0&q=80&w=1080'
    },
    {
      id: '4',
      title: 'Rave Revolution',
      date: 'Jan 5',
      location: 'Underground Tunnel System',
      imageUrl: 'https://images.unsplash.com/photo-1465917031443-a76ab279572f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxyYXZlJTIwdW5kZXJncm91bmR8ZW58MXx8fHwxNzU4MjI1NTEzfDA&ixlib=rb-4.1.0&q=80&w=1080'
    }
  ],
  settings: {
    notifications: true,
    preferredVenues: ['The Basement Club', 'Metro Convention Center']
  }
};

export default function Home() {
  const [activeTab, setActiveTab] = useState<'events' | 'profile'>('events');
  const [events, setEvents] = useState(mockEvents);
  const [profile, setProfile] = useState(mockProfile);

  const handleToggleSubscription = (eventId: string) => {
    setEvents(prev => prev.map(event => 
      event.id === eventId 
        ? { 
            ...event, 
            isSubscribed: !event.isSubscribed,
            participantCount: event.isSubscribed 
              ? event.participantCount - 1 
              : event.participantCount + 1
          }
        : event
    ));

    // Update profile subscribed events
    const event = events.find(e => e.id === eventId);
    if (event) {
      if (event.isSubscribed) {
        // Remove from subscribed events
        setProfile(prev => ({
          ...prev,
          subscribedEvents: prev.subscribedEvents.filter(e => e.id !== eventId)
        }));
      } else {
        // Add to subscribed events
        setProfile(prev => ({
          ...prev,
          subscribedEvents: [...prev.subscribedEvents, {
            id: event.id,
            title: event.title,
            date: event.date,
            location: event.location,
            imageUrl: event.imageUrl
          }]
        }));
      }
    }
  };

  const handleUpdateProfile = (updates: Partial<typeof profile>) => {
    setProfile(prev => ({ ...prev, ...updates }));
  };

  const handleTabChange = useCallback((tab: 'events' | 'profile') => {
    setActiveTab(tab);
  }, []);

  return (
    <div className="bg-background text-foreground dark pb-24">
      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          className="min-h-screen page-transition"
          animate={{ opacity: 1, x: 0 }}
          exit={{ opacity: 0, x: -40 }}
          transition={{ duration: 0.3 }}
        >
          {activeTab === 'events' ? (
            <EventFeed 
              events={events} 
              onToggleSubscription={handleToggleSubscription} 
            />
          ) : (
            <ProfilePage 
              profile={profile} 
              onUpdateProfile={handleUpdateProfile} 
            />
          )}
        </motion.div>
      </AnimatePresence>

      <div className="fixed bottom-0 left-0 right-0 z-50">
        <BottomNavigation 
          activeTab={activeTab} 
          onTabChange={handleTabChange} 
        />
      </div>
    </div>
  );
}
