package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/gql/generated"
	"github.com/japiirainen/go-oluet-api/gql/resolvers"
	"github.com/japiirainen/go-oluet-api/handlers"
	"github.com/japiirainen/go-oluet-api/middleware"
	log "github.com/sirupsen/logrus"
)

const defaultPort = "5000"

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	psql := db.Connect()
	defer psql.CloseConnection()
	r := mux.NewRouter().StrictSlash(false)

	//Home routes
	home := r.Path("/").Subrouter()
	home.Methods("GET").HandlerFunc(handlers.HomeHandler)

	//internal routes
	internalBase := mux.NewRouter()
	r.PathPrefix("/internal").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middleware.CheckAuth),
		negroni.Wrap(internalBase),
	))
	internal := internalBase.PathPrefix("/internal").Subrouter()
	internal.Methods("GET").HandlerFunc(handlers.GetInternal)
	internal.Methods("POST").HandlerFunc(handlers.DailyUpdate)

	//graphql routes
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		DB: psql,
	}}))
	r.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
	log.Infof("playground @ -> http://localhost:%s/graphql", port)
	log.Infof("query graphql @ -> http://localhost:%s/query", port)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),

		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	s.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
