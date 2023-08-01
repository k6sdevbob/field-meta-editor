package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/foghorn-tech/kanaries-dsl/parser"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

type api struct {
	db *sql.DB
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func main() {
	// Create a mock database connection
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create api and init database with filling example data
	app := &api{
		db: db,
	}
	if err = app.InitDB(); err != nil {
		log.Fatalf("%v", err)
	}

	// Create router and listen
	router := mux.NewRouter()

	router.Handle("/meta/query", corsMiddleware(http.HandlerFunc(app.QueryMetaHandler))).Methods("POST", http.MethodOptions)
	router.Handle("/meta/update", corsMiddleware(http.HandlerFunc(app.UpdateMetaHandler))).Methods("POST", http.MethodOptions)
	router.Handle("/dsl/query", corsMiddleware(http.HandlerFunc(app.QueryDatesetHandler))).Methods("POST", http.MethodOptions)

	http.ListenAndServe(":23402", router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *api) QueryMetaHandler(w http.ResponseWriter, r *http.Request) {
	type updateRequest struct {
		DataID string `json:"dataId"`
	}
	var ur updateRequest
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(ur.DataID)

	meta, err := a.QueryMeta(id)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := SuccessResponse{Success: true, Data: meta}
	json.NewEncoder(w).Encode(resp)
}

func (a *api) UpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	type updateRequest struct {
		DataID string `json:"dataId"`
		Meta   []Meta `json:"meta"`
	}

	var ur updateRequest
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(ur.DataID)

	err = a.UpdateMeta(id, ur.Meta)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Meta updated successfully")

	meta, err := a.QueryMeta(id)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := SuccessResponse{Success: true, Data: meta}
	json.NewEncoder(w).Encode(resp)
}

func (a *api) QueryDatesetHandler(w http.ResponseWriter, r *http.Request) {
	type queryRequest struct {
		DataID  string         `json:"dataId"`
		Payload parser.Payload `json:"payload"`
	}
	var ur queryRequest
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(ur.DataID)

	queryDataset, err := a.QueryDataset(id)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	baseParser := parser.BaseParser{}
	dataset := parser.Dataset{
		Table: queryDataset.Name,
		Type:  parser.DatasetType(queryDataset.Type),
	}

	sql, err := baseParser.Parse(dataset, ur.Payload)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := a.QueryDatasource(sql)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := ErrorResponse{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := SuccessResponse{Success: true, Data: res}
	json.NewEncoder(w).Encode(resp)
}
