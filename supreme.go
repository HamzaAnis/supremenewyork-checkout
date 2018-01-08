package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type AddProduct struct {
	Utf8   string `json:"utf8,omitempty"`
	Style  int    `json:"style"`
	Size   int    `json:"size"`
	Commit string `json:"commit,omitempty"`
	Qty    int    `json:"qty"`
}

type CheckoutInfo struct {
	Utf8              string `json:"utf8,omitempty"`
	AuthenticityToken string `json:"authenticity_token,omitempty"`
	BillingName       string `json:"order[billing_name]"`
	Email             string `json:"order[email]"`
	Tel               string `json:"order[tel]"`
	Address1          string `json:"order[billing_address]"`
	Address2          string `json:"order[billing_address_2]"`
	Address3          string `json:"order[billing_address_3]"`
	City              string `json:"order[billing_city]"`
	Zip               string `json:"order[billing_zip]"`
	Country           string `json:"order[billing_country]"`
	SameAsBilling     bool   `json:"same_as_billing_address"`
	StoreCreditID     string `json:"store_credit_id"`
	CardType          string `json:"credit_card[type]"`
	CNB               string `json:"credit_card[cnb]"`
	Month             string `json:"credit_card[month]"`
	Year              string `json:"credit_card[year]"`
	VVal              string `json:"credit_card[vval]"`
	Terms             bool   `json:"order[terms]"`
	Captcha           string `json:"g-recaptcha-response"`
}

func main() {

	//the product below is this http://www.supremenewyork.com/shop/accessories/jtn5gqd12/livrjqw6h
	//which is the same as http://www.supremenewyork.com/shop/302969/

	// URL, _ := url.Parse("http://www.supremenewyork.com")

	// cookieJar, _ := cookiejar.New(nil)

	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "hasShownCookieNotice",
		Value:  "1",
		Path:   "/",
		Domain: "supremenewyork.com",
	}
	cookies = append(cookies, cookie)

	form := url.Values{}
	form.Add("utf8", "✓")
	form.Add("style", "20429")
	form.Add("size", "42323")
	form.Add("commit", "add to basket")

	m := map[string]interface{}{
		"utf8":   "✓",
		"style":  "20429",
		"size":   "42323",
		"commit": "add to basket"}
	request := gorequest.New()
	_, body, errs := request.Post("http://www.supremenewyork.com/shop/302969/add.json").
		Param("utf8", "✓").
		Param("style", "20429").
		Param("size", "42323").
		Param("commit", "add to basket").
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Firefox/57.0").
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("Accept-Encoding", "gzip, deflate, br").
		Set("Referer", "https://www.supremenewyork.com").
		Set("Origin", "https://www.supremenewyork.com").
		Set("Connection", "keep-alive").
		Set("Accept", "application/json").
		Set("X-Requested-Width", "XMLHttpRequest").
		Send(m).
		AddCookies(cookies).
		End()
	// Set("X-CSRF-Token", token).
	if errs != nil {
		log.Println(errs)
		os.Exit(1)
	}
	// req, body, errs := request.Get("https://www.supremenewyork.com").End()
	// fmt.Println("REsp is " + resp.Status)
	// if errs != nil {
	// 	fmt.Print("errrs")
	// 	fmt.Print(errs)
	// }
	// resp.Body.Close()
	fmt.Print("Status of the item is " + body)

	_, body, err := request.Get("http://www.supremenewyork.com/checkout").End()

	if err != nil {
		log.Println(errs)
		os.Exit(1)
	}
	pageOutput := body
	i := strings.Index(pageOutput, "csrf-token")
	chars := pageOutput[i:]
	x := strings.Index(chars, " />")
	temp := chars[:x]
	temp = temp[21:]
	g := len(temp)
	temp = temp[:(g - 1)]
	token := temp
	fmt.Println("After getting the checkout page " + body)
	checkoutinfo := CheckoutInfo{
		AuthenticityToken: token,
		BillingName:       "",
		Email:             "", //string@domain.com
		Tel:               "",
		Address1:          "",
		City:              "",
		Zip:               "",
		Country:           "",
		SameAsBilling:     true,
		CardType:          "visa",
		CNB:               "", // 1111 1111 1111 1111
		Month:             "", // 01
		Year:              "", // 2020
		VVal:              "", //123
		Terms:             true,
		Captcha:           "",
	}
	m = map[string]interface{}{
		"utf8":                     "✓",
		"authenticity_token":       token,
		"order[billing_name]":      checkoutinfo.BillingName,
		"order[email]":             checkoutinfo.Email,
		"order[tel]":               checkoutinfo.Tel,
		"order[billing_address]":   checkoutinfo.Address1,
		"order[billing_address_2]": "",
		"order[billing_address_3]": "",
		"order[billing_city]":      checkoutinfo.City,
		"order[billing_zip]":       checkoutinfo.Zip,
		"order[billing_country]":   checkoutinfo.Country,
		"same_as_billing_address":  "1",
		"store_credit_id":          "",
		"credit_card[type]":        checkoutinfo.CardType,
		"credit_card[cnb]":         checkoutinfo.CNB,
		"credit_card[month]":       checkoutinfo.Month,
		"credit_card[year]":        checkoutinfo.Year,
		"credit_card[vval]":        checkoutinfo.VVal,
		"order[terms]":             "0",
		"hpcvv":                    "",
		"cnt":                      "1"}

	_, body1, err1 := request.Get("http://www.supremenewyork.com/checkout.js").
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Firefox/57.0").
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("Accept-Encoding", "gzip").
		Set("Referer", "https://www.supremenewyork.com/checkout").
		Set("Origin", "https://www.supremenewyork.com").
		Set("Connection", "keep-alive").
		Set("Accept", "application/json").
		Set("X-Requested-Width", "XMLHttpRequest").
		Send(m).
		AddCookies(cookies).
		End()

	if err1 != nil {
		log.Println(err1)
		os.Exit(1)
	}
	fmt.Print("\nAfter getting JS file " + string(body1))

}
