package anaconda

import(
	"time"
)

type PromotedTweets struct {
	Data [] struct{
		TweetId                                   string                      `json:"tweet_id"`
		Cancelled                                 bool                        `json:"cancelled"`
		CreatedAt                                 time.Time                   `json:"created_at"`
		UpdatedAt                                 time.Time                   `json:"updated_at"`
		Deleted                                   bool                        `json:"deleted"`
		Id                                        string                      `json:"id"`
		Paused                                    bool                        `json:"paused"`
		ApprovalStatus                            string                      `json:"approval_status"`
		LineItemId                                string                      `json:"line_item_id"`
	} `json:"data"`
	DataType                                    string                      `json:"data_type"`
	Request                                     struct{
		Params                                       struct{
			AccountId   string `json:"account_id"`
			LineItemId  string `json:"line_item_id"`
			WithDeleted bool   `json:"with_deleted"`
		} `json:"params"`
	} `json:"request"`
	TotalCount                                  int64                       `json:"total_count"`
}