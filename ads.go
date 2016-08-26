package anaconda

import (
	"fmt"
	"net/url"
)

func (a TwitterApi) GetAccounts(values url.Values) (accounts Accounts, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{AdsUrl + "/accounts", values, &accounts, _GET, response_ch}
	return accounts, (<-response_ch).err
}

func (a TwitterApi) GetAccountById(id string, values url.Values) (account Account, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s", AdsUrl, id), values, &account, _GET, response_ch}
	return account, (<-response_ch).err
}

func (a TwitterApi) GetAccountStats(id string, values url.Values) (stats Stats, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/stats/accounts/%s", AdsUrl, id), values, &stats, _GET, response_ch}
	return stats, (<-response_ch).err
}

func (a TwitterApi) GetCampaigns(accountid string, values url.Values) (campaigns Campaigns, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/campaigns", AdsUrl, accountid), values, &campaigns, _GET, response_ch}
	return campaigns, (<-response_ch).err
}

func (a TwitterApi) GetCampaignById(accountid, campaignid string, values url.Values) (campaign Campaign, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/campaigns/%s", AdsUrl, accountid, campaignid), values, &campaign, _GET, response_ch}
	return campaign, (<-response_ch).err
}

func (a TwitterApi) GetFundingInstruments(accountid string, values url.Values) (instruments FundingInstruments, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/funding_instruments", AdsUrl, accountid), values, &instruments, _GET, response_ch}
	return instruments, (<-response_ch).err
}

func (a TwitterApi) GetFundingInstrumentById(accountid, instrumentid string, values url.Values) (instrument FundingInstrument, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/funding_instruments/%s", AdsUrl, accountid, instrumentid), values, &instrument, _GET, response_ch}
	return instrument, (<-response_ch).err
}

func (a TwitterApi) GetLineItems(accountid string, values url.Values) (items LineItems, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/line_items", AdsUrl, accountid), values, &items, _GET, response_ch}
	return items, (<-response_ch).err
}

func (a TwitterApi) GetLineItemById(accountid, itemid string, values url.Values) (item LineItem, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/line_items/%s", AdsUrl, accountid, itemid), values, &item, _GET, response_ch}
	return item, (<-response_ch).err
}

func (a TwitterApi) GetPromotedTweets(accountid string, values url.Values) (tweets PromotedTweets, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/promoted_tweets", AdsUrl, accountid), values, &tweets, _GET, response_ch}
	return tweets, (<-response_ch).err
}

func (a TwitterApi) GetPromotedTweetById(accountid, tweetid string, values url.Values) (tweet PromotedTweet, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/promoted_tweets/%s", AdsUrl, accountid, tweetid), values, &tweet, _GET, response_ch}
	return tweet, (<-response_ch).err
}

func (a TwitterApi) GetTargetingCriteria(accountid string, values url.Values) (tc TargetingCriteria, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/targeting_criteria", AdsUrl, accountid), values, &tc, _GET, response_ch}
	return tc, (<-response_ch).err
}

func (a TwitterApi) GetAccountVideos(accountid string, values url.Values) (videos AccountVideos, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/videos", AdsUrl, accountid), values, &videos, _GET, response_ch}
	return videos, (<-response_ch).err
}

func (a TwitterApi) GetAccountVideoById(accountid, videoid string, values url.Values) (video AccountVideo, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/accounts/%s/videos/%s", AdsUrl, accountid, videoid), values, &video, _GET, response_ch}
	return video, (<-response_ch).err
}

func (a TwitterApi) GetAllCampaigns(accountid string, values url.Values) (campaigns Campaigns, err error) {

	// get first fetch
	if campaigns, err = a.GetCampaigns(accountid,values); err != nil {
		return
	}

	// init cursor variable
	cursor := campaigns.NextCursor

	// get all data
	for cursor != "" {

		// add cursor to query string
		values.Add("cursor",campaigns.NextCursor)

		// get next fetch
		next, e := a.GetCampaigns(accountid,values); if e != nil {
			return Campaigns{}, e
		}

		// append data
		campaigns.Data = append(campaigns.Data, next.Data...)

		// set pagination
		campaigns.Pagination = next.Pagination

		// set next cursor
		cursor = next.NextCursor

		// del previous cursor from query string
		values.Del("cursor")
	}

	return
}

func (a TwitterApi) GetAllFundingInstruments(accountid string, values url.Values) (instruments FundingInstruments, err error) {

	// get first fetch
	if instruments, err = a.GetFundingInstruments(accountid,values); err != nil {
		return
	}

	// init cursor variable
	cursor := instruments.NextCursor

	// get all data
	for cursor != "" {

		// add cursor to query string
		values.Add("cursor",instruments.NextCursor)

		// get next fetch
		next, e := a.GetFundingInstruments(accountid,values); if e != nil {
			return FundingInstruments{}, e
		}

		// append data
		instruments.Data = append(instruments.Data, next.Data...)

		// set pagination
		instruments.Pagination = next.Pagination

		// set next cursor
		cursor = next.NextCursor

		// del previous cursor from query string
		values.Del("cursor")
	}

	return
}

func (a TwitterApi) GetAllLineItems(accountid string, values url.Values) (items LineItems, err error) {

	// get first fetch
	if items, err = a.GetLineItems(accountid,values); err != nil {
		return
	}

	// init cursor variable
	cursor := items.NextCursor

	// get all data
	for cursor != "" {

		// add cursor to query string
		values.Add("cursor",items.NextCursor)

		// get next fetch
		next, e := a.GetLineItems(accountid,values); if e != nil {
			return LineItems{}, e
		}

		// append data
		items.Data = append(items.Data, next.Data...)

		// set pagination
		items.Pagination = next.Pagination

		// set next cursor
		cursor = next.NextCursor

		// del previous cursor from query string
		values.Del("cursor")
	}

	return
}

func (a TwitterApi) GetAllPromotedTweets(accountid string, values url.Values) (tweets PromotedTweets, err error) {

	// get first fetch
	if tweets, err = a.GetPromotedTweets(accountid,values); err != nil {
		return
	}

	// init cursor variable
	cursor := tweets.NextCursor

	// get all data
	for cursor != "" {

		// add cursor to query string
		values.Add("cursor",tweets.NextCursor)

		// get next fetch
		next, e := a.GetPromotedTweets(accountid,values); if e != nil {
			return PromotedTweets{}, e
		}

		// append data
		tweets.Data = append(tweets.Data, next.Data...)

		// set pagination
		tweets.Pagination = next.Pagination

		// set next cursor
		cursor = next.NextCursor

		// del previous cursor from query string
		values.Del("cursor")
	}

	return
}