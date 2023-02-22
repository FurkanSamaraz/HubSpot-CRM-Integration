package crm_structures

import "time"

type CreateEmails struct {
	From struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"from"`
	To []struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"to"`
	Cc  []interface{} `json:"cc"`
	Bcc []interface{} `json:"bcc"`
}
type GetEmails struct {
}
type UpdateEmails struct {
	Properties struct {
		HsTimestamp      time.Time `json:"hs_timestamp"`
		HubspotOwnerID   string    `json:"hubspot_owner_id"`
		HsEmailDirection string    `json:"hs_email_direction"`
		HsEmailStatus    string    `json:"hs_email_status"`
		HsEmailSubject   string    `json:"hs_email_subject"`
		HsEmailText      string    `json:"hs_email_text"`
	} `json:"properties"`
}
