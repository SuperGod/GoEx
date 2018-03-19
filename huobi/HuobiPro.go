package huobi

import (
	"net/http"
)

type HuobiPro struct {
	*HuoBi_V2
}

func NewHuobiProWithAddr(client *http.Client, apikey, secretkey, accountId, addr string) *HuobiPro {
	hbv2 := new(HuoBi_V2)
	hbv2.accountId = accountId
	hbv2.accessKey = apikey
	hbv2.secretKey = secretkey
	hbv2.httpClient = client
	hbv2.baseUrl = addr
	return &HuobiPro{hbv2}
}

func NewHuobiProMargin(client *http.Client, apikey, secretkey, subType string) (pro *HuobiPro, err error) {
	hbv2 := new(HuoBi_V2)
	hbv2.accessKey = apikey
	hbv2.secretKey = secretkey
	hbv2.httpClient = client
	hbv2.baseUrl = "https://api.huobipro.com"
	hbv2.accountId, err = hbv2.GetMarginAccountId(subType)
	if err != nil {
		return
	}
	hbv2.isMargin = true
	pro = &HuobiPro{hbv2}
	return
}

func NewHuobiPro(client *http.Client, apikey, secretkey string) (pro *HuobiPro, err error) {
	hbv2 := new(HuoBi_V2)
	hbv2.accessKey = apikey
	hbv2.secretKey = secretkey
	hbv2.httpClient = client
	hbv2.baseUrl = "https://api.huobipro.com"
	hbv2.accountId, err = hbv2.GetAccountId()
	if err != nil {
		return
	}
	pro = &HuobiPro{hbv2}
	return
}

func (hbpro *HuobiPro) GetExchangeName() string {
	return "huobi.pro"
}
