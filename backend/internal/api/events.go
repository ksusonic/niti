package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

// under auth middleware
func (a *API) Events(ctx context.Context, request genapi.EventsRequestObject) (genapi.EventsResponseObject, error) {
	tgUserID := models.MustTGUserID(ctx)

	limit := utils.Deref(request.Params.Limit)
	if limit <= 0 {
		limit = 30
	} else if limit > 100 {
		limit = 100
	}

	offset := utils.Deref(request.Params.Offset)

	events, err := a.eventsRepo.ListEvents(ctx, tgUserID, limit, offset)
	if err != nil {
		a.logger.Error("list events", zap.Error(err), zap.Int("limit", limit), zap.Int("offset", offset))

		return nil, err
	}

	return genapi.Events200JSONResponse(utils.Map(events, func(e models.EventEnriched) genapi.Event {
		return genapi.Event{
			Id:                e.ID,
			Title:             e.Title,
			Description:       e.Description,
			Location:          e.Location,
			VideoUrl:          e.VideoURL,
			StartsAt:          e.StartsAt,
			IsSubscribed:      e.IsSubscribed,
			ParticipantsCount: e.ParticipantsCount,
			Djs: utils.Map(e.DJs, func(dj models.DJ) genapi.DJ {
				return genapi.DJ{
					StageName: dj.StageName,
					AvatarUrl: dj.AvatarURL,
					Socials: utils.Map(dj.Socials, func(social models.Social) genapi.Social {
						return genapi.Social{
							Name: social.Name,
							Url:  social.URL,
							Icon: social.Icon,
						}
					}),
				}
			}),
		}
	})), nil
}
