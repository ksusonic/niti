//go:build integration

package subscriptions_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/storage/repository/subscriptions"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/stretchr/testify/require"
)

func setupTestPool(t *testing.T) *pgxpool.Pool {
	require.NoError(t, godotenv.Load("../../../../.env"))

	var cfg config.PostgresConfig
	require.NoError(t, env.Parse(&cfg))

	config, err := pgxpool.ParseConfig(cfg.DSN)
	require.NoError(t, err)

	// Disable prepared statement caching to avoid conflicts in tests
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}

func createTestUser(t *testing.T, pool *pgxpool.Pool, ctx context.Context) int64 {
	usersRepo := users.New(pool)
	telegramID := int64(uuid.New().ID())

	user := &models.User{
		TelegramID: telegramID,
		Username:   "testuser",
		FirstName:  "Test",
		LastName:   nil,
		AvatarURL:  nil,
		IsDJ:       false,
	}

	_, err := usersRepo.Create(ctx, user)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = usersRepo.Delete(ctx, telegramID)
	})

	return telegramID
}

func createTestEvent(t *testing.T, pool *pgxpool.Pool, ctx context.Context, createdBy int64) int32 {
	var eventID int32
	err := pool.QueryRow(
		ctx,
		`INSERT INTO events (title, description, location, starts_at, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		"Test Event",
		"Test event description",
		"Test location",
		"2024-12-31 23:59:59",
		createdBy,
	).Scan(&eventID)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM events WHERE id = $1", eventID)
	})

	return eventID
}

func TestSubscriptionsRepository_CreateAndDelete(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Create test event
	eventID := createTestEvent(t, pool, ctx, userID)

	// Test CREATE subscription
	err := repo.CreateSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)

	// Verify subscription exists
	exists, err := repo.GetSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)
	require.True(t, exists)

	// Test CREATE subscription again (should not error due to ON CONFLICT DO NOTHING)
	err = repo.CreateSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)

	// Verify subscription still exists
	exists, err = repo.GetSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)
	require.True(t, exists)

	// Test DELETE subscription
	err = repo.DeleteSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)

	// Verify subscription no longer exists
	exists, err = repo.GetSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)
	require.False(t, exists)

	// Test DELETE subscription again (should not error)
	err = repo.DeleteSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)
}

func TestSubscriptionsRepository_CreateWithNonExistentUser(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user for event creation
	userID := createTestUser(t, pool, ctx)
	eventID := createTestEvent(t, pool, ctx, userID)

	// Try to create subscription with non-existent user
	nonExistentUserID := int64(999999999)
	err := repo.CreateSubscription(ctx, nonExistentUserID, int(eventID))
	require.Error(t, err)
	// Should fail due to foreign key constraint
}

func TestSubscriptionsRepository_CreateWithNonExistentEvent(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Try to create subscription with non-existent event
	nonExistentEventID := int(999999999)
	err := repo.CreateSubscription(ctx, userID, nonExistentEventID)
	require.Error(t, err)
	// Should fail due to foreign key constraint
}

func TestSubscriptionsRepository_GetNonExistentSubscription(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user and event but no subscription
	userID := createTestUser(t, pool, ctx)
	eventID := createTestEvent(t, pool, ctx, userID)

	// Check if subscription exists (it shouldn't)
	exists, err := repo.GetSubscription(ctx, userID, int(eventID))
	require.NoError(t, err)
	require.False(t, exists)
}

func TestSubscriptionsRepository_GetSubscriptionsByEventIDs(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Create multiple test events
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)
	eventID3 := createTestEvent(t, pool, ctx, userID)

	// Subscribe to events 1 and 3
	err := repo.CreateSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID3))
	require.NoError(t, err)

	// Test getting subscriptions for all events
	eventIDs := []int{int(eventID1), int(eventID2), int(eventID3)}
	subscribedEventIDs, err := repo.GetSubscriptionsByEventIDs(ctx, userID, eventIDs)
	require.NoError(t, err)
	require.Len(t, subscribedEventIDs, 2)
	require.Contains(t, subscribedEventIDs, int(eventID1))
	require.Contains(t, subscribedEventIDs, int(eventID3))
	require.NotContains(t, subscribedEventIDs, int(eventID2))

	// Test getting subscriptions for subset of events
	subsetEventIDs := []int{int(eventID1), int(eventID2)}
	subscribedSubsetEventIDs, err := repo.GetSubscriptionsByEventIDs(ctx, userID, subsetEventIDs)
	require.NoError(t, err)
	require.Len(t, subscribedSubsetEventIDs, 1)
	require.Contains(t, subscribedSubsetEventIDs, int(eventID1))
	require.NotContains(t, subscribedSubsetEventIDs, int(eventID2))
}

func TestSubscriptionsRepository_GetSubscriptionsByEventIDs_EmptyArray(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Test with empty event IDs array
	eventIDs := []int{}
	subscribedEventIDs, err := repo.GetSubscriptionsByEventIDs(ctx, userID, eventIDs)
	require.NoError(t, err)
	require.Empty(t, subscribedEventIDs)
}

func TestSubscriptionsRepository_GetSubscriptionsByEventIDs_NoSubscriptions(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Create test events but don't subscribe to any
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)

	// Test getting subscriptions when user has no subscriptions
	eventIDs := []int{int(eventID1), int(eventID2)}
	subscribedEventIDs, err := repo.GetSubscriptionsByEventIDs(ctx, userID, eventIDs)
	require.NoError(t, err)
	require.Empty(t, subscribedEventIDs)
}

func TestSubscriptionsRepository_GetUserSubscriptions(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Create multiple test events
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)
	eventID3 := createTestEvent(t, pool, ctx, userID)

	// Initially user has no subscriptions
	subscribedEventIDs, err := repo.GetUserSubscriptions(ctx, userID)
	require.NoError(t, err)
	require.Empty(t, subscribedEventIDs)

	// Subscribe to events 1 and 3
	err = repo.CreateSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID3))
	require.NoError(t, err)

	// Get all user subscriptions
	subscribedEventIDs, err = repo.GetUserSubscriptions(ctx, userID)
	require.NoError(t, err)
	require.Len(t, subscribedEventIDs, 2)
	require.Contains(t, subscribedEventIDs, int(eventID1))
	require.Contains(t, subscribedEventIDs, int(eventID3))
	require.NotContains(t, subscribedEventIDs, int(eventID2))

	// Subscribe to event 2
	err = repo.CreateSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)

	// Get all user subscriptions again
	subscribedEventIDs, err = repo.GetUserSubscriptions(ctx, userID)
	require.NoError(t, err)
	require.Len(t, subscribedEventIDs, 3)
	require.Contains(t, subscribedEventIDs, int(eventID1))
	require.Contains(t, subscribedEventIDs, int(eventID2))
	require.Contains(t, subscribedEventIDs, int(eventID3))

	// Delete one subscription
	err = repo.DeleteSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)

	// Verify subscription is removed
	subscribedEventIDs, err = repo.GetUserSubscriptions(ctx, userID)
	require.NoError(t, err)
	require.Len(t, subscribedEventIDs, 2)
	require.Contains(t, subscribedEventIDs, int(eventID1))
	require.Contains(t, subscribedEventIDs, int(eventID3))
	require.NotContains(t, subscribedEventIDs, int(eventID2))
}

func TestSubscriptionsRepository_GetAllSubscriptions(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test users
	userID1 := createTestUser(t, pool, ctx)
	userID2 := createTestUser(t, pool, ctx)

	// Create test events
	eventID1 := createTestEvent(t, pool, ctx, userID1)
	eventID2 := createTestEvent(t, pool, ctx, userID1)
	eventID3 := createTestEvent(t, pool, ctx, userID2)

	// Initially no subscriptions
	subscriptions, err := repo.GetAllSubscriptions(ctx, nil, 0, 10)
	require.NoError(t, err)
	initialCount := len(subscriptions)

	// Create subscriptions with some delay to ensure different creation times
	err = repo.CreateSubscription(ctx, userID1, int(eventID1))
	require.NoError(t, err)

	err = repo.CreateSubscription(ctx, userID2, int(eventID2))
	require.NoError(t, err)

	err = repo.CreateSubscription(ctx, userID1, int(eventID3))
	require.NoError(t, err)

	// Get all subscriptions
	subscriptions, err = repo.GetAllSubscriptions(ctx, nil, 0, 10)
	require.NoError(t, err)
	require.Len(t, subscriptions, initialCount+3)

	// Verify subscriptions are ordered by creation time (DESC)
	if len(subscriptions) >= 3 {
		// The last 3 subscriptions should be our test data (most recent first)
		lastThree := subscriptions[len(subscriptions)-3:]
		require.Equal(t, userID1, lastThree[0].UserID)
		require.Equal(t, int(eventID3), lastThree[0].EventID)
		require.Equal(t, userID2, lastThree[1].UserID)
		require.Equal(t, int(eventID2), lastThree[1].EventID)
		require.Equal(t, userID1, lastThree[2].UserID)
		require.Equal(t, int(eventID1), lastThree[2].EventID)
	}

	// Test pagination with limit
	subscriptionsPage1, err := repo.GetAllSubscriptions(ctx, nil, 0, 2)
	require.NoError(t, err)
	require.LessOrEqual(t, len(subscriptionsPage1), 2)

	// Test pagination with offset
	subscriptionsPage2, err := repo.GetAllSubscriptions(ctx, nil, 1, 2)
	require.NoError(t, err)
	require.LessOrEqual(t, len(subscriptionsPage2), 2)

	// If we have enough data, verify offset works
	if len(subscriptions) > 1 {
		require.NotEqual(t, subscriptionsPage1[0], subscriptionsPage2[0])
	}
}

func TestSubscriptionsRepository_GetAllSubscriptions_Pagination(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Test default limit when limit is 0
	subscriptions, err := repo.GetAllSubscriptions(ctx, nil, 0, 0)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)

	// Test negative offset (should be corrected to 0)
	subscriptions, err = repo.GetAllSubscriptions(ctx, nil, -5, 10)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)

	// Test negative limit (should use default)
	subscriptions, err = repo.GetAllSubscriptions(ctx, nil, 0, -1)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)
}

func TestSubscriptionsRepository_GetSubscriptionsCount(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Get initial count
	initialCount, err := repo.GetSubscriptionsCount(ctx, nil)
	require.NoError(t, err)
	require.GreaterOrEqual(t, initialCount, int64(0))

	// Create test user and events
	userID := createTestUser(t, pool, ctx)
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)

	// Create subscriptions
	err = repo.CreateSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)

	// Verify count increased
	newCount, err := repo.GetSubscriptionsCount(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, initialCount+2, newCount)

	// Delete one subscription
	err = repo.DeleteSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)

	// Verify count decreased
	finalCount, err := repo.GetSubscriptionsCount(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, initialCount+1, finalCount)
}

func TestSubscriptionsRepository_GetAllSubscriptions_EmptyDatabase(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// This test assumes we might have some existing data, so we just test the method works
	subscriptions, err := repo.GetAllSubscriptions(ctx, nil, 0, 10)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)

	count, err := repo.GetSubscriptionsCount(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, int64(len(subscriptions)), count)
}

func TestSubscriptionsRepository_GetAllSubscriptions_WithUserFilter(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Create multiple test events
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)
	eventID3 := createTestEvent(t, pool, ctx, userID)
	eventID4 := createTestEvent(t, pool, ctx, userID)

	// Create subscriptions
	err := repo.CreateSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID3))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID4))
	require.NoError(t, err)

	// Test getting all user subscriptions
	userSubscriptions, err := repo.GetAllSubscriptions(ctx, &userID, 0, 10)
	require.NoError(t, err)
	require.Len(t, userSubscriptions, 4)

	// Verify all subscriptions belong to the user
	for _, sub := range userSubscriptions {
		require.Equal(t, userID, sub.UserID)
	}

	// Test pagination with limit
	page1, err := repo.GetAllSubscriptions(ctx, &userID, 0, 2)
	require.NoError(t, err)
	require.Len(t, page1, 2)

	page2, err := repo.GetAllSubscriptions(ctx, &userID, 2, 2)
	require.NoError(t, err)
	require.Len(t, page2, 2)

	// Verify pages don't overlap
	page1EventIDs := make(map[int]bool)
	for _, sub := range page1 {
		page1EventIDs[sub.EventID] = true
	}
	for _, sub := range page2 {
		require.False(t, page1EventIDs[sub.EventID], "Pages should not overlap")
	}

	// Test empty result for non-existent user
	nonExistentUserID := int64(999999999)
	emptyResult, err := repo.GetAllSubscriptions(ctx, &nonExistentUserID, 0, 10)
	require.NoError(t, err)
	require.Empty(t, emptyResult)
}

func TestSubscriptionsRepository_GetSubscriptionsCount_WithUserFilter(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	// Create test user
	userID := createTestUser(t, pool, ctx)

	// Initially user has no subscriptions
	count, err := repo.GetSubscriptionsCount(ctx, &userID)
	require.NoError(t, err)
	require.Equal(t, int64(0), count)

	// Create test events and subscriptions
	eventID1 := createTestEvent(t, pool, ctx, userID)
	eventID2 := createTestEvent(t, pool, ctx, userID)
	eventID3 := createTestEvent(t, pool, ctx, userID)

	err = repo.CreateSubscription(ctx, userID, int(eventID1))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)
	err = repo.CreateSubscription(ctx, userID, int(eventID3))
	require.NoError(t, err)

	// Verify count
	count, err = repo.GetSubscriptionsCount(ctx, &userID)
	require.NoError(t, err)
	require.Equal(t, int64(3), count)

	// Delete one subscription
	err = repo.DeleteSubscription(ctx, userID, int(eventID2))
	require.NoError(t, err)

	// Verify count decreased
	count, err = repo.GetSubscriptionsCount(ctx, &userID)
	require.NoError(t, err)
	require.Equal(t, int64(2), count)

	// Test count for non-existent user
	nonExistentUserID := int64(999999999)
	count, err = repo.GetSubscriptionsCount(ctx, &nonExistentUserID)
	require.NoError(t, err)
	require.Equal(t, int64(0), count)
}

func TestSubscriptionsRepository_PaginationEdgeCases(t *testing.T) {
	pool := setupTestPool(t)
	repo := subscriptions.New(pool)
	ctx := context.Background()

	userID := createTestUser(t, pool, ctx)

	// Test negative offset and limit (should be corrected)
	subscriptions, err := repo.GetAllSubscriptions(ctx, &userID, -5, -10)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)

	// Test zero limit (should use default)
	subscriptions, err = repo.GetAllSubscriptions(ctx, &userID, 0, 0)
	require.NoError(t, err)
	require.NotNil(t, subscriptions)

	// Test large offset (should return empty result)
	subscriptions, err = repo.GetAllSubscriptions(ctx, &userID, 1000000, 10)
	require.NoError(t, err)
	require.Empty(t, subscriptions)
}
