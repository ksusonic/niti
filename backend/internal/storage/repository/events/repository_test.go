//go:build integration

package events_test

import (
	"context"
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
		user := &models.User{
			TelegramID: creatorTelegramID,
			Username:   "testuser",
			FirstName:  "Test",
		}

		_, err := userRepo.Create(ctx, user)
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

		// LIST events to verify creation
		events, err := repo.ListEvents(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(events), 1) // May have other events in DB

		// Find our event in the list
		var foundEvent *models.Event
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
		eventsAfterSecond, err := repo.ListEvents(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(eventsAfterSecond), 2) // Should have at least our 2 events

		// DELETE first event
		err = repo.DeleteEvent(ctx, eventID)
		require.NoError(t, err)

		// LIST after delete
		eventsAfterDelete, err := repo.ListEvents(ctx)
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
		user := &models.User{
			TelegramID: creatorID,
			Username:   "listuser",
			FirstName:  "List",
		}

		_, err := userRepo.Create(ctx, user)
		require.NoError(t, err)

		// Get initial count
		initialEvents, err := repo.ListEvents(ctx)
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
		events, err := repo.ListEvents(ctx)
		require.NoError(t, err)
		require.Equal(t, initialCount+1, len(events))

		// Verify our event is in the list
		var foundEvent bool
		for _, e := range events {
			if e.ID == eventID {
				foundEvent = true
				require.Equal(t, "List Test Event", e.Title)
				break
			}
		}
		require.True(t, foundEvent, "Created event should be in the list")
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
		allEvents, err := repo.ListEvents(ctx)
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
