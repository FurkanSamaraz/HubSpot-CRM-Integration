package crm_structures

type CreateQuotes struct {
	Properties struct {
		HsTitle          string `json:"hs_title"`
		HsExpirationDate string `json:"hs_expiration_date"`
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
type GetQuotes struct {
}

//* PATCH request to https://api.hubapi.com/crm/v3/objects/quote/{QUOTE_ID}

type UpdateQuotes struct {
	Properties struct {
		HsStatus string `json:"hs_status"`
	} `json:"properties"`
}
