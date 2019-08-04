/*
 * HoloStorage Accessor API
 *
 * API to access holograms and metadata from HoloStorage
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

type AccessorConfig struct {
	FhirURL       string `json:"fhirUrl"`
	BlobStoreName string `json:"blobStoreName"`
	BlobKey       string `json:"blobKey"`
}

var accessorConfig AccessorConfig

// NewRouter returns a new router.
func NewRouter(confFile string) *gin.Engine {
	err := LoadConfiguration(confFile, &accessorConfig)
	if err != nil {
		log.Printf("Error with configuration file: " + confFile)
		log.Fatalln(err)
	}

	log.Printf("Fhir URL: %s", accessorConfig.FhirURL)
	log.Printf("Blob Store: %s", accessorConfig.BlobStoreName)

	router := gin.Default()

	router.Static("/app/1.0.0/ui/", "./third_party/swaggerui")

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/api/1.0.0/",
		Index,
	},

	{
		"AuthorsAidGet",
		http.MethodGet,
		"/api/1.0.0/authors/:aid",
		AuthorsAidGet,
	},

	{
		"AuthorsAidPut",
		http.MethodPut,
		"/api/1.0.0/authors/:aid",
		AuthorsAidPut,
	},

	{
		"AuthorsGet",
		http.MethodGet,
		"/api/1.0.0/authors",
		AuthorsGet,
	},

	{
		"HologramsGet",
		http.MethodGet,
		"/api/1.0.0/holograms",
		HologramsGet,
	},

	{
		"HologramsHidDelete",
		http.MethodDelete,
		"/api/1.0.0/holograms/:hid",
		HologramsHidDelete,
	},

	{
		"HologramsHidDownloadGet",
		http.MethodGet,
		"/api/1.0.0/holograms/:hid/download",
		HologramsHidDownloadGet,
	},

	{
		"HologramsHidGet",
		http.MethodGet,
		"/api/1.0.0/holograms/:hid",
		HologramsHidGet,
	},

	{
		"HologramsPost",
		http.MethodPost,
		"/api/1.0.0/holograms",
		HologramsPost,
	},

	{
		"PatientsGet",
		http.MethodGet,
		"/api/1.0.0/patients",
		PatientsGet,
	},

	{
		"PatientsPidGet",
		http.MethodGet,
		"/api/1.0.0/patients/:pid",
		PatientsPidGet,
	},

	{
		"PatientsPidPut",
		http.MethodPut,
		"/api/1.0.0/patients/:pid",
		PatientsPidPut,
	},
}
