package facebook

import (
	"fmt"
	"ftb/config"
	"net/http"
	"time"
)

const (
	accessToken = "EAAINSbsnPggBANzqj6okv2Np94bkSPNttSryVmAD7piJwyDXGPJNImt59V9pzdpzYbmRYwQ2vRyzekESdPMh9jgVF4qQKmPGcUrROcP2jdcjsz3CnhjRiEgvWCdSoeYE7mHLOwcZCrg0z7X5FohqgSaMoLP0be9aiQEC1mNZBsNSRcKRZAKUXeXGqhiZC1sZD"
	pageID      = "105705114157272"
)

func Init(config *config.Configuration) {
	client := &http.Client{Timeout: time.Duration(4) * time.Second}

	getPageAccessTokenUrl := "/" + pageID + "?fields=access_token&access_token=" + accessToken
	pageAccessToken := getPageAccessToken(*client, "GET", getPageAccessTokenUrl)

	getLeadFormUrl := "/" + pageID + "/leadgen_forms?access_token=" + pageAccessToken.PageAccessToken
	leadForm := getLeadFormData(*client, "GET", getLeadFormUrl)
	fmt.Println(leadForm)

	for _, lead := range leadForm.Data {
		getLeadsUrl := "/" + lead.ID + "/leads?access_token=" + pageAccessToken.PageAccessToken
		leadData := getLeadsData(*client, "GET", getLeadsUrl)
		fmt.Println(lead.Name)
		for _, item := range leadData.Data {
			fmt.Println(item.ID)
			fmt.Println(item.CreatedTime)
			for _, it := range item.FieldData {
				fmt.Println(it.Name)
				fmt.Println(it.Values)
			}
		}
	}

}
