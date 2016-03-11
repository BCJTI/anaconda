package anaconda

import(
	"time"
)

type FundingInstrument struct {
	Data [] struct{
		AccountId                                 string                      `json:"account_id"`
		Cancelled                                 bool                        `json:"cancelled"`
		CreatedAt                                 time.Time                   `json:"created_at"`
		CreditLimitLocalMicro                     int64                       `json:"credit_limit_local_micro"`
		CreditRemainingLocalMicro                 int64                       `json:"credit_remaining_local_micro"`
		Currency                                  string                      `json:"currency"`
		Deleted                                   bool                        `json:"deleted"`
		Description                               string                      `json:"description"`
		EndTime                                   time.Time                   `json:"end_time"`
		FundedAmountLocalMicro                    int64                       `json:"funded_amount_local_micro"`
		Id                                        string                      `json:"id"`
		StartTime                                 time.Time                   `json:"start_time"`
		Type                                      string                      `json:"type"`
		UpdatedAt                                 time.Time                   `json:"updated_at"`
	} `json:"data"`
	DataType                                    string                      `json:"data_type"`
	Request                                     struct{
		Params                                       struct{
			AccountId   string     `json:"account_id"`
			FundingInstrumentIds []string   `json:"funding_instrument_ids"`
			WithDeleted bool       `json:"with_deleted"`
		} `json:"params"`
	} `json:"request"`
	TotalCount                                   int64                      `json:"total_count"`
}