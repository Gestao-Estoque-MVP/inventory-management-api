package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/diogoX451/inventory-management-api/internal/graph/middleware"
	"github.com/diogoX451/inventory-management-api/internal/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
	userRole := middleware.CtxValue(ctx)
	if userRole.Role != role {
		return nil, &gqlerror.Error{
			Message: "Acess Denied Role",
		}
	}

	return next(ctx)
}
