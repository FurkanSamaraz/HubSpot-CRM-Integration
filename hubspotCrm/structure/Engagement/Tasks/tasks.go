package crm_structures

import "time"

type CreateTasks struct {
	Properties struct {
		HsTimestamp    time.Time `json:"hs_timestamp"`
		HsTaskBody     string    `json:"hs_task_body"`
		HubspotOwnerID string    `json:"hubspot_owner_id"`
		HsTaskSubject  string    `json:"hs_task_subject"`
		HsTaskStatus   string    `json:"hs_task_status"`
		HsTaskPriority string    `json:"hs_task_priority"`
		HsTaskType     string    `json:"hs_task_type"`
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
type GetTasks struct {
}
type UpdateTasks struct {
	Properties struct {
		HsTimestamp    time.Time `json:"hs_timestamp"`
		HsTaskBody     string    `json:"hs_task_body"`
		HubspotOwnerID string    `json:"hubspot_owner_id"`
		HsTaskSubject  string    `json:"hs_task_subject"`
		HsTaskStatus   string    `json:"hs_task_status"`
		HsTaskPriority string    `json:"hs_task_priority"`
	} `json:"properties"`
}
