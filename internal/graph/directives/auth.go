package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/diogoX451/inventory-management-api/internal/graph/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	log.Printf("[DEBUG] Auth", tokenData)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
