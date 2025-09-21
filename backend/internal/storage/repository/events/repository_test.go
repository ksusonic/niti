//go:build integration

package events_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
	eventsrepo "github.com/ksusonic/niti/backend/internal/storage/repository/events"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
	"github.com/stretchr/testify/require"
)

func TestEventsRepository_CRUD(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// First create users that will be referenced by events
		creatorTelegramID := int64(12345)
		viewerTelegramID := int64(67890)

		user := &models.User{
			TelegramID: creatorTelegramID,
			Username:   "testuser",
			FirstName:  "Test",
		}

		_, err := userRepo.Create(ctx, user)
		require.NoError(t, err)

		viewer := &models.User{
			TelegramID: viewerTelegramID,
			Username:   "viewer",
			FirstName:  "Viewer",
		}

		_, err = userRepo.Create(ctx, viewer)
		require.NoError(t, err)

		// Test data
		location := "Test Venue"
		videoURL := "https://example.com/stream"
		now := time.Now()

		event := &models.Event{
			Title:       "Test Event",
			Description: "This is a test event",
			Location:    &location,
			VideoURL:    &videoURL,
			StartsAt:    &now,
			CreatedBy:   &creatorTelegramID,
		}

		// CREATE
		eventID, err := repo.CreateEvent(ctx, event)
		require.NoError(t, err)
		require.NotZero(t, eventID)

		// GET BY ID
		retrievedEvent, err := repo.GetEventByID(ctx, eventID)
		require.NoError(t, err)
		require.NotNil(t, retrievedEvent)
		require.Equal(t, eventID, retrievedEvent.ID)
		require.Equal(t, "Test Event", retrievedEvent.Title)
		require.Equal(t, "This is a test event", retrievedEvent.Description)
		require.Equal(t, location, *retrievedEvent.Location)
		require.Equal(t, videoURL, *retrievedEvent.VideoURL)
		require.Equal(t, creatorTelegramID, *retrievedEvent.CreatedBy)

		// LIST events to verify creation (using viewer user ID)
		events, err := repo.ListEvents(ctx, viewerTelegramID, 10, 0)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(events), 1) // May have other events in DB

		// Find our event in the list
		var foundEvent *models.EventEnriched
		for _, e := range events {
			if e.ID == eventID {
				foundEvent = &e
				break
			}
		}
		require.NotNil(t, foundEvent, "Created event should be found in list")
		require.Equal(t, "Test Event", foundEvent.Title)
		require.Equal(t, "This is a test event", foundEvent.Description)
		require.Equal(t, location, *foundEvent.Location)
		require.Equal(t, videoURL, *foundEvent.VideoURL)
		require.Equal(t, creatorTelegramID, *foundEvent.CreatedBy)
		require.False(t, foundEvent.IsSubscribed) // viewer is not subscribed
		require.NotNil(t, foundEvent.ParticipantsCount)
		require.Equal(t, 0, *foundEvent.ParticipantsCount) // no participants yet
		require.NotNil(t, foundEvent.DJs)

		// CREATE another event
		startsAt2 := time.Now().Add(24 * time.Hour)
		event2 := &models.Event{
			Title:       "Second Event",
			Description: "Second description",
			Location:    nil,
			VideoURL:    nil,
			StartsAt:    &startsAt2,
			CreatedBy:   &creatorTelegramID,
		}

		eventID2, err := repo.CreateEvent(ctx, event2)
		require.NoError(t, err)
		require.NotZero(t, eventID2)

		// LIST events after creating second one
		eventsAfterSecond, err := repo.ListEvents(ctx, viewerTelegramID, 10, 0)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(eventsAfterSecond), 2) // Should have at least our 2 events

		// DELETE first event
		err = repo.DeleteEvent(ctx, eventID)
		require.NoError(t, err)

		// LIST after delete
		eventsAfterDelete, err := repo.ListEvents(ctx, viewerTelegramID, 10, 0)
		require.NoError(t, err)

		// Verify first event is deleted
		var deletedEventFound bool
		for _, e := range eventsAfterDelete {
			if e.ID == eventID {
				deletedEventFound = true
				break
			}
		}
		require.False(t, deletedEventFound, "Deleted event should not be found in list")

		// Verify second event still exists
		var secondEventFound bool
		for _, e := range eventsAfterDelete {
			if e.ID == eventID2 {
				secondEventFound = true
				require.Equal(t, "Second Event", e.Title)
				break
			}
		}
		require.True(t, secondEventFound, "Second event should still exist")
	})
	require.NoError(t, err)
}

