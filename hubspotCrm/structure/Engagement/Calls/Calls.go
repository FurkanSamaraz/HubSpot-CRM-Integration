package crm_structures

import "time"

type GetCalls struct {
}
type UpdateCalls struct {
	Properties struct {
		HsTimestamp        time.Time `json:"hs_timestamp"`
		HsCallTitle        string    `json:"hs_call_title"`
		HubspotOwnerID     string    `json:"hubspot_owner_id"`
		HsCallBody         string    `json:"hs_call_body"`
		HsCallDuration     string    `json:"hs_call_duration"`
		HsCallFromNumber   string    `json:"hs_call_from_number"`
		HsCallToNumber     string    `json:"hs_call_to_number"`
		HsCallRecordingURL string    `json:"hs_call_recording_url"`
		HsCallStatus       string    `json:"hs_call_status"`
	} `json:"properties"`
}
type CreateCalls struct {
	Properties struct {
		HsTimestamp        time.Time `json:"hs_timestamp"`
		HsCallTitle        string    `json:"hs_call_title"`
		HubspotOwnerID     string    `json:"hubspot_owner_id"`
		HsCallBody         string    `json:"hs_call_body"`
		HsCallDuration     string    `json:"hs_call_duration"`
		HsCallFromNumber   string    `json:"hs_call_from_number"`
		HsCallToNumber     string    `json:"hs_call_to_number"`
		HsCallRecordingURL string    `json:"hs_call_recording_url"`
		HsCallStatus       string    `json:"hs_call_status"`
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
