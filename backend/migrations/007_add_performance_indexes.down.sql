-- Migration down: Remove performance indexes added in 007_add_performance_indexes.up.sql
-- This migration removes all the composite and covering indexes added for query optimization

-- Remove partial index for events with location
DROP INDEX IF EXISTS idx_events_location_starts_at;

-- Remove index for user subscription counts and aggregations
DROP INDEX IF EXISTS idx_subscriptions_user_created;

-- Remove composite index for djs stage_name with id
DROP INDEX IF EXISTS idx_djs_stage_name_id;

-- Remove partial index for active/future events
DROP INDEX IF EXISTS idx_events_future_starts_at;

-- Remove index for DJ lookups in event_djs joins
DROP INDEX IF EXISTS idx_event_djs_dj_event;

-- Remove composite index for subscriptions event-user optimization
DROP INDEX IF EXISTS idx_subscriptions_event_user;

-- Remove covering index for events
DROP INDEX IF EXISTS idx_events_covering;

-- Remove composite index for event_djs optimization
DROP INDEX IF EXISTS idx_event_djs_event_order_dj;

-- Remove composite index for subscriptions user-event optimization
DROP INDEX IF EXISTS idx_subscriptions_user_event_created;

-- Remove composite index for events starts_at with id
DROP INDEX IF EXISTS idx_events_starts_at_id_desc;