func TestEventsRepository_CreateEvent(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Create user first
		creatorID := int64(123)
		user := &models.User{
			TelegramID: creatorID,
			Username:   "creator",
			FirstName:  "Creator",
		}

		_, err := userRepo.Create(ctx, user)
		require.NoError(t, err)

		now := time.Now()

		// Test minimal event creation
		event := &models.Event{
			Title:     "Minimal Event",
			CreatedBy: &creatorID,
		}

		eventID, err := repo.CreateEvent(ctx, event)
		require.NoError(t, err)
		require.NotZero(t, eventID)

		// Verify the event was created
		retrievedEvent, err := repo.GetEventByID(ctx, eventID)
		require.NoError(t, err)
		require.Equal(t, "Minimal Event", retrievedEvent.Title)
		require.Equal(t, creatorID, *retrievedEvent.CreatedBy)

		// Test full event creation
		location := "Full Event Location"
		videoURL := "https://example.com/full-stream"
		fullEvent := &models.Event{
			Title:       "Full Event",
			Description: "Complete event description",
			Location:    &location,
			VideoURL:    &videoURL,
			StartsAt:    &now,
			CreatedBy:   &creatorID,
		}

		fullEventID, err := repo.CreateEvent(ctx, fullEvent)
		require.NoError(t, err)
		require.NotZero(t, fullEventID)

		// Verify the full event was created correctly
		retrievedFullEvent, err := repo.GetEventByID(ctx, fullEventID)
		require.NoError(t, err)
		require.Equal(t, "Full Event", retrievedFullEvent.Title)
		require.Equal(t, "Complete event description", retrievedFullEvent.Description)
		require.Equal(t, location, *retrievedFullEvent.Location)
		require.Equal(t, videoURL, *retrievedFullEvent.VideoURL)
		require.Equal(t, creatorID, *retrievedFullEvent.CreatedBy)
		require.WithinDuration(t, now, *retrievedFullEvent.StartsAt, time.Second)
	})
	require.NoError(t, err)
}

func TestEventsRepository_GetEventByID_NotFound(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Try to get a non-existent event
		_, err := repo.GetEventByID(ctx, 999999)
		require.Error(t, err)
	})
	require.NoError(t, err)
}

func TestEventsRepository_DeleteEvent_NotFound(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Try to delete a non-existent event
		err := repo.DeleteEvent(ctx, 999999)
		require.NoError(t, err) // DELETE operations typically don't fail if the record doesn't exist
	})
	require.NoError(t, err)
}

func TestEventsRepository_ListEvents(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Create user first
		creatorID := int64(456)
		viewerID := int64(789)

		user := &models.User{
			TelegramID: creatorID,
			Username:   "listuser",
			FirstName:  "List",
		}

		_, err := userRepo.Create(ctx, user)
		require.NoError(t, err)

		viewer := &models.User{
			TelegramID: viewerID,
			Username:   "viewer",
			FirstName:  "Viewer",
		}

		_, err = userRepo.Create(ctx, viewer)
		require.NoError(t, err)

		// Get initial count
		initialEvents, err := repo.ListEvents(ctx, viewerID, 10, 0)
		require.NoError(t, err)
		initialCount := len(initialEvents)

		// Create an event
		event := &models.Event{
			Title:     "List Test Event",
			CreatedBy: &creatorID,
		}

		eventID, err := repo.CreateEvent(ctx, event)
		require.NoError(t, err)

		// List events after creation
		events, err := repo.ListEvents(ctx, viewerID, 10, 0)
		require.NoError(t, err)
		require.Equal(t, initialCount+1, len(events))

		// Verify our event is in the list
		var foundEvent *models.EventEnriched
		for _, e := range events {
			if e.ID == eventID {
				foundEvent = &e
				require.Equal(t, "List Test Event", e.Title)
				require.False(t, e.IsSubscribed) // viewer is not subscribed
				require.NotNil(t, e.ParticipantsCount)
				require.Equal(t, 0, *e.ParticipantsCount)
				break
			}
		}
		require.NotNil(t, foundEvent, "Created event should be in the list")
	})
	require.NoError(t, err)
}

