package anaconda

import(
	"time"
)

type LineItems struct {
	Data []struct {
		BidType                              string           `json:"bid_type"`
		Name                                 string           `json:"name"`
		Placements                           []string         `json:"placements"`
		ProductType                          string           `json:"product_type"`
		BidAmountLocalMicro                  int64            `json:"bid_amount_local_micro"`
		AutomaticallySelectBid               bool             `json:"automatically_select_bid"`
		AdvertiserDomain                     string           `json:"advertiser_domain"`
		PrimaryWebEventTag                   string           `json:"primary_web_event_tag"`
		ChargeBy                             string           `json:"charge_by"`
		BidUnit                              string           `json:"bid_unit"`
		TotalBudgetAmountLocalMicro          int64            `json:"total_budget_amount_local_micro"`
		Objective                            string           `json:"objective"`
		Id                                   string           `json:"id"`
		Paused                               bool             `json:"paused"`
		AccountId                            string           `json:"account_id"`
		Optimization                         string           `json:"optimization"`
		Categories                           []string         `json:"categories"`
		Currency                             string           `json:"currency"`
		CreatedAt                            time.Time        `json:"created_at"`
		UpdatedAt                            time.Time        `json:"updated_at"`
		IncludeSentiment                     string           `json:"include_sentiment"`
		CampaignId                           string           `json:"campaign_id"`
		Deleted                              bool             `json:"deleted"`
	} `json:"data"`
	DataType                               string           `json:"data_type"`
	Request                                struct{
		Params                                  struct{
			WithDeleted bool       `json:"with_deleted"`
			AccountId   string     `json:"account_id"`
			CampaignIds []string   `json:"campaign_ids"`
		} `json:"params"`
	} `json:"request"`
	TotalCount                            int64             `json:"total_count"`
	NextCursor                            string            `json:"next_cursor"`
}