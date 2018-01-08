package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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
		Set("Referer", "https://www.supremenewyork.com/checkout").
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
	}
	// req, body, errs := request.Get("https://www.supremenewyork.com").End()
	// fmt.Println("REsp is " + resp.Status)
	// if errs != nil {
	// 	fmt.Print("errrs")
	// 	fmt.Print(errs)
	// }
	// resp.Body.Close()
	fmt.Print(body)
}
