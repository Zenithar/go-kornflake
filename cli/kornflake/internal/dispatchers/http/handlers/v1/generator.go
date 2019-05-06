package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "go.zenithar.org/kornflake/internal/services/v1"
	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
	"go.zenithar.org/pkg/web/respond"
)

type generatorCtrl struct {
	snfg   v1.SnowflakeGenerator
	bfg10k v1.BigflakeGenerator // Quake inside !

	snfgReq *snowflakev1.GetRequest
	bfgReq  *bigflakev1.GetRequest
}

// -----------------------------------------------------------------------------

// GeneratorRoutes returns identifier generator related API
func GeneratorRoutes(snowflakes v1.SnowflakeGenerator, bigflakes v1.BigflakeGenerator) http.Handler {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &generatorCtrl{
		snfg:    snowflakes,
		snfgReq: &snowflakev1.GetRequest{},
		bfg10k:  bigflakes,
		bfgReq:  &bigflakev1.GetRequest{},
	}

	// Map routes
	r.Get("/snowflake", ctrl.snowflake())
	r.Post("/snowflake", ctrl.snowflake())
	r.Get("/bigflake", ctrl.bigflake())
	r.Post("/bigflake", ctrl.bigflake())

	// Return router
	return r
}

// -----------------------------------------------------------------------------

func (c *generatorCtrl) snowflake() http.HandlerFunc {
	type response struct {
		ID string `json:"id"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		res, err := c.snfg.Get(ctx, c.snfgReq)
		if err != nil {
			respond.WithError(w, r, http.StatusInternalServerError, "Unable to generate identifier")
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			ID: res.Identifier,
		})
	}
}

func (c *generatorCtrl) bigflake() http.HandlerFunc {
	type response struct {
		ID string `json:"id"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		res, err := c.bfg10k.Get(ctx, c.bfgReq)
		if err != nil {
			respond.WithError(w, r, http.StatusInternalServerError, "Unable to generate identifier")
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			ID: res.Identifier,
		})
	}
}
