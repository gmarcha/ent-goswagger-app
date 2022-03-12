// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/gmarcha/ent-goswagger-app/internal/clients/database"
	"github.com/gmarcha/ent-goswagger-app/internal/clients/env"
	"github.com/gmarcha/ent-goswagger-app/internal/clients/redis"
	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/restapi/operations"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/auth"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/event"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/eventtype"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/role"
	"github.com/gmarcha/ent-goswagger-app/internal/modules/user"
)

//go:generate swagger generate server --target ../../goswagger --name Tutor --spec ../../../docs/swagger.yaml --principal models.Principal --exclude-main

func configureFlags(api *operations.TutorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TutorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	env.Init()

	db := database.Init()
	rdb := redis.Init()

	auth.Init(api, db, rdb)
	role.Init(api, db)
	user.Init(api, db)
	event.Init(api, db)
	eventtype.Init(api, db)

	api.ServerShutdown = func() {
		db.Close()
		rdb.Close()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
