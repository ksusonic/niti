package public

import (
	"context"
	"errors"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pkg/publicapi"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (a *API) GetProfile(ctx context.Context, _ publicapi.GetProfileRequestObject) (publicapi.GetProfileResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	eg, gCtx := errgroup.WithContext(ctx)

	var (
		user   *models.User
		events []models.EventEnriched
	)

	eg.Go(func() (err error) {
		user, err = a.usersRepo.Get(gCtx, tgUserID)
		return err
	})

	eg.Go(func() (err error) {
		events, err = a.eventsRepo.GetUserEvents(gCtx, tgUserID)
		return err
	})

	if err := eg.Wait(); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			a.logger.Error("not found user", zap.Int64("user_id", tgUserID))

			return publicapi.GetProfile404JSONResponse{Message: "profile not found"}, nil
		}

		a.logger.Error("get user", zap.Int64("user_id", tgUserID))

		return nil, err
	}

	return publicapi.GetProfile200JSONResponse{
		TelegramId: user.TelegramID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		AvatarUrl:  user.AvatarURL,
		IsDj:       user.IsDJ,
		Subscriptions: utils.Map(events, func(event models.EventEnriched) publicapi.Event {
			return publicapi.Event{
				Description:       event.Description,
				Id:                event.ID,
				Location:          event.Location,
				ParticipantsCount: event.ParticipantsCount,
				StartsAt:          event.StartsAt,
				Title:             event.Title,
				VideoUrl:          event.VideoURL,
				IsSubscribed:      event.IsSubscribed,
				Djs: utils.Map(event.DJs, func(dj models.DJ) publicapi.DJ {
					return publicapi.DJ{
						StageName: dj.StageName,
						AvatarUrl: dj.AvatarURL,
						Socials: utils.Map(dj.Socials, func(social models.Social) publicapi.Social {
							return publicapi.Social{
								Name: social.Name,
								Url:  social.URL,
								Icon: social.Icon,
							}
						}),
					}
				}),
			}
		}),
	}, nil
}
