package restapi

import (
	"crypto/tls"
	"github.com/craguilar/cidm/restapi/operations"
	"github.com/craguilar/cidm/restapi/operations/login"
	"github.com/craguilar/cidm/restapi/operations/resource"
	"github.com/craguilar/cidm/restapi/operations/token"
	"github.com/craguilar/cidm/services"
	"github.com/craguilar/cidm/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/gorilla/handlers"
	"github.com/tylerb/graceful"
	"net/http"
)

//go:generate swagger generate server --target $GOPATH\src\github.com\craguilar\cidm --name cidm --spec $GOPATH\src\github.com\craguilar\cidm\swagger\swagger.yaml

func configureAPI(api *operations.CidmAPI) http.Handler {

	// configure the api here
	api.ServeError = errors.ServeError

	// Custom logging Default one is log.Printf Expected interface func(string, ...interface{})
	api.Logger = utils.Logger().Infof

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Header authorization Applies when the Authorization header is set with the Basic scheme
	//TODO Implement authorizer
	api.APIAuthorizer = &services.CAuthorizer{}
	api.BearerAuth = services.HandleBearerAuth

	// Get Login handler
	api.LoginGetLoginDetailsByIDHandler = login.GetLoginDetailsByIDHandlerFunc(services.GetLoginDetailsById)
	api.LoginGetLoginDetailsHandler = login.GetLoginDetailsHandlerFunc(services.GetLoginDetails)
	api.LoginAddLoginDetailsHandler = login.AddLoginDetailsHandlerFunc(services.AddLoginDetails)
	api.LoginUpdateLoginDetailsHandler = login.UpdateLoginDetailsHandlerFunc(services.UpdateLoginDetails)
	// Login & callback Handler
	api.LoginGetLoginMockHandler = login.GetLoginMockHandlerFunc(services.HandleLoginMockRequest)
	api.LoginGetLoginHandler = login.GetLoginHandlerFunc(services.HandleLoginRequest)
	api.LoginGetLoginCallbackHandler = login.GetLoginCallbackHandlerFunc(services.HandleCallback)
	//Configuration
	api.LoginGetLoginConfigDetailsHandler = login.GetLoginConfigDetailsHandlerFunc(services.GetLoginConfigDetails)
	api.LoginUpdateLoginConfigDetailsHandler = login.UpdateLoginConfigDetailsHandlerFunc(services.UpdateLoginConfigDetails)
	//Resources
	api.ResourceGetResourceByIDHandler = resource.GetResourceByIDHandlerFunc(services.GetResourceById)
	api.ResourceAddResourceHandler = resource.AddResourceHandlerFunc(services.AddResource)
	api.ResourceUpdateResourceHandler = resource.UpdateResourceHandlerFunc(services.UpdateResource)
	api.ResourceDeleteResourceHandler = resource.DeleteResourceHandlerFunc(services.DeleteResource)
	//Token Authentication
	api.TokenTokenValidateHandler = token.TokenValidateHandlerFunc(services.HandleTokenValidate)
	api.TokenTokenValidatPayloadHandler = token.TokenValidatPayloadHandlerFunc(services.HandleTokenPayloadValidate)
	//Token Authorization
	api.TokenGetUserPermissionHandler = token.GetUserPermissionHandlerFunc(services.GetUserPermission)
	api.TokenAddUserPermissionHandler = token.AddUserPermissionHandlerFunc(services.AddUserPermission)
	api.TokenUpdateUserPermissionHandler = token.UpdateUserPermissionHandlerFunc(services.UpdateUserPermission)
	api.TokenDeleteUserPermissionHandler = token.DeleteUserPermissionHandlerFunc(services.DeleteUserPermission)

	//TODO implement Accounting for AAA compliance.

	api.ServerShutdown = customShutdown

	return setupGlobalMiddleware(api.Serve(setupMiddleware))
}

func customShutdown() {
	//Custom shutdown
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddleware(handler http.Handler) http.Handler {

	cors := handlers.AllowedOrigins([]string{"*"})
	//TODO Implement micro cache for efficient return of web service rest API -
	/*
	   //cacheenabled, err := strconv.ParseBool(services.GetVariable("MICROCACHENOTENABLED"))
	   Microcache github.com/httpimp/microcache
	   cache := microcache.New(microcache.Config{
	     Nocache:              true,
	     Timeout:              50 * time.Millisecond, //Max time to return response
	     TTL:                  25 * time.Second,      //
	     StaleIfError:         3600 * time.Second,
	     StaleRecache:         false,
	     StaleWhileRevalidate: 50 * time.Second,
	     CollapsedForwarding:  true,
	     HashQuery:            true,
	     QueryIgnore:          []string{},
	     Exposed:              true,
	     Driver:               microcache.NewDriverLRU(1e4),
	     Compressor:           microcache.CompressorSnappy{},
	   })
	*/
	//return cache.Middleware(handlers.CORS(cors)(handler))
	return handlers.CORS(cors)(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func configureFlags(api *operations.CidmAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}
