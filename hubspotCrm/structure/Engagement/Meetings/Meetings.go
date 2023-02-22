package crm_structures

import "time"

type GetMeetings struct {
}
type CreateMeetings struct {
	Properties struct {
		HsTimestamp            time.Time `json:"hs_timestamp"`
		HubspotOwnerID         string    `json:"hubspot_owner_id"`
		HsMeetingTitle         string    `json:"hs_meeting_title"`
		HsMeetingBody          string    `json:"hs_meeting_body"`
		HsInternalMeetingNotes string    `json:"hs_internal_meeting_notes"`
		HsMeetingExternalURL   string    `json:"hs_meeting_external_url"`
		HsMeetingLocation      string    `json:"hs_meeting_location"`
		HsMeetingStartTime     time.Time `json:"hs_meeting_start_time"`
		HsMeetingEndTime       time.Time `json:"hs_meeting_end_time"`
		HsMeetingOutcome       string    `json:"hs_meeting_outcome"`
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
type UpdateMeetings struct {
	Properties struct {
		HsTimestamp            time.Time `json:"hs_timestamp"`
		HubspotOwnerID         string    `json:"hubspot_owner_id"`
		HsMeetingTitle         string    `json:"hs_meeting_title"`
		HsMeetingBody          string    `json:"hs_meeting_body"`
		HsInternalMeetingNotes string    `json:"hs_internal_meeting_notes"`
		HsMeetingExternalURL   string    `json:"hs_meeting_external_url"`
		HsMeetingLocation      string    `json:"hs_meeting_location"`
		HsMeetingStartTime     time.Time `json:"hs_meeting_start_time"`
		HsMeetingEndTime       time.Time `json:"hs_meeting_end_time"`
		HsMeetingOutcome       string    `json:"hs_meeting_outcome"`
	} `json:"properties"`
}
