package anaconda

import (
	"encoding/json"
	"time"
	"gopkg.in/guregu/null.v3"
)

const (
    AdsUrl = "https://ads-api.twitter.com/1"
)

type Pagination struct {
	DataType                string                  `json:"data_type"`
	NextCursor              string                  `json:"next_cursor"`
	TotalCount              int64                   `json:"total_count"`
	Request                 interface{}             `json:"request"`
}

type Account struct {
	Data                    AccountModel            `json:"data"`
	Pagination
}

type Accounts struct {
	Data                    []AccountModel          `json:"data"`
	Pagination
}

type AccountModel struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	ApprovalStatus          string                  `json:"approval_status"`
	Salt                    string                  `json:"salt"`
	Deleted                 bool                    `json:"deleted"`
	Timezone                string                  `json:"timezone"`
	TimezoneSwitchAt        time.Time               `json:"timezone_switch_at"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type Campaign struct {
	Data                    CampaignModel           `json:"data"`
	Pagination
}

type Campaigns struct {
	Data                    []CampaignModel         `json:"data"`
	Pagination
}

type CampaignModel struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	AccountId               string                  `json:"account_id"`
	FundingInstrumentId     string                  `json:"funding_instrument_id"`
	Currency                string                  `json:"currency"`
	Paused                  bool                    `json:"paused"`
	Deleted                 bool                    `json:"deleted"`
	StandardDelivery        bool                    `json:"standard_delivery"`
	TotalBudget             int64                   `json:"total_budget_amount_local_micro"`
	DailyBudget             int64                   `json:"daily_budget_amount_local_micro"`
	StartTime               time.Time               `json:"start_time"`
	EndTime                 null.Time               `json:"end_time"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type FundingInstrument struct {
	Data                    FundingModel            `json:"data"`
	DataType                string                  `json:"data_type"`
	Request                 interface{}             `json:"request"`
}

type FundingInstruments struct {
	Data                    []FundingModel          `json:"data"`
	Pagination
}

type FundingModel struct {
	Id                      string                  `json:"id"`
	AccountId               string                  `json:"account_id"`
	Currency                string                  `json:"currency"`
	Type                    string                  `json:"type"`
	Description             string                  `json:"description"`
	CreditLimit             int64                   `json:"credit_limit_local_micro"`
	FundedAmount            int64                   `json:"funded_amount_local_micro"`
	Deleted                 bool                    `json:"deleted"`
	Cancelled               bool                    `json:"cancelled"`
	StartTime               time.Time               `json:"start_time"`
	EndTime                 time.Time               `json:"end_time"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type LineItem struct {
	Data                    LineItemModel           `json:"data"`
	Pagination
}

type LineItems struct {
	Data                    []LineItemModel         `json:"data"`
	Pagination
}

type LineItemModel struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	AccountId               string                  `json:"account_id"`
	CampaignId              string                  `json:"campaign_id"`
	BidType                 string                  `json:"bid_type"`
	Placements              []string                `json:"placements"`
	ProductType             string                  `json:"product_type"`
	BidAmount               int64                   `json:"bid_amount_local_micro"`
	AutomaticallySelectBid  bool                    `json:"automatically_select_bid"`
	AdvertiserDomain        string                  `json:"advertiser_domain"`
	PrimaryWebEventTag      string                  `json:"primary_web_event_tag"`
	ChargeBy                string                  `json:"charge_by"`
	BidUnit                 string                  `json:"bid_unit"`
	TotalBudget             int64                   `json:"total_budget_amount_local_micro"`
	Objective               string                  `json:"objective"`
	Paused                  bool                    `json:"paused"`
	Deleted                 bool                    `json:"deleted"`
	Optimization            string                  `json:"optimization"`
	Categories              []string                `json:"categories"`
	Currency                string                  `json:"currency"`
	IncludeSentiment        string                  `json:"include_sentiment"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type PromotedTweet struct {
	Data                    PromotedTweetModel      `json:"data"`
	Pagination
}

type PromotedTweets struct {
	Data                    []PromotedTweetModel    `json:"data"`
	Pagination
}

type PromotedTweetModel struct {
	Id                      string                  `json:"id"`
	TweetId                 string                  `json:"tweet_id"`
	LineItemId              string                  `json:"line_item_id"`
	Paused                  bool                    `json:"paused"`
	Deleted                 bool                    `json:"deleted"`
	Cancelled               bool                    `json:"cancelled"`
	ApprovalStatus          string                  `json:"approval_status"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type AccountVideo struct {
	Data                    VideoModel              `json:"data"`
	Pagination
}

type AccountVideos struct {
	Data                    []VideoModel            `json:"data"`
	Pagination
}

type VideoModel struct {
	Id                      string                  `json:"id"`
	Title                   string                  `json:"title"`
	Tweeted                 bool                    `json:"tweeted"`
	ReadyToTweet            bool                    `json:"ready_to_tweet"`
	Deleted                 bool                    `json:"deleted"`
	Duration                int64                   `json:"duration"`
	Description             string                  `json:"description"`
	PreviewUrl              string                  `json:"preview_url"`
	ReasonsNotServable      []string                `json:"reasons_not_servable"`
	CreatedAt               time.Time               `json:"created_at"`
	UpdatedAt               time.Time               `json:"updated_at"`
}

type Data struct {
	Id                      int64                   `json:"id"`
	AccountId               string                  `json:"account_id"`
	URL                     string                  `json:"url"`
	Entity                  string                  `json:"entity"`
	EntityIds               []json.Number           `json:"entity_ids"`
	Placement               string                  `json:"placement"`
	Granularity             string                  `json:"granularity"`
	SegmentationType        string                  `json:"segmentation_type"`
	Country                 string                  `json:"country"`
	Platform                string                  `json:"platform"`
	Status                  string                  `json:"status"`
	MetricGroups            []string                `json:"metric_groups"`
	StartTime               time.Time               `json:"start_time"`
	EndTime                 time.Time               `json:"end_time"`
}

type Job struct {
	Data                    Data                    `json:"data"`
	Pagination
}

type Jobs struct {
	Data                    []Data                  `json:"data"`
	Pagination
}