func TestEventsRepository_GetUserEvents(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		userID := int64(12345)
		otherUserID := int64(67890)

		// Create users first
		user1 := &models.User{
			TelegramID: userID,
			Username:   "user1",
			FirstName:  "User",
		}
		user2 := &models.User{
			TelegramID: otherUserID,
			Username:   "user2",
			FirstName:  "Other",
		}

		_, err := userRepo.Create(ctx, user1)
		require.NoError(t, err)
		_, err = userRepo.Create(ctx, user2)
		require.NoError(t, err)

		// Create some events
		now := time.Now()
		event1 := &models.Event{
			Title:       "User Event 1",
			Description: "First user event",
			StartsAt:    &now,
			CreatedBy:   &userID,
		}

		event1ID, err := repo.CreateEvent(ctx, event1)
		require.NoError(t, err)

		future := now.Add(24 * time.Hour)
		event2 := &models.Event{
			Title:       "User Event 2",
			Description: "Second user event",
			StartsAt:    &future,
			CreatedBy:   &userID,
		}

		event2ID, err := repo.CreateEvent(ctx, event2)
		require.NoError(t, err)

		// Create event for other user
		otherEvent := &models.Event{
			Title:       "Other User Event",
			Description: "Event by different user",
			StartsAt:    &now,
			CreatedBy:   &otherUserID,
		}

		otherEventID, err := repo.CreateEvent(ctx, otherEvent)
		require.NoError(t, err)

		// Test GetUserEvents with no subscriptions
		// Note: GetUserEvents returns events the user is subscribed to, not events they created
		userEvents, err := repo.GetUserEvents(ctx, userID)
		require.NoError(t, err)
		require.Empty(t, userEvents) // No subscriptions yet

		// Test GetUserEvents for non-existent user
		nonExistentUserEvents, err := repo.GetUserEvents(ctx, int64(999999))
		require.NoError(t, err)
		require.Empty(t, nonExistentUserEvents)

		// Verify events were created (using ListEvents)
		allEvents, err := repo.ListEvents(ctx, userID, 10, 0)
		require.NoError(t, err)

		// Verify our events exist in the list
		eventIDs := []int{event1ID, event2ID, otherEventID}
		foundIDs := make([]int, 0)
		for _, event := range allEvents {
			for _, targetID := range eventIDs {
				if event.ID == targetID {
					foundIDs = append(foundIDs, event.ID)
					break
				}
			}
		}
		require.Len(t, foundIDs, 3, "All created events should be found")
	})
	require.NoError(t, err)
}

