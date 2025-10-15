-- =========================================
-- SEED DATA
-- =========================================

-- Insert mock profiles
insert into public.profiles (username, display_name, role, bio, avatar_url, social_links)
values
  -- Main user profile
  ('ravekid2024', 'RaveKid2024', 'dj', 'Electronic music enthusiast and weekend DJ. Love deep house and techno vibes.', 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/ravekid2024", "soundcloud": "https://soundcloud.com/ravekid2024", "spotify": "https://spotify.com/ravekid2024"}'),
  
  -- DJs for Event 1: Underground Beats
  ('djnexus', 'DJ Nexus', 'dj', 'Deep house DJ based in downtown.', 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/djnexus", "soundcloud": "https://soundcloud.com/djnexus"}'),
  ('lunaeclipse', 'Luna Eclipse', 'dj', 'Techno and trance specialist.', 'https://images.unsplash.com/photo-1599423424751-54e0c1187a02?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/lunaeclipse", "spotify": "https://spotify.com/lunaeclipse"}'),
  
  -- DJs for Event 2: Neon Nights Festival
  ('electricstorm', 'Electric Storm', 'dj', 'EDM powerhouse.', 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/electricstorm", "soundcloud": "https://soundcloud.com/electricstorm"}'),
  ('synthiawave', 'Synthia Wave', 'dj', 'Progressive house specialist.', 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/synthiawave", "spotify": "https://spotify.com/synthiawave"}'),
  ('bassprophet', 'Bass Prophet', 'dj', 'Bass music expert.', 'https://images.unsplash.com/photo-1599566150163-29194dcaad36?w=150&h=150&fit=crop&crop=face', '{"soundcloud": "https://soundcloud.com/bassprophet"}'),
  
  -- DJs for Event 3: Warehouse Sessions
  ('concretejungle', 'Concrete Jungle', 'dj', 'Raw techno specialist.', 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/concretejungle", "soundcloud": "https://soundcloud.com/concretejungle"}'),
  
  -- DJs for Event 4: Rave Revolution
  ('quantumbeats', 'Quantum Beats', 'dj', 'Drum & bass pioneer.', 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/quantumbeats", "soundcloud": "https://soundcloud.com/quantumbeats", "spotify": "https://spotify.com/quantumbeats"}'),
  ('hardcorehero', 'Hardcore Hero', 'dj', 'Hardcore music legend.', 'https://images.unsplash.com/photo-1463453091185-61582044d556?w=150&h=150&fit=crop&crop=face', '{"soundcloud": "https://soundcloud.com/hardcorehero"}');

-- Insert events
insert into public.events (title, description, location, start_time, end_time, banner_url, created_by)
values
  (
    'Underground Beats',
    'A night of deep house and techno in the heart of the city. Experience the underground scene with top DJs spinning the latest tracks.',
    'The Basement Club, Downtown',
    '2025-12-15T22:00:00Z',
    '2025-12-16T01:00:00Z',
    'https://images.unsplash.com/photo-1643236990197-9c95f22d3c20?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxESiUyMG1peGluZyUyMG11c2ljfGVufDF8fHx8MTc1ODIyNTQ5NXww&ixlib=rb-4.1.0&q=80&w=1080',
    1
  ),
  (
    'Neon Nights Festival',
    'Multi-stage electronic music festival featuring the biggest names in EDM, progressive house, and trance.',
    'Metro Convention Center',
    '2025-12-22T18:00:00Z',
    '2025-12-23T02:00:00Z',
    'https://images.unsplash.com/photo-1630497326964-62cd41a012d7?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVjdHJvbmljJTIwbXVzaWMlMjBmZXN0aXZhbHxlbnwxfHx8fDE3NTgxNTA5NTZ8MA&ixlib=rb-4.1.0&q=80&w=1080',
    1
  ),
  (
    'Warehouse Sessions',
    'Raw and unfiltered techno in an authentic warehouse setting. Pure underground vibes.',
    'Industrial District Warehouse',
    '2025-12-28T23:00:00Z',
    '2025-12-29T02:00:00Z',
    'https://images.unsplash.com/photo-1709239511642-1677410c306a?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxjbHViJTIwcGFydHklMjBsaWdodHN8ZW58MXx8fHwxNzU4MjI1NTA1fDA&ixlib=rb-4.1.0&q=80&w=1080',
    1
  ),
  (
    'Rave Revolution',
    'Old school meets new school in this epic rave experience. Featuring classic breaks, drum & bass, and hardcore.',
    'Underground Tunnel System',
    '2026-01-05T21:00:00Z',
    '2026-01-06T00:00:00Z',
    'https://images.unsplash.com/photo-1465917031443-a76ab279572f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxyYXZlJTIwdW5kZXJncm91bmR8ZW58MXx8fHwxNzU4MjI1NTEzfDA&ixlib=rb-4.1.0&q=80&w=1080',
    1
  );

-- Event lineup
insert into public.event_lineup (event_id, dj_id, start_time, end_time, "order")
values
  -- Event 1: Underground Beats
  (1, 2, '2025-12-15T22:00:00Z', '2025-12-15T23:30:00Z', 1),
  (1, 3, '2025-12-15T23:30:00Z', '2025-12-16T01:00:00Z', 2),
  
  -- Event 2: Neon Nights Festival
  (2, 4, '2025-12-22T18:00:00Z', '2025-12-22T19:30:00Z', 1),
  (2, 5, '2025-12-22T19:30:00Z', '2025-12-22T21:00:00Z', 2),
  (2, 6, '2025-12-22T21:00:00Z', '2025-12-22T23:00:00Z', 3),
  
  -- Event 3: Warehouse Sessions
  (3, 7, '2025-12-28T23:00:00Z', '2025-12-29T02:00:00Z', 1),
  
  -- Event 4: Rave Revolution
  (4, 8, '2026-01-05T21:00:00Z', '2026-01-05T22:30:00Z', 1),
  (4, 9, '2026-01-05T22:30:00Z', '2026-01-06T00:00:00Z', 2);

-- Event participants
-- The participant counts will be tracked in a separate way since we don't want to create thousands of fake users
-- Instead, we'll just subscribe the main user to events 2 and 4

-- Event 2: Neon Nights Festival (main user subscribed)
insert into public.event_participants (event_id, user_id, status)
values (2, 1, 'going');

-- Event 4: Rave Revolution (main user subscribed)
insert into public.event_participants (event_id, user_id, status)
values (4, 1, 'going');

-- Note: Participant counts (127, 2341, 89, 456) can be stored as metadata 
-- or calculated from actual participants when real users start joining events

-- DJ sets for main user profile
insert into public.dj_sets (dj_id, title, venue, date)
values
  (1, 'Late Night Sessions', 'Club Voltage', '2026-01-12T22:00:00Z'),
  (1, 'Weekend Warrior', 'The Underground', '2026-01-19T22:00:00Z');

-- User settings
insert into public.user_settings (user_id, push_notifications)
values
  (1, true);
