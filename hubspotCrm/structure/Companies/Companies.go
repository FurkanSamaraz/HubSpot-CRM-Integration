package crm_structures

type GetCompanies struct {
}

type CreateCompanies struct {
	Properties struct {
		Name           string `json:"name"`
		Domain         string `json:"domain"`
		City           string `json:"city"`
		Industry       string `json:"industry"`
		Phone          string `json:"phone"`
		State          string `json:"state"`
		Lifecyclestage string `json:"lifecyclestage"`
	} `json:"properties"`
}

type UpdateCompanies struct {
}
