package anaconda

import (
	"time"
	"encoding/json"
)

type Stats struct {
	Data                    []StatsModel        `json:"data"`
	DataType                string              `json:"data_type"`
	NextCursor              string              `json:"next_cursor"`
	TotalCount              int64               `json:"total_count"`
	Request                 StatsRequest        `json:"request"`
}

type StatsModel struct {
	Id                      string              `json:"id"`
	IdData                  []Metrics           `json:"id_data"`
}

type StatsRequest struct {
	Params                  StatsParams         `json:"params"`
}

type StatsParams struct {
	Entity                  string              `json:"entity"`
	EntityIds               []json.Number       `json:"entity_ids"`
	Segmentation            string              `json:"segmentation_type"`
	Placement               string              `json:"placement"`
	Granularity             string              `json:"granularity"`
	Platform                string              `json:"platform"`
	Country                 string              `json:"country"`
	MetricGroups            []string            `json:"metric_groups"`
	StartTime               time.Time           `json:"start_time"`
	EndTime                 time.Time           `json:"end_time"`
}

type Metrics struct {
	Metric                  Metric              `json:"metrics"`
	Segment                 SegmentType         `json:"segment"`
}

type Metric struct {
	AppClicks               []int64             `json:"app_clicks"`
	BilledCharge            []int64             `json:"billed_charge_local_micro"`
	BilledEngagements       []int64             `json:"billed_engagements"`
	CardEngagements         []int64             `json:"card_engagements"`
	CarouselSwipes          []int64             `json:"carousel_swipes"`
	Clicks                  []int64             `json:"clicks"`
	Engagements             []int64             `json:"engagements"`
	Follows                 []int64             `json:"follows"`
	Impressions             []int64             `json:"impressions"`
	Likes                   []int64             `json:"likes"`
	MediaViews              []int64             `json:"media_views"`
	QualifiedImpressions    []int64             `json:"qualified_impressions"`
	Replies                 []int64             `json:"replies"`
	Retweets                []int64             `json:"retweets"`
	UrlClicks               []int64             `json:"url_clicks"`
	Mobile                  MobileType          `json:"-"`
	Video                   VideoType           `json:"-"`
	Web                     WebType             `json:"-"`
}

type WebType struct {
	ConversionCustom        ConversionType      `json:"conversion_custom"`
	ConversionDownloads     ConversionType      `json:"conversion_downloads"`
	ConversionPurchases     ConversionType      `json:"conversion_purchases"`
	ConversionSignUps       ConversionType      `json:"conversion_sign_ups"`
	ConversionSiteVisits    ConversionType      `json:"conversion_site_visits"`
}

type ConversionType struct {
	Metric                  []int64             `json:"metric"`
	OrderQuantity           []int64             `json:"order_quantity"`
	SaleAmount              []int64             `json:"sale_amount"`
}

type MobileType struct {
	AchievementsUnlocked    MobileConversion    `json:"mobile_conversion_achievements_unlocked"`
	AddToCarts              MobileConversion    `json:"mobile_conversion_add_to_carts"`
	AddToWishLists          MobileConversion    `json:"mobile_conversion_add_to_wishlists"`
	CheckoutsInitiated      MobileConversion    `json:"mobile_conversion_checkouts_initiated"`
	ContentViews            MobileConversion    `json:"mobile_conversion_content_views"`
	Installs                MobileConversion    `json:"mobile_conversion_installs"`
	Invites                 MobileConversion    `json:"mobile_conversion_invites"`
	KeyPageViews            MobileConversion    `json:"mobile_conversion_key_page_views"`
	LevelsAchieved          MobileConversion    `json:"mobile_conversion_levels_achieved"`
	Logins                  MobileConversion    `json:"mobile_conversion_logins"`
	PaymentInfoAdditions    MobileConversion    `json:"mobile_conversion_payment_info_additions"`
	Rates                   MobileConversion    `json:"mobile_conversion_rates"`
	ReEngages               MobileConversion    `json:"mobile_conversion_re_engages"`
	Reservations            MobileConversion    `json:"mobile_conversion_reservations"`
	Searches                MobileConversion    `json:"mobile_conversion_searches"`
	Shares                  MobileConversion    `json:"mobile_conversion_shares"`
	SpentCredits            MobileConversion    `json:"mobile_conversion_spent_credits"`
	TutorialsCompleted      MobileConversion    `json:"mobile_conversion_tutorials_completed"`
	Updates                 MobileConversion    `json:"mobile_conversion_updates"`
}

type MobileConversion struct {
	Assisted                []int64             `json:"assisted"`
	OrderQuantity           []int64             `json:"order_quantity"`
	PostEngagement          []int64             `json:"post_engagement"`
	PostView                []int64             `json:"post_view"`
	SaleAmount              []int64             `json:"sale_amount"`
}

type VideoType struct {
	ContentStarts           []int64             `json:"video_content_starts"`
	CtaClicks               []int64             `json:"video_cta_clicks"`
	MrcViews                []int64             `json:"video_mrc_views"`
	TotalViews              []int64             `json:"video_total_views"`
	Views25                 []int64             `json:"video_views_25"`
	Views50                 []int64             `json:"video_views_50"`
	Views75                 []int64             `json:"video_views_75"`
	Views100                []int64             `json:"video_views_100"`
}

type SegmentType struct{
	SegmentName             string              `json:"segment_name"`
	SegmentValue            string              `json:"segment_value"`
}