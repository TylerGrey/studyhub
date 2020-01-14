package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TylerGrey/studyhub/api/app/loader"
	"github.com/graph-gophers/graphql-go"
)

// GraphQL GraphQL 핸들러
type GraphQL struct {
	Schema  *graphql.Schema
	Loaders loader.Collection
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx       := h.Loaders.Attach(r.Context()) // Attach dataloaders onto the request context.
	response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
