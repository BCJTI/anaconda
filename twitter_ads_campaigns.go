package anaconda

import(
	"time"
)

type Campaigns struct {
	Data [] struct {
		Name                                     string                 `json:"name"`
		CreatedAt                                time.Time              `json:"created_at"`
		EndTime                                  time.Time              `json:"end_time"`
		UpdatedAt                                time.Time              `json:"updated_at"`
		AccountId                                string                 `json:"account_id"`
		Deleted                                  bool                   `json:"deleted"`
		Id                                       string                 `json:"id"`
		Paused                                   bool                   `json:"paused"`
		TotalBudgetAmountLocalMicro              int64                  `json:"total_budget_amount_local_micro"`
		Currency                                 string                 `json:"currency"`
		DailyBudgetAmountLocalMicro              int64                  `json:"daily_budget_amount_local_micro"`
		FundingInstrumentId                      string                 `json:"funding_instrument_id"`
		StartTime                                time.Time              `json:"start_time"`
		StandardDelivery                         bool                   `json:"standard_delivery"`
	} `json:"data"`
	DataType                                    string                      `json:"data_type"`
	Request                                     struct{
		Params                                       struct{
			WithDeleted bool       `json:"with_deleted"`
			AccountId   string     `json:"account_id"`
		} `json:"params"`
	} `json:"request"`
	TotalCount                            int64             `json:"total_count"`

}
