package gateway

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"github.com/purwandi/platform/gateway/resolver"
	"github.com/purwandi/platform/gateway/schema"
)

// Server ...
type Server struct {
	authorization *Authorization
	resolver      *resolver.Resolver
}

// NewServer is for creating resolver
func NewServer(r *resolver.Resolver) *Server {
	return &Server{
		resolver:      r,
		authorization: NewAuthorizationMiddleware(r.UserService),
	}
}

// Mount for mounting server
func (s *Server) Mount(e *echo.Echo) {
	e.POST("/query", s.GraphQL, s.authorization.Check)
}

// GraphQL ...
func (s *Server) GraphQL(c echo.Context) error {
	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
	}

	handler := &relay.Handler{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), s.resolver, opts...),
	}

	handler.ServeHTTP(c.Response(), c.Request())
	return nil
}
