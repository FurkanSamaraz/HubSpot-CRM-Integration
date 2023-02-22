package crm_structures

import "time"

type GetDeals struct {
}

// * /crm/v3/objects/deals
type CreateDeals struct {
	Properties struct {
		Amount         string    `json:"amount"`
		Closedate      time.Time `json:"closedate"`
		Dealname       string    `json:"dealname"`
		Pipeline       string    `json:"pipeline"`
		Dealstage      string    `json:"dealstage"`
		HubspotOwnerID string    `json:"hubspot_owner_id"`
	} `json:"properties"`
}

// * /crm/v3/objects/deals/{dealId}
type UpdateDeals struct {
}
