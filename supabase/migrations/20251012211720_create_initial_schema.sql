-- =========================================
-- PROFILES
-- =========================================
create table if not exists public.profiles (
  id bigserial primary key,
  username text unique not null,
  display_name text,
  avatar_url text,
  bio text,
  role text check (role in ('dj', 'fan', 'organizer')) default 'fan',
  social_links jsonb default '{}'::jsonb,
  created_at timestamptz default now()
);

-- =========================================
-- EVENTS
-- =========================================
create table if not exists public.events (
  id bigserial primary key,
  title text not null,
  description text,
  location text,
  start_time timestamptz not null,
  end_time timestamptz,
  banner_url text,
  created_by bigint references public.profiles(id) on delete set null,
  created_at timestamptz default now()
);

-- =========================================
-- EVENT LINEUP
-- =========================================
create table if not exists public.event_lineup (
  id bigserial primary key,
  event_id bigint references public.events(id) on delete cascade,
  dj_id bigint references public.profiles(id) on delete cascade,
  start_time timestamptz not null,
  end_time timestamptz,
  "order" int,
  created_at timestamptz default now()
);

-- =========================================
-- EVENT PARTICIPANTS
-- =========================================
create table if not exists public.event_participants (
  id bigserial primary key,
  event_id bigint references public.events(id) on delete cascade,
  user_id bigint references public.profiles(id) on delete cascade,
  status text check (status in ('going', 'interested', 'not_going')) default 'going',
  joined_at timestamptz default now(),
  unique (event_id, user_id)
);

-- =========================================
-- USER SETTINGS
-- =========================================
create table if not exists public.user_settings (
  user_id bigint primary key references public.profiles(id) on delete cascade,
  push_notifications boolean default true,
  telegram_chat_id text,
  updated_at timestamptz default now()
);

-- =========================================
-- DJ SETS (optional table)
-- =========================================
create table if not exists public.dj_sets (
  id bigserial primary key,
  dj_id bigint references public.profiles(id) on delete cascade,
  title text not null,
  event_id bigint references public.events(id) on delete set null,
  venue text,
  date timestamptz not null,
  created_at timestamptz default now()
);

-- =========================================
-- INDEXES
-- =========================================
create index if not exists idx_events_start_time on public.events (start_time);
create index if not exists idx_event_lineup_event_id on public.event_lineup (event_id);
create index if not exists idx_event_participants_user_id on public.event_participants (user_id);
create index if not exists idx_dj_sets_dj_id on public.dj_sets (dj_id);
