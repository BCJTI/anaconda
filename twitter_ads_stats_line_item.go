package anaconda

import(
	"time"
)

type StatsLineItems struct {
	Data [] struct{
		EndTime                                                       time.Time                `json:"end_time"`
		Id                                                            string                   `json:"id"`
		promoted_tweet_profile_retweets                               []int64                  `json:"promoted_tweet_profile_retweets"`
		promoted_tweet_timeline_card_engagements                      []int64                  `json:"promoted_tweet_timeline_card_engagements"`
		promoted_account_follow_rate                                  []int64                  `json:"promoted_account_follow_rate"`
		promoted_video_views_100                                      []int64                  `json:"promoted_video_views_100"`
		promoted_tweet_search_url_clicks                              []int64                  `json:"promoted_tweet_search_url_clicks"`
		promoted_tweet_search_clicks                                  []int64                  `json:"promoted_tweet_search_clicks"`
		promoted_video_views_75                                       []int64                  `json:"promoted_video_views_75"`
		promoted_tweet_search_favorites                               []int64                  `json:"promoted_tweet_search_favorites"`
		billed_charge_local_micro                                     []int64                  `json:"billed_charge_local_micro"`
		promoted_tweet_search_retweets                                []int64                  `json:"promoted_tweet_search_retweets"`
		promoted_tweet_timeline_media_views                           []int64                  `json:"promoted_tweet_timeline_media_views"`
		promoted_tweet_profile_follows                                []int64                  `json:"promoted_tweet_profile_follows"`
		promoted_tweet_timeline_favorites                             []int64                  `json:"promoted_tweet_timeline_favorites"`
		promoted_video_total_views                                    []int64                  `json:"promoted_video_total_views"`
		promoted_tweet_app_open_attempts                              []int64                  `json:"promoted_tweet_app_open_attempts"`
		promoted_tweet_profile_impressions                            []int64                  `json:"promoted_tweet_profile_impressions"`
		promoted_tweet_search_card_engagements                        []int64                  `json:"promoted_tweet_search_card_engagements"`
		promoted_tweet_profile_favorites                              []int64                  `json:"promoted_tweet_profile_favorites"`
		promoted_account_impressions                                  []int64                  `json:"promoted_account_impressions"`
		promoted_tweet_search_replies                                 []int64                  `json:"promoted_tweet_search_replies"`
		promoted_tweet_search_follows                                 []int64                  `json:"promoted_tweet_search_follows"`
		promoted_video_views_50                                       []int64                  `json:"promoted_video_views_50"`
		promoted_tweet_profile_media_views                            []int64                  `json:"promoted_tweet_profile_media_views"`
		promoted_tweet_search_qualified_impressions                   []int64                  `json:"promoted_tweet_search_qualified_impressions"`
		promoted_video_views_25                                       []int64                  `json:"promoted_video_views_25"`
		promoted_tweet_timeline_clicks                                []int64                  `json:"promoted_tweet_timeline_clicks"`
		promoted_tweet_timeline_impressions                           []int64                  `json:"promoted_tweet_timeline_impressions"`
		promoted_account_follows                                      []int64                  `json:"promoted_account_follows"`
		promoted_tweet_timeline_url_clicks                            []int64                  `json:"promoted_tweet_timeline_url_clicks"`
		promoted_tweet_timeline_qualified_impressions                 []int64                  `json:"promoted_tweet_timeline_qualified_impressions"`
		billed_follows                                                []int64                  `json:"billed_follows"`
		promoted_tweet_search_media_views                             []int64                  `json:"promoted_tweet_search_media_views"`
		promoted_tweet_timeline_retweets                              []int64                  `json:"promoted_tweet_timeline_retweets"`
		promoted_tweet_profile_url_clicks                             []int64                  `json:"promoted_tweet_profile_url_clicks"`
		promoted_tweet_timeline_replies                               []int64                  `json:"promoted_tweet_timeline_replies"`
		promoted_tweet_app_install_attempts                           []int64                  `json:"promoted_tweet_app_install_attempts"`
		promoted_tweet_profile_replies                                []int64                  `json:"promoted_tweet_profile_replies"`
		billed_engagements                                            []int64                  `json:"billed_engagements"`
		promoted_tweet_profile_card_engagements                       []int64                  `json:"promoted_tweet_profile_card_engagements"`
		promoted_tweet_timeline_follows                               []int64                  `json:"promoted_tweet_timeline_follows"`
		promoted_tweet_profile_clicks                                 []int64                  `json:"promoted_tweet_profile_clicks"`
		promoted_tweet_search_engagements                             []int64                  `json:"promoted_tweet_search_engagements"`
		promoted_tweet_timeline_engagements                           []int64                  `json:"promoted_tweet_timeline_engagements"`
		promoted_tweet_search_impressions                             []int64                  `json:"promoted_tweet_search_impressions"`
		promoted_tweet_profile_qualified_impressions                  []int64                  `json:"promoted_tweet_profile_qualified_impressions"`
		promoted_tweet_profile_engagements                            []int64                  `json:"promoted_tweet_profile_engagements"`
		segment                                                       struct{
			segmentation_type       string     `json:"segmentation_type"`
			name                    string     `json:"name"`
			segmentation_value      string     `json:"segmentation_value"`
		} `json:"segment"`
		StartTime                                                     time.Time                `json:"start_time"`
		Granularity                                                   string                   `json:"granularity"`
		UpdatedAt                                                     time.Time                `json:"updated_at"`
	} `json:"data"`
	DataType                                                        string                   `json:"data_type"`
	Request                                                         struct{
		Params struct{
			Granularity    string     `json:"granularity"`
			AccountId      string     `json:"account_id"`
			Metrics        []string   `json:"metrics"`
			LineItemIds    []string   `json:"line_item_ids"`
			StartTime      time.Time  `json:"start_time"`
			EndTime        time.Time  `json:"end_time"`
		} `json:"params"`
	} `json:"request"`
	TotalCount                                                      int64                   `json:"total_count"`
}