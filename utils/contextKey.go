package utils

import (
	"context"

	"github.com/govnocods/JWT-Authorization/models"
)

type ContextKey string
var UserCtxKey = ContextKey("user")

func GetCtxUser(ctx context.Context) *models.User {
	user, ok := ctx.Value(UserCtxKey).(*models.User)
	if !ok {
		return nil
	} else {
		return user
	}
}