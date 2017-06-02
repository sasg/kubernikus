package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewKubernikusAPI creates a new Kubernikus instance
func NewKubernikusAPI(spec *loads.Document) *KubernikusAPI {
	return &KubernikusAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		GetAPIHandler: GetAPIHandlerFunc(func(params GetAPIParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAPI has not yet been implemented")
		}),
		GetAPIV1ClustersHandler: GetAPIV1ClustersHandlerFunc(func(params GetAPIV1ClustersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetAPIV1Clusters has not yet been implemented")
		}),
		GetAPIV1ClustersNameHandler: GetAPIV1ClustersNameHandlerFunc(func(params GetAPIV1ClustersNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetAPIV1ClustersName has not yet been implemented")
		}),
		PostAPIV1ClustersHandler: PostAPIV1ClustersHandlerFunc(func(params PostAPIV1ClustersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostAPIV1Clusters has not yet been implemented")
		}),

		// Applies when the "x-auth-token" header is set
		KeystoneAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (keystone) x-auth-token from header param [x-auth-token] has not yet been implemented")
		},
	}
}

/*KubernikusAPI the kubernikus API */
type KubernikusAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// KeystoneAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-auth-token provided in the header
	KeystoneAuth func(string) (interface{}, error)

	// GetAPIHandler sets the operation handler for the get API operation
	GetAPIHandler GetAPIHandler
	// GetAPIV1ClustersHandler sets the operation handler for the get API v1 clusters operation
	GetAPIV1ClustersHandler GetAPIV1ClustersHandler
	// GetAPIV1ClustersNameHandler sets the operation handler for the get API v1 clusters name operation
	GetAPIV1ClustersNameHandler GetAPIV1ClustersNameHandler
	// PostAPIV1ClustersHandler sets the operation handler for the post API v1 clusters operation
	PostAPIV1ClustersHandler PostAPIV1ClustersHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *KubernikusAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KubernikusAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KubernikusAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KubernikusAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KubernikusAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KubernikusAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KubernikusAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KubernikusAPI
func (o *KubernikusAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeystoneAuth == nil {
		unregistered = append(unregistered, "XAuthTokenAuth")
	}

	if o.GetAPIHandler == nil {
		unregistered = append(unregistered, "GetAPIHandler")
	}

	if o.GetAPIV1ClustersHandler == nil {
		unregistered = append(unregistered, "GetAPIV1ClustersHandler")
	}

	if o.GetAPIV1ClustersNameHandler == nil {
		unregistered = append(unregistered, "GetAPIV1ClustersNameHandler")
	}

	if o.PostAPIV1ClustersHandler == nil {
		unregistered = append(unregistered, "PostAPIV1ClustersHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KubernikusAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KubernikusAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "keystone":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.KeystoneAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *KubernikusAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *KubernikusAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KubernikusAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the kubernikus API
func (o *KubernikusAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KubernikusAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api"] = NewGetAPI(o.context, o.GetAPIHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/clusters"] = NewGetAPIV1Clusters(o.context, o.GetAPIV1ClustersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/clusters/{name}"] = NewGetAPIV1ClustersName(o.context, o.GetAPIV1ClustersNameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/clusters"] = NewPostAPIV1Clusters(o.context, o.PostAPIV1ClustersHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KubernikusAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *KubernikusAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}