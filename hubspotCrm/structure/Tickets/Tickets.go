package crm_structures

type Create_Tickets struct {
	Properties struct {
		HsPipeline       string `json:"hs_pipeline"`
		HsPipelineStage  string `json:"hs_pipeline_stage"`
		HsTicketPriority string `json:"hs_ticket_priority"`
		Subject          string `json:"subject"`
	} `json:"properties"`
}
type Get_Tickets struct {
}

// * /crm/v3/objects/tickets/{ticketId}
type Update_Tickets struct {
}
