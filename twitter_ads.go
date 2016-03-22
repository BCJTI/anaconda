package anaconda

import "time"

const (
    AdsUrl = "https://ads-api.twitter.com/1"
)

type Account struct {
    Data                AccountModel    `json:"data"`
    DataType            string          `json:"data_type"`
}

type Accounts struct {
    Data                []AccountModel  `json:"data"`
    DataType            string          `json:"data_type"`
    NextCursor          string          `json:"next_cursor"`
    TotalCount          int64           `json:"total_count"`
}

type AccountModel struct {
    Id                  string          `json:"id"`
    Name                string          `json:"name"`
    ApprovalStatus      string          `json:"approval_status"`
    Salt                string          `json:"salt"`
    Deleted             bool            `json:"deleted"`
    Timezone            string          `json:"timezone"`
    TimezoneSwitchAt    time.Time       `json:"timezone_switch_at"`
    CreatedAt           time.Time       `json:"created_at"`
    UpdatedAt           time.Time       `json:"updated_at"`
}