package crm_structures

// * https://api.hubspot.com/crm/v3/objects/contacts

type CreateContacts struct {
	Properties struct {
		Email          string `json:"email"`
		Firstname      string `json:"firstname"`
		Lastname       string `json:"lastname"`
		Phone          string `json:"phone"`
		Company        string `json:"company"`
		Website        string `json:"website"`
		Lifecyclestage string `json:"lifecyclestage"`
	} `json:"properties"`
}
type GetContacts struct {
}

// * https://api.hubspot.com/crm/v3/objects/contacts?properties=email&associations=companies.
// * crm/v3/objects/contacts/{contactId}
type UpdateContacts struct {
}
