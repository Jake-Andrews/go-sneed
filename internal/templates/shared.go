package templates

import (
	"context"
	"strconv"

    "go-sneed/internal/models"
)

func AuthenticatedUser(ctx context.Context) models.AuthenticatedUser {
	user, ok := ctx.Value(models.UserReqKey).(models.AuthenticatedUser)
	if !ok {
		return models.AuthenticatedUser{}
	}
	return user
}

func String(i int) string {
	return strconv.Itoa(i)
}
