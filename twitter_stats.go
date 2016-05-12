package anaconda

type Stats struct {
	Data                    []StatsModel        `json:"data"`
	TimeSeriesLength        int64               `json:"time_series_length"`
	Pagination
}

type StatsModel struct {
	Id                      string              `json:"id"`
	IdData                  []Metrics           `json:"id_data"`
}

type Metrics struct {
	AppClicks               string              `json:"app_clicks"`
	BilledCharge            []int64             `json:"billed_charge_local_micro"`
	BilledEngagements       []int64             `json:"billed_engagements"`
	CardEngagements         string              `json:"card_engagements"`
	CarouselSwipes          string              `json:"carousel_swipes"`
	Clicks                  []int64             `json:"clicks"`
	ConversionCustom        MetricType          `json:"conversion_custom"`
	ConversionDownloads     MetricType          `json:"conversion_downloads"`
	ConversionPurchases     MetricType          `json:"conversion_purchases"`
	ConversionSignUps       MetricType          `json:"conversion_sign_ups"`
	ConversionSiteVisits    MetricType          `json:"conversion_site_visits"`
	Engagements             []int64             `json:"engagements"`
	Follows                 []int64             `json:"follows"`
	Impressions             []int64             `json:"impressions"`
	Likes                   []int64             `json:"likes"`
	MediaViews              []int64             `json:"media_views"`
	Mobile                  MobileType
	QualifiedImpressions    string              `json:"qualified_impressions"`
	Replies                 []int64             `json:"replies"`
	Retweets                []int64             `json:"retweets"`
	Segment                 SegmentType         `json:"segment"`
	UrlClicks               []int64             `json:"url_clicks"`
	Video                   VideoType
}

type MetricType struct {
	Metric                  string              `json:"metric"`
	OrderQuantity           string              `json:"order_quantity"`
	SaleAmount              string              `json:"sale_amount"`
}

type MobileType struct {
	AchievementsUnlocked    MobileConversion    `json:"mobile_conversion_achievements_unlocked"`
	AddToCarts              MobileConversion    `json:"mobile_conversion_add_to_carts"`
	AddToWishlists          MobileConversion    `json:"mobile_conversion_add_to_wishlists"`
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
	Assisted                string              `json:"assisted"`
	OrderQuantity           string              `json:"order_quantity"`
	PostEngagement          string              `json:"post_engagement"`
	PostView                string              `json:"post_view"`
	SaleAmount              string              `json:"sale_amount"`
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