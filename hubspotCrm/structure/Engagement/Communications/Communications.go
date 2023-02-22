package crm_repository

import "time"

type CreateCommunications struct {
	Properties struct {
		HsCommunicationChannelType string    `json:"hs_communication_channel_type"`
		HsCommunicationLoggedFrom  string    `json:"hs_communication_logged_from"`
		HsCommunicationBody        string    `json:"hs_communication_body"`
		HsTimestamp                time.Time `json:"hs_timestamp"`
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
type GetCommunications struct {
}
type UpdateCommunications struct {
	Properties struct {
		HsCommunicationBody string `json:"hs_communication_body"`
	} `json:"properties"`
}
