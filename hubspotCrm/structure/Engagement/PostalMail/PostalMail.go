package crm_repository

type CreatePostalMail struct {
	Properties struct {
		HsTimestamp      string `json:"hs_timestamp"`
		HsPostalMailBody string `json:"hs_postal_mail_body"`
		HubspotOwnerID   string `json:"hubspot_owner_id"`
		HsAttachmentIds  string `json:"hs_attachment_ids"`
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
type GetPostalMail struct {
}
type UpdatePostalMail struct {
	Properties struct {
		HsPostalMailBody string `json:"hs_postal_mail_body"`
	} `json:"properties"`
}
