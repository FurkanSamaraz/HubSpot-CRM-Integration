package crm_structures

import "time"

type CreateProducts struct {
	Description              string `json:"description"`
	HsCostOfGoodsSold        string `json:"hs_cost_of_goods_sold"`
	HsRecurringBillingPeriod string `json:"hs_recurring_billing_period"`
	HsSku                    string `json:"hs_sku"`
	Name                     string `json:"name"`
	Price                    string `json:"price"`
}
type GetProducts struct {
	Results []struct {
		Properties struct {
			Createdate               time.Time `json:"createdate"`
			Description              string    `json:"description"`
			HsCostOfGoodsSold        string    `json:"hs_cost_of_goods_sold"`
			HsLastmodifieddate       time.Time `json:"hs_lastmodifieddate"`
			HsRecurringBillingPeriod string    `json:"hs_recurring_billing_period"`
			HsSku                    string    `json:"hs_sku"`
			Name                     string    `json:"name"`
			Price                    string    `json:"price"`
		} `json:"properties"`
	} `json:"results"`
	Paging struct {
		Next struct {
			After string `json:"after"`
			Link  string `json:"link"`
		} `json:"next"`
	} `json:"paging"`
}
type UpdateProducts struct {
}
