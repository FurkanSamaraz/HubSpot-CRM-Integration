package crm_repository

import "time"

type CreateNotes struct {
	Properties struct {
		HsTimestamp     time.Time `json:"hs_timestamp"`
		HsNoteBody      string    `json:"hs_note_body"`
		HubspotOwnerID  string    `json:"hubspot_owner_id"`
		HsAttachmentIds string    `json:"hs_attachment_ids"`
	} `json:"properties"`
	Associations []struct {
		To struct {
			ID int `json:"id"`
		} `json:"to"`
		Types []struct {
			AssociationCategory string `json:"associationCategory"`
			AssociationTypeID   int    `json:"associationTypeId"`
		} `json:"types"`
	} `json:"associations"`
}
type GetNotes struct {
}
type UpdateNotes struct {
	Properties struct {
		HsNoteBody      string `json:"hs_note_body"`
		HsAttachmentIds string `json:"hs_attachment_ids"`
	} `json:"properties"`
}
