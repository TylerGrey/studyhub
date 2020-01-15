package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/TylerGrey/studyhub/api/app/loader"
	"github.com/TylerGrey/studyhub/api/app/resolvers/args"
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

	contentType := strings.SplitN(r.Header.Get("Content-Type"), ";", 2)[0]
	if r.Method == "POST" {
		switch contentType {
		case "text/plain", "application/json":
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case "multipart/form-data":
			if err := r.ParseMultipartForm(4 * 1024 * 1024); err != nil {
				panic(err)
			}

			// Unmarshal uploads
			var uploads = map[args.FileInput][]string{}
			var uploadsMap = map[string][]string{}
			if err := json.Unmarshal([]byte(r.Form.Get("map")), &uploadsMap); err != nil {
				panic(err)
			} else {
				for key, path := range uploadsMap {
					if file, header, err := r.FormFile(key); err != nil {
						panic(err)
					} else {
						uploads[args.FileInput{
							File:     file,
							Size:     header.Size,
							Filename: header.Filename,
						}] = path
					}
				}
			}

			var operations interface{}
			// Unmarshal operations
			if err := json.Unmarshal([]byte(r.Form.Get("operations")), &operations); err != nil {
				panic(err)
			}

			// set uploads to operations
			for file, paths := range uploads {
				for _, path := range paths {
					if err := set(file, operations, path); err != nil {
						panic(err)
					}
				}
			}

			switch data := operations.(type) {
			case map[string]interface{}:
				if value, ok := data["operationName"]; ok && value != nil {
					params.OperationName = value.(string)
				}
				if value, ok := data["query"]; ok && value != nil {
					params.Query = value.(string)
				}
				if value, ok := data["variables"]; ok && value != nil {
					params.Variables = value.(map[string]interface{})
				}
			case []interface{}:
				result := make([]interface{}, len(data))
				for _, operation := range data {
					data := operation.(map[string]interface{})
					if value, ok := data["operationName"]; ok {
						params.OperationName = value.(string)
					}
					if value, ok := data["query"]; ok {
						params.Query = value.(string)
					}
					if value, ok := data["variables"]; ok {
						params.Variables = value.(map[string]interface{})
					}
				}
				if err := json.NewEncoder(w).Encode(result); err != nil {
					panic(err)
				}
			default:
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}

	ctx := h.Loaders.Attach(r.Context()) // Attach dataloaders onto the request context.
	response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func set(v interface{}, m interface{}, path string) error {
	var parts []interface{}
	for _, p := range strings.Split(path, ".") {
		if isNumber, err := regexp.MatchString(`\d+`, p); err != nil {
			return err
		} else if isNumber {
			index, _ := strconv.Atoi(p)
			parts = append(parts, index)
		} else {
			parts = append(parts, p)
		}
	}
	for i, p := range parts {
		last := i == len(parts)-1
		switch idx := p.(type) {
		case string:
			if last {
				m.(map[string]interface{})[idx] = v
			} else {
				m = m.(map[string]interface{})[idx]
			}
		case int:
			if last {
				m.([]interface{})[idx] = v
			} else {
				m = m.([]interface{})[idx]
			}
		}
	}
	return nil
}
