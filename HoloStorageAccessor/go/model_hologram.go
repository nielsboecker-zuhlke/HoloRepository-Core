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
	"errors"
	"time"
)

// Hologram - Metadata of a hologram
type Hologram struct {
	// UUID of a hologram
	Hid string `json:"hid,omitempty"`

	// Summary of the hologram
	Title string `json:"title,omitempty"`

	// Further elaborations on the hologram with additional details
	Description string `json:"description,omitempty"`

	// The mimetype of the hologram file
	ContentType string `json:"contentType,omitempty"`

	// The size of the hologram in KB
	FileSizeInKb uint32 `json:"fileSizeInKb,omitempty"`

	// The body part the hologram represents
	BodySite string `json:"bodySite,omitempty"`

	// The date when the imaging was done
	DateOfImaging time.Time `json:"dateOfImaging,omitempty"`

	// The date the hologram was generated
	CreationDate time.Time `json:"creationDate,omitempty"`

	// The way the hologram was created
	CreationMode string `json:"creationMode,omitempty"`

	// The description associated with the creation method of the hologram
	CreationDescription string `json:"creationDescription,omitempty"`

	// The UUID of the author of this hologram
	Aid string `json:"aid,omitempty"`

	// The UUID of the patient associated with this hologram
	Pid string `json:"pid,omitempty"`
}

type ReferenceFHIR struct {
	Reference string `json:"reference,omitempty"`
}

type CodeableConceptFHIR struct {
	Text string `json:"text,omitempty"`
}

type ContentFHIR struct {
	Attachment AttachmentFHIR `json:"attachment,omitempty"`
}
type AttachmentFHIR struct {
	ContentType string `json:"contentType,omitempty"`
	URL         string `json:"url,omitempty"`
	Size        uint32 `json:"size,omitempty"`
	Title       string `json:"title,omitempty"`
}
type HologramDocumentReferenceFHIR struct {
	ResourceType string              `json:"resourceType,omitempty"`
	Id           string              `json:"id,omitempty"`
	Status       string              `json:"status,omitempty"`
	Date         time.Time           `json:"date,omitempty"`
	Type         CodeableConceptFHIR `json:"type,omitempty"`
	Content      []ContentFHIR       `json:"content,omitempty"`
	HologramMeta string              `json:"description,omitempty"`
	Subject      ReferenceFHIR       `json:"subject,omitempty"`
	Author       []ReferenceFHIR     `json:"author,omitempty"`
}

type HologramMeta struct {
	Description         string    `json:"description,omitempty"`
	CreationDescription string    `json:"creationDescription,omitempty"`
	BodySite            string    `json:"bodySite,omitempty"`
	DateOfImaging       time.Time `json:"dateOfImaging,omitempty"`
}

func (h Hologram) ToFHIR() HologramDocumentReferenceFHIR {
	hologramDocRef := HologramDocumentReferenceFHIR{ResourceType: "DocumentReference", Status: "Current"}
	hologramDocRef.Id = h.Hid
	hologramDocRef.Date = h.CreationDate
	hologramDocRef.Type = CodeableConceptFHIR{Text: h.CreationMode}

	// Process Attachment in ContentFHIR
	attachmentData := AttachmentFHIR{ContentType: h.ContentType, Size: h.FileSizeInKb * 1024, Title: h.Title}
	if attachmentData != (AttachmentFHIR{}) {
		contentData := ContentFHIR{Attachment: attachmentData}
		hologramDocRef.Content = append(hologramDocRef.Content, contentData)
	}
	// Process HologramMeta
	holoMeta := HologramMeta{
		Description:         h.Description,
		CreationDescription: h.CreationDescription,
		DateOfImaging:       h.DateOfImaging,
		BodySite:            h.BodySite,
	}
	if holoMeta != (HologramMeta{}) {
		holoMetadata, _ := json.Marshal(holoMeta)
		hologramDocRef.HologramMeta = string(holoMetadata)
	}

	return hologramDocRef
}

func (h *HologramDocumentReferenceFHIR) UpdateAttachmentURL(url string) error {
	if len(h.Content) > 0 {
		h.Content[0].Attachment.URL = url
		return nil
	} else {
		return errors.New("No content within DocumentReference")
	}
}
