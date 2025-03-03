package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"io"

	"net/http"
)

var AppLogger *zap.Logger

func main() {
	var err error
	AppLogger, err = zap.NewProduction()
	AppLogger.Sugar().Infow("AppLogger created")
	AppLogger.Sugar().Info("app started")
	if err != nil {
		panic(err)
	}
	DB := DBHandler{}
	err = DB.InitDB()
	if err != nil {
		AppLogger.Sugar().Fatal(err)
	}
	defer DB.Close()
	resp, err := http.Get("https://api.deps.dev/v3/projects/github.com%2Fcli%2Fcli")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			AppLogger.Sugar().Warn("failed to close response body", zap.Error(err))
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		AppLogger.Sugar().Error(err.Error())
	}
	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		AppLogger.Sugar().Error(err.Error())
		panic(err)
	}
	AppLogger.Sugar().Info("body unmarshalled")
	marshal, err := json.Marshal(response.ScoreCard.Checks)
	if err != nil {
		AppLogger.Sugar().Error(err.Error())
		panic(err)
	}
	err = DB.InitPackage(response.ProjectKey.ID, string(body), string(marshal))
	if err != nil {
		AppLogger.Sugar().Error(err.Error())
	}
	handlers := New(&DB)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.HandleGet)
	router.Post("/", handlers.HandlePost)
	router.Delete("/", handlers.HandleDelete)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		AppLogger.Sugar().Error(err.Error())
		return
	}
}
