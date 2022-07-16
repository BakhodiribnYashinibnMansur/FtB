package facebook

import (
	"encoding/json"
	"fmt"
	"ftb/model"
	"net/http"
)

func getPageAccessToken(client http.Client, method, url string) model.PageAccessTokenModel {
	var model model.PageAccessTokenModel
	requestUrl := "https://graph.facebook.com/v14.0" + url

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}

	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}

	fmt.Printf("AccessToken ::: %s  \n ID:::%v", model.PageAccessToken, model.ID)
	fmt.Println()
	defer resp.Body.Close()
	fmt.Printf("Body : %s \n ", resp.Body)
	fmt.Printf("Response status : %s \n", resp.Status)
	return model
}

func getLeadFormData(client http.Client, method, url string) model.LeadFromModel {
	var model model.LeadFromModel
	requestUrl := "https://graph.facebook.com/v14.0" + url

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}

	defer resp.Body.Close()
	return model
}
func getLeadsData(client http.Client, method, url string) model.LeadModel {
	var model model.LeadModel
	requestUrl := "https://graph.facebook.com/v14.0" + url

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		fmt.Printf("Error %s", err)
		return model
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&model)
	if err != nil {
		fmt.Printf("Error::: %s \n", err)
	}

	defer resp.Body.Close()
	return model
}
