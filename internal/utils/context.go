package utils

import(
	"context"
	"doctor-booking-system/internal/db/models"
)

type contextKey string

const userCtxKey = contextKey("currentUser")

func WithUser(ctx context.Context,user *models.User) context.Context{
	return context.WithValue(ctx,userCtxKey,user)
}

func GetUserFromContext(ctx context.Context) (*models.User,bool){
	user,ok := ctx.Value(userCtxKey).(*models.User)
	return user,ok
}