func TestEventsRepository_ListEvents_LimitAndOffset(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Create user first
		creatorID := int64(789)
		viewerID := int64(987)

		user := &models.User{
			TelegramID: creatorID,
			Username:   "paginationuser",
			FirstName:  "Pagination",
		}

		_, err := userRepo.Create(ctx, user)
		require.NoError(t, err)

		viewer := &models.User{
			TelegramID: viewerID,
			Username:   "viewer",
			FirstName:  "Viewer",
		}

		_, err = userRepo.Create(ctx, viewer)
		require.NoError(t, err)

		// Create multiple events with different start times to test ordering
		baseTime := time.Now()
		eventIDs := make([]int, 5)

		for i := 0; i < 5; i++ {
			startsAt := baseTime.Add(time.Duration(i) * time.Hour)
			event := &models.Event{
				Title:       fmt.Sprintf("Event %d", i+1),
				Description: fmt.Sprintf("Description for event %d", i+1),
				StartsAt:    &startsAt,
				CreatedBy:   &creatorID,
			}

			eventID, err := repo.CreateEvent(ctx, event)
			require.NoError(t, err)
			eventIDs[i] = eventID
		}

		// Test limit functionality
		t.Run("Limit", func(t *testing.T) {
			// Get first 2 events
			events, err := repo.ListEvents(ctx, viewerID, 2, 0)
			require.NoError(t, err)
			require.LessOrEqual(t, len(events), 2)

			// Get first 3 events
			events3, err := repo.ListEvents(ctx, viewerID, 3, 0)
			require.NoError(t, err)
			require.LessOrEqual(t, len(events3), 3)
			require.GreaterOrEqual(t, len(events3), len(events))

			// Test with limit 0 (should return no events)
			eventsZero, err := repo.ListEvents(ctx, viewerID, 0, 0)
			require.NoError(t, err)
			require.Empty(t, eventsZero)

			// Test with very large limit
			eventsLarge, err := repo.ListEvents(ctx, viewerID, 1000, 0)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(eventsLarge), 5) // Should contain our 5 events at minimum
		})

		// Test offset functionality
		t.Run("Offset", func(t *testing.T) {
			// Get all events first to understand total count
			allEvents, err := repo.ListEvents(ctx, viewerID, 1000, 0)
			require.NoError(t, err)
			totalCount := len(allEvents)

			if totalCount > 1 {
				// Test offset 1
				eventsOffset1, err := repo.ListEvents(ctx, viewerID, 1000, 1)
				require.NoError(t, err)
				require.Equal(t, totalCount-1, len(eventsOffset1))

				// Verify first event from offset 1 is different from first event without offset
				if len(allEvents) > 1 && len(eventsOffset1) > 0 {
					require.NotEqual(t, allEvents[0].ID, eventsOffset1[0].ID)
				}
			}

			if totalCount > 2 {
				// Test offset 2
				eventsOffset2, err := repo.ListEvents(ctx, viewerID, 1000, 2)
				require.NoError(t, err)
				require.Equal(t, totalCount-2, len(eventsOffset2))
			}

			// Test offset larger than total count
			eventsOffsetLarge, err := repo.ListEvents(ctx, viewerID, 1000, totalCount+10)
			require.NoError(t, err)
			require.Empty(t, eventsOffsetLarge)
		})

		// Test limit and offset combined
		t.Run("LimitAndOffset", func(t *testing.T) {
			// Get events in pages
			page1, err := repo.ListEvents(ctx, viewerID, 2, 0)
			require.NoError(t, err)

			page2, err := repo.ListEvents(ctx, viewerID, 2, 2)
			require.NoError(t, err)

			page3, err := repo.ListEvents(ctx, viewerID, 2, 4)
			require.NoError(t, err)

			// Verify no overlap between pages
			if len(page1) > 0 && len(page2) > 0 {
				for _, event1 := range page1 {
					for _, event2 := range page2 {
						require.NotEqual(t, event1.ID, event2.ID, "Events should not overlap between pages")
					}
				}
			}

			if len(page2) > 0 && len(page3) > 0 {
				for _, event2 := range page2 {
					for _, event3 := range page3 {
						require.NotEqual(t, event2.ID, event3.ID, "Events should not overlap between pages")
					}
				}
			}
		})

		// Test ordering (events should be ordered by starts_at DESC)
		t.Run("Ordering", func(t *testing.T) {
			events, err := repo.ListEvents(ctx, viewerID, 10, 0)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(events), 5) // Should have our 5 test events

			// Verify descending order by starts_at
			for i := 1; i < len(events); i++ {
				if events[i-1].StartsAt != nil && events[i].StartsAt != nil {
					require.True(t,
						events[i-1].StartsAt.After(*events[i].StartsAt) || events[i-1].StartsAt.Equal(*events[i].StartsAt),
						"Events should be ordered by starts_at DESC")
				}
			}
		})

		// Test negative values
		t.Run("NegativeValues", func(t *testing.T) {
			// Test negative limit (PostgreSQL rejects this)
			// We need to test this in a separate transaction since it will abort the current one
			err := repo.WithRollback(ctx, func(ctx context.Context) {
				_, err := repo.ListEvents(ctx, viewerID, -1, 0)
				require.Error(t, err)
				require.Contains(t, err.Error(), "LIMIT must not be negative")
			})
			require.NoError(t, err)

			// Test negative offset (PostgreSQL rejects this too)
			// We need to test this in a separate transaction since it will abort the current one
			err = repo.WithRollback(ctx, func(ctx context.Context) {
				_, err := repo.ListEvents(ctx, viewerID, 10, -1)
				require.Error(t, err)
				require.Contains(t, err.Error(), "OFFSET must not be negative")
			})
			require.NoError(t, err)
		})
	})
	require.NoError(t, err)
}

