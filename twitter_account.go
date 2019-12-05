package anaconda

import (
	"database/sql"
)

type Settings struct {
	AlwaysUseHTTPS           bool            `json:"always_use_https"`
	DiscoverableByEmail      bool            `json:"discoverable_by_email"`
	GeoEnabled               bool            `json:"geo_enabled"`
	Language                 string          `json:"language"`
	Protected                bool            `json:"protected"`
	ScreenName               string          `json:"screen_name"`
	ShowAllInlineMedia       bool            `json:"show_all_inline_media"`
	SleepTime                SleepTime       `json:"sleep_time"`
	TimeZone                 TimeZone        `json:"time_zone"`
	Locations                []TrendLocation `json:"trend_location"`
	UseCookiePersonalization bool            `json:"use_cookie_personalization"`
	AllowContributorRequest  string          `json:"allow_contributor_request"`
}

type SleepTime struct {
	Enabled   bool         `json:"enabled"`
	EndTime   sql.NullTime `json:"end_time"`
	StartTime sql.NullTime `json:"start_time"`
}

type TimeZone struct {
	Name       string `json:"name"`
	TzInfoName string `json:"tzinfo_name"`
	UtcOffset  int64  `json:"utc_offset"`
}
