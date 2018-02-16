package authn

import (
	"context"
	"log"

	"github.com/twitchtv/twirp"
)

func checkAuthN(ctx context.Context) bool {
	authN := ctx.Value("Authorization").(string)
	log.Printf("Authorization: %v", authN)
	return false
}

//NewAuthNServerHooks Checks authN
func NewAuthNServerHooks() *twirp.ServerHooks {
	hooks := &twirp.ServerHooks{}

	hooks.RequestReceived = func(ctx context.Context) (context.Context, error) {
		if !checkAuthN(ctx) {
			twerr := twirp.NewError(twirp.Unauthenticated, "Unauthenticated")
			return nil, twerr
		}
		return ctx, nil
	}

	return hooks
}