func TestEventsRepository_ListEvents_SubscriptionStatus(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	userRepo := users.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Create users
		creatorID := int64(111)
		viewerID := int64(222)
		subscriberID := int64(333)

		creator := &models.User{
			TelegramID: creatorID,
			Username:   "creator",
			FirstName:  "Creator",
		}

		viewer := &models.User{
			TelegramID: viewerID,
			Username:   "viewer",
			FirstName:  "Viewer",
		}

		subscriber := &models.User{
			TelegramID: subscriberID,
			Username:   "subscriber",
			FirstName:  "Subscriber",
		}

		_, err := userRepo.Create(ctx, creator)
		require.NoError(t, err)
		_, err = userRepo.Create(ctx, viewer)
		require.NoError(t, err)
		_, err = userRepo.Create(ctx, subscriber)
		require.NoError(t, err)

		// Create an event
		event := &models.Event{
			Title:     "Subscription Test Event",
			CreatedBy: &creatorID,
		}

		eventID, err := repo.CreateEvent(ctx, event)
		require.NoError(t, err)

		// Test viewing events as different users
		// Viewer (not subscribed)
		viewerEvents, err := repo.ListEvents(ctx, viewerID, 10, 0)
		require.NoError(t, err)

		var viewerEvent *models.EventEnriched
		for _, e := range viewerEvents {
			if e.ID == eventID {
				viewerEvent = &e
				break
			}
		}
		require.NotNil(t, viewerEvent)
		require.False(t, viewerEvent.IsSubscribed)
		require.NotNil(t, viewerEvent.ParticipantsCount)
		require.Equal(t, 0, *viewerEvent.ParticipantsCount)

		// TODO: Add subscription functionality tests when subscription repository is available
		// This would involve:
		// 1. Creating a subscription for subscriberID to eventID
		// 2. Verifying that ListEvents(ctx, subscriberID, ...) shows IsSubscribed: true
		// 3. Verifying that ParticipantsCount increases to 1
	})
	require.NoError(t, err)
}

func TestEventsRepository_GetUserEvents_EmptyResult(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Test with user that has no subscriptions
		userEvents, err := repo.GetUserEvents(ctx, int64(12345))
		require.NoError(t, err)
		require.Empty(t, userEvents)
	})
	require.NoError(t, err)
}

func TestEventsRepository_CreateEvent_ValidationError(t *testing.T) {
	pool := base.SetupTestPool(t)
	repo := eventsrepo.New(pool)
	ctx := context.Background()

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Test creating event with non-existent user (foreign key violation)
		nonExistentUserID := int64(999999)
		event := &models.Event{
			Title:     "Invalid Event",
			CreatedBy: &nonExistentUserID,
		}

		_, err := repo.CreateEvent(ctx, event)
		require.Error(t, err)
		require.Contains(t, err.Error(), "foreign key constraint")
	})
	require.NoError(t, err)
}
