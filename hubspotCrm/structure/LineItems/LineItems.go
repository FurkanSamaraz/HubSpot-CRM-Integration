package crm_structures

//* POST request to https://api.hubapi.com/crm/v3/objects/line_item

type CreateLineItems struct {
	Properties struct {
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Name     string `json:"name"`
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
type GetLineItems struct {
}

//* PATCH request to https://api.hubapi.com/crm/v3/objects/line_item/{lineItemId}

type UpdateLineItems struct {
	Properties struct {
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Name     string `json:"name"`
	} `json:"properties"`
}
