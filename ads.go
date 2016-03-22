package anaconda

import (
    "net/url"
)

func (a TwitterApi) GetAccounts(values url.Values) (accounts Accounts, err error) {
    values = cleanValues(values)
    response_ch := make(chan response)
    a.queryQueue <- query{AdsUrl + "/accounts", values, &accounts, _GET, response_ch}
    return accounts, (<-response_ch).err
}

func (a TwitterApi) GetAccountById(id string, values url.Values) (account Account, err error) {
    values = cleanValues(values)
    response_ch := make(chan response)
    a.queryQueue <- query{AdsUrl + "/accounts/" + id, values, &account, _GET, response_ch}
    return account, (<-response_ch).err
}