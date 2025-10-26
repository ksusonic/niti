-- =========================================
-- SEED DATA
-- =========================================

-- Insert mock profiles and capture their IDs
with inserted_profiles as (
  insert into public.profiles (username, display_name, role, bio, avatar_url, social_links)
  values
    -- Main user profile
    ('ravekid2024', 'RaveKid2024', 'dj', 'Любитель электронной музыки и диджей по выходным. Обожаю deep house и techno.', 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/ravekid2024", "soundcloud": "https://soundcloud.com/ravekid2024", "spotify": "https://spotify.com/ravekid2024"}'),
    
    -- DJs for Event 1: Underground Beats
    ('djnexus', 'DJ Nexus', 'dj', 'Deep house диджей из центра города.', 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/djnexus", "soundcloud": "https://soundcloud.com/djnexus"}'),
    ('lunaeclipse', 'Luna Eclipse', 'dj', 'Специалист по techno и trance.', 'https://images.unsplash.com/photo-1599423424751-54e0c1187a02?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/lunaeclipse", "spotify": "https://spotify.com/lunaeclipse"}'),
    
    -- DJs for Event 2: Neon Nights Festival
    ('electricstorm', 'Electric Storm', 'dj', 'EDM мастер.', 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/electricstorm", "soundcloud": "https://soundcloud.com/electricstorm"}'),
    ('synthiawave', 'Synthia Wave', 'dj', 'Специалист по progressive house.', 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/synthiawave", "spotify": "https://spotify.com/synthiawave"}'),
    ('bassprophet', 'Bass Prophet', 'dj', 'Эксперт по басовой музыке.', 'https://images.unsplash.com/photo-1599566150163-29194dcaad36?w=150&h=150&fit=crop&crop=face', '{"soundcloud": "https://soundcloud.com/bassprophet"}'),
    
    -- DJs for Event 3: Warehouse Sessions
    ('concretejungle', 'Concrete Jungle', 'dj', 'Специалист по жёсткому techno.', 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/concretejungle", "soundcloud": "https://soundcloud.com/concretejungle"}'),
    
    -- DJs for Event 4: Rave Revolution
    ('quantumbeats', 'Quantum Beats', 'dj', 'Пионер drum & bass.', 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=150&h=150&fit=crop&crop=face', '{"instagram": "https://instagram.com/quantumbeats", "soundcloud": "https://soundcloud.com/quantumbeats", "spotify": "https://spotify.com/quantumbeats"}'),
    ('hardcorehero', 'Hardcore Hero', 'dj', 'Легенда hardcore музыки.', 'https://images.unsplash.com/photo-1463453091185-61582044d556?w=150&h=150&fit=crop&crop=face', '{"soundcloud": "https://soundcloud.com/hardcorehero"}')
  returning id, username
),
profile_ids as (
  select 
    (select id from inserted_profiles where username = 'ravekid2024') as main_user_id,
    (select id from inserted_profiles where username = 'djnexus') as djnexus_id,
    (select id from inserted_profiles where username = 'lunaeclipse') as lunaeclipse_id,
    (select id from inserted_profiles where username = 'electricstorm') as electricstorm_id,
    (select id from inserted_profiles where username = 'synthiawave') as synthiawave_id,
    (select id from inserted_profiles where username = 'bassprophet') as bassprophet_id,
    (select id from inserted_profiles where username = 'concretejungle') as concretejungle_id,
    (select id from inserted_profiles where username = 'quantumbeats') as quantumbeats_id,
    (select id from inserted_profiles where username = 'hardcorehero') as hardcorehero_id
),
-- Insert events
inserted_events as (
  insert into public.events (title, description, location, start_time, end_time, banner_url, created_by)
  select
    title, description, location, start_time, end_time, banner_url, (select main_user_id from profile_ids)
  from (
    values
      (
        'Underground Beats',
        'Ночь deep house и techno в сердце города. Погрузитесь в андеграундную сцену с топовыми диджеями, которые крутят свежайшие треки.',
        'The Basement Club, центр',
        '2025-12-15T22:00:00Z'::timestamptz,
        '2025-12-16T01:00:00Z'::timestamptz,
        'https://images.unsplash.com/photo-1643236990197-9c95f22d3c20?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxESiUyMG1peGluZyUyMG11c2ljfGVufDF8fHx8MTc1ODIyNTQ5NXww&ixlib=rb-4.1.0&q=80&w=1080'
      ),
      (
        'Neon Nights Festival',
        'Многоступенчатый фестиваль электронной музыки с крупнейшими именами в EDM, progressive house и trance.',
        'Конгресс-центр Метро',
        '2025-12-22T18:00:00Z'::timestamptz,
        '2025-12-23T02:00:00Z'::timestamptz,
        'https://images.unsplash.com/photo-1630497326964-62cd41a012d7?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVjdHJvbmljJTIwbXVzaWMlMjBmZXN0aXZhbHxlbnwxfHx8fDE3NTgxNTA5NTZ8MA&ixlib=rb-4.1.0&q=80&w=1080'
      ),
      (
        'Warehouse Sessions',
        'Сырой и неотфильтрованный techno в настоящем складском пространстве. Чистый андеграундный вайб.',
        'Склад в промышленном районе',
        '2025-12-28T23:00:00Z'::timestamptz,
        '2025-12-29T02:00:00Z'::timestamptz,
        'https://images.unsplash.com/photo-1709239511642-1677410c306a?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxjbHViJTIwcGFydHklMjBsaWdodHN8ZW58MXx8fHwxNzU4MjI1NTA1fDA&ixlib=rb-4.1.0&q=80&w=1080'
      ),
      (
        'Rave Revolution',
        'Старая школа встречается с новой школой в этом эпическом рейв-событии. Классические брейки, drum & bass и hardcore.',
        'Подземная туннельная система',
        '2026-01-05T21:00:00Z'::timestamptz,
        '2026-01-06T00:00:00Z'::timestamptz,
        'https://images.unsplash.com/photo-1465917031443-a76ab279572f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxyYXZlJTIwdW5kZXJncm91bmR8ZW58MXx8fHwxNzU4MjI1NTEzfDA&ixlib=rb-4.1.0&q=80&w=1080'
      )
  ) as t(title, description, location, start_time, end_time, banner_url)
  returning id, title
),
event_ids as (
  select 
    (select id from inserted_events where title = 'Underground Beats') as event1_id,
    (select id from inserted_events where title = 'Neon Nights Festival') as event2_id,
    (select id from inserted_events where title = 'Warehouse Sessions') as event3_id,
    (select id from inserted_events where title = 'Rave Revolution') as event4_id
),
-- Insert event lineup
inserted_lineup as (
  insert into public.event_lineup (event_id, dj_id, start_time, end_time, "order")
  select event_id, dj_id, start_time, end_time, "order"
  from (
    select (select event1_id from event_ids) as event_id, (select djnexus_id from profile_ids) as dj_id, '2025-12-15T22:00:00Z'::timestamptz as start_time, '2025-12-15T23:30:00Z'::timestamptz as end_time, 1 as "order"
    union all
    select (select event1_id from event_ids), (select lunaeclipse_id from profile_ids), '2025-12-15T23:30:00Z'::timestamptz, '2025-12-16T01:00:00Z'::timestamptz, 2
    union all
    select (select event2_id from event_ids), (select electricstorm_id from profile_ids), '2025-12-22T18:00:00Z'::timestamptz, '2025-12-22T19:30:00Z'::timestamptz, 1
    union all
    select (select event2_id from event_ids), (select synthiawave_id from profile_ids), '2025-12-22T19:30:00Z'::timestamptz, '2025-12-22T21:00:00Z'::timestamptz, 2
    union all
    select (select event2_id from event_ids), (select bassprophet_id from profile_ids), '2025-12-22T21:00:00Z'::timestamptz, '2025-12-22T23:00:00Z'::timestamptz, 3
    union all
    select (select event3_id from event_ids), (select concretejungle_id from profile_ids), '2025-12-28T23:00:00Z'::timestamptz, '2025-12-29T02:00:00Z'::timestamptz, 1
    union all
    select (select event4_id from event_ids), (select quantumbeats_id from profile_ids), '2026-01-05T21:00:00Z'::timestamptz, '2026-01-05T22:30:00Z'::timestamptz, 1
    union all
    select (select event4_id from event_ids), (select hardcorehero_id from profile_ids), '2026-01-05T22:30:00Z'::timestamptz, '2026-01-06T00:00:00Z'::timestamptz, 2
  ) lineup
  returning id
),
-- Insert event participants
inserted_participants as (
  insert into public.event_participants (event_id, user_id, status)
  select event_id, user_id, status
  from (
    select (select event2_id from event_ids) as event_id, (select main_user_id from profile_ids) as user_id, 'going' as status
    union all
    select (select event4_id from event_ids), (select main_user_id from profile_ids), 'going'
  ) participants
  returning id
),
-- Insert DJ sets
inserted_dj_sets as (
  insert into public.dj_sets (dj_id, title, venue, date)
  select dj_id, title, venue, date
  from (
    select (select main_user_id from profile_ids) as dj_id, 'Поздние ночные сессии' as title, 'Club Voltage' as venue, '2026-01-12T22:00:00Z'::timestamptz as date
    union all
    select (select main_user_id from profile_ids), 'Воин выходного дня', 'The Underground', '2026-01-19T22:00:00Z'::timestamptz
  ) sets
  returning id
)
-- Insert user settings
insert into public.user_settings (user_id, push_notifications)
select (select main_user_id from profile_ids), true;
