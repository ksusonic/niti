-- Migration: Add performance indexes based on query execution plan analysis
-- This migration adds composite indexes to optimize the complex query involving
-- events, subscriptions, event_djs, and djs tables

-- Composite index for events table to optimize sorting by starts_at DESC with id
-- This helps with the main sort operation and provides covering index benefits
CREATE INDEX idx_events_starts_at_id_desc ON events(starts_at DESC, id);

-- Composite index for subscriptions to optimize user-based event queries
-- This replaces the need for separate lookups and joins
CREATE INDEX idx_subscriptions_user_event_created ON subscriptions(user_id, event_id, created_at);

-- Composite index for event_djs to optimize event-based DJ aggregations
-- This helps with the GroupAggregate operation in the query plan
CREATE INDEX idx_event_djs_event_order_dj ON event_djs(event_id, order_in_lineup, dj_id);

-- Covering index for events to include frequently accessed columns
-- This can eliminate the need to go back to the table for these columns
CREATE INDEX idx_events_covering ON events(id, starts_at DESC)
INCLUDE (title, description, location, cover_url, video_url, created_by);

-- Composite index for better join performance between subscriptions and events
-- This optimizes the hash join operations seen in the execution plan
CREATE INDEX idx_subscriptions_event_user ON subscriptions(event_id, user_id);

-- Index to optimize DJ lookups in event_djs joins
-- This helps with the hash join between event_djs and djs
CREATE INDEX idx_event_djs_dj_event ON event_djs(dj_id, event_id);

-- Partial index for active/future events only
-- This can significantly reduce index size and improve performance for current events
CREATE INDEX idx_events_future_starts_at ON events(starts_at DESC)
WHERE starts_at > CURRENT_TIMESTAMP;

-- Composite index for djs to optimize stage_name searches with id
-- This helps when joining djs table and sorting by stage_name
CREATE INDEX idx_djs_stage_name_id ON djs(stage_name, id);

-- Index to optimize user subscription counts and aggregations
CREATE INDEX idx_subscriptions_user_created ON subscriptions(user_id, created_at DESC);

-- Partial index for events with location (many queries filter by location)
CREATE INDEX idx_events_location_starts_at ON events(location, starts_at DESC)
WHERE location IS NOT NULL;
