-- =========================================
-- SEED DATA
-- =========================================

-- Insert mock profiles
insert into public.profiles (username, display_name, role, bio, avatar_url, social_links)
values
  ('ravekid2024', 'RaveKid2024', 'dj', 'Electronic music enthusiast and weekend DJ. Love deep house and techno vibes.', 'https://example.com/avatar1.png', '{"instagram": "https://instagram.com/ravekid2024"}'),
  ('djnexus', 'DJ Nexus', 'dj', 'Deep house DJ based in downtown.', 'https://example.com/avatar2.png', '{}'),
  ('lunaeclipse', 'Luna Eclipse', 'dj', 'Techno and trance specialist.', 'https://example.com/avatar3.png', '{}'),
  ('partyfan', 'PartyFan', 'fan', 'Just here for the beats.', 'https://example.com/avatar4.png', '{}');

-- Insert events
insert into public.events (title, description, location, start_time, end_time, banner_url, created_by)
values
  ('Underground Beats', 'A night of deep house and techno in the heart of the city.', 'The Basement Club, Downtown', '2025-12-15T22:00:00Z', '2025-12-16T01:00:00Z', 'https://example.com/event1.png', 1),
  ('Neon Nights Festival', 'Massive electronic festival featuring top DJs.', 'Metro Convention Center', '2025-12-22T18:00:00Z', '2025-12-23T02:00:00Z', 'https://example.com/event2.png', 1);

-- Lineup for Underground Beats
insert into public.event_lineup (event_id, dj_id, start_time, end_time, "order")
values
  (1, 2, '2025-12-15T22:00:00Z', '2025-12-15T23:30:00Z', 1),
  (1, 3, '2025-12-15T23:30:00Z', '2025-12-16T01:00:00Z', 2);

-- Participants
insert into public.event_participants (event_id, user_id, status)
values
  (1, 4, 'going');

-- DJ sets
insert into public.dj_sets (dj_id, title, venue, date)
values
  (1, 'Late Night Sessions', 'Club Voltage', '2026-01-12T22:00:00Z'),
  (1, 'Weekend Warrior', 'The Underground', '2026-01-19T22:00:00Z');

-- User settings
insert into public.user_settings (user_id, push_notifications)
values
  (4, true);
