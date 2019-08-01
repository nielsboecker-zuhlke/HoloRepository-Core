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
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorsAidGet - Get a single author metadata in HoloStorage
func AuthorsAidGet(c *gin.Context) {
	id := c.Param("aid")
	fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Practitioner/"+id)
	result := SingleFHIRQuery(FHIRRequest{httpMethod: http.MethodGet, qid: id, url: fhirURL})

	if result.err != nil {
		c.JSON(http.StatusInternalServerError, Error{ErrorCode: "500", ErrorMessage: result.err.Error()})
		return
	}
	var data PractitionerFHIR
	err := json.Unmarshal(result.response, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if data.ID != id {
		errMsg := "aid '" + id + "' cannot be found"
		c.JSON(http.StatusNotFound, Error{ErrorCode: "404", ErrorMessage: errMsg})
		return
	}
	c.JSON(http.StatusOK, data.ToAPISpec())
}

// AuthorsAidPut - Add or update author information
func AuthorsAidPut(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: "Expected Content-Type: 'application/json', got '" + contentType + "'"})
		return
	}

	var data Author
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: err.Error()})
		return
	}

	id := c.Param("aid")
	if data.Aid != id {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: "aid in param and body do not match"})
		return
	}

	dataFhir := data.ToFHIR()
	jsonData, _ := json.Marshal(dataFhir)
	fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Practitioner/"+id)
	result := SingleFHIRQuery(FHIRRequest{httpMethod: http.MethodPut, qid: id, url: fhirURL, body: string(jsonData)})
	if result.err != nil {
		c.JSON(http.StatusInternalServerError, Error{ErrorCode: "500", ErrorMessage: result.err.Error()})
		return
	}
	c.JSON(result.statusCode, dataFhir)
}

// AuthorsGet - Mass query for author metadata in HoloStorage
func AuthorsGet(c *gin.Context) {
	fhirRequests := make(map[string]FHIRRequest)
	ids := ParseQueryIDs(c.Query("aid"))

	for _, id := range ids {
		fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Practitioner/"+id)
		fhirRequests[id] = FHIRRequest{httpMethod: http.MethodGet, qid: id, url: fhirURL}
	}

	results := BatchFHIRQuery(fhirRequests)

	authorsMap := make(map[string]Author)
	for id, result := range results {
		var tempAuthor PractitionerFHIR
		err := json.Unmarshal(result.response, &tempAuthor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if tempAuthor.ID != id {
			authorsMap[id] = Author{}
		} else {
			authorsMap[id] = tempAuthor.ToAPISpec()
		}
	}

	c.JSON(http.StatusOK, authorsMap)
}

// HologramsGet - Mass query for hologram metadata based on hologram ids
func HologramsGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HologramsHidDelete - Delete a hologram in HoloStorage
func HologramsHidDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HologramsHidDownloadGet - Download holograms models based on the hologram id
func HologramsHidDownloadGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HologramsHidGet - Get a single hologram metadata based on the hologram id
func HologramsHidGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HologramsPost - Upload hologram to HoloStorage
func HologramsPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PatientsGet - Mass query for patients metadata in HoloStorage
func PatientsGet(c *gin.Context) {
	fhirRequests := make(map[string]FHIRRequest)
	ids := ParseQueryIDs(c.Query("pid"))

	for _, id := range ids {
		fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Patient/"+id)
		fhirRequests[id] = FHIRRequest{httpMethod: http.MethodGet, qid: id, url: fhirURL}
	}

	results := BatchFHIRQuery(fhirRequests)

	patientsMap := make(map[string]Patient)
	var emptyData Patient
	for id, result := range results {
		var tempData PatientFHIR
		err := json.Unmarshal(result.response, &tempData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if tempData.ID != id {
			patientsMap[id] = emptyData
		} else {
			patientsMap[id] = tempData.ToAPISpec()
		}
	}

	c.JSON(http.StatusOK, patientsMap)
}

// PatientsPidGet - Get a single patient metadata in HoloStorage
func PatientsPidGet(c *gin.Context) {
	id := c.Param("pid")
	fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Patient/"+id)
	result := SingleFHIRQuery(FHIRRequest{httpMethod: http.MethodGet, qid: id, url: fhirURL})

	if result.err != nil {
		c.JSON(http.StatusInternalServerError, Error{ErrorCode: "500", ErrorMessage: result.err.Error()})
		return
	}
	var data PatientFHIR
	err := json.Unmarshal(result.response, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if data.ID != id {
		errMsg := "id '" + id + "' cannot be found"
		c.JSON(http.StatusNotFound, Error{ErrorCode: "404", ErrorMessage: errMsg})
		return
	}
	c.JSON(http.StatusOK, data.ToAPISpec())
}

// PatientsPidPut - Add or update basic patient information
func PatientsPidPut(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType != "application/json" {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: "Expected Content-Type: 'application/json', got '" + contentType + "'"})
		return
	}

	var data Patient
	id := c.Param("pid")

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: err.Error()})
		return
	}

	if data.Pid != id {
		c.JSON(http.StatusBadRequest, Error{ErrorCode: "400", ErrorMessage: "pid in param and body do not match"})
		return
	}

	dataFhir := data.ToFHIR()
	jsonData, _ := json.Marshal(dataFhir)
	fhirURL, _ := ConstructURL(accessorConfig.FhirURL, "Patient/"+id)
	result := SingleFHIRQuery(FHIRRequest{httpMethod: http.MethodPut, qid: id, url: fhirURL, body: string(jsonData)})
	if result.err != nil {
		c.JSON(http.StatusInternalServerError, Error{ErrorCode: "500", ErrorMessage: result.err.Error()})
		return
	}
	c.JSON(result.statusCode, dataFhir)
}
