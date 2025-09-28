package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pkg/openapi"
	"go.uber.org/zap"
)

func (a *API) Events(ctx context.Context, request openapi.EventsRequestObject) (openapi.EventsResponseObject, error) {
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

	return openapi.Events200JSONResponse(utils.Map(events, func(e models.EventEnriched) openapi.Event {
		return openapi.Event{
			Id:                e.ID,
			Title:             e.Title,
			Description:       e.Description,
			Location:          e.Location,
			VideoUrl:          e.VideoURL,
			StartsAt:          e.StartsAt,
			IsSubscribed:      e.IsSubscribed,
			ParticipantsCount: e.ParticipantsCount,
			Djs: utils.Map(e.DJs, func(dj models.DJ) openapi.DJ {
				return openapi.DJ{
					StageName: dj.StageName,
					AvatarUrl: dj.AvatarURL,
					Socials: utils.Map(dj.Socials, func(social models.Social) openapi.Social {
						return openapi.Social{
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
