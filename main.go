package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
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

var channel = make(chan string)

func main() {

	//the product below is this http://www.supremenewyork.com/shop/accessories/jtn5gqd12/livrjqw6h
	//which is the same as http://www.supremenewyork.com/shop/302969/
	reqJsonUM := AddProduct{
		Style: 20429,
		Size:  42323,
		Qty:   1,
	}

	reqJson, err := json.Marshal(&reqJsonUM)
	println(string(reqJson))

	if err != nil {
		panic(err)
	}

	URL, _ := url.Parse("http://www.supremenewyork.com")

	cookieJar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "hasShownCookieNotice",
		Value:  "1",
		Path:   "/",
		Domain: "supremenewyork.com",
	}
	cookies = append(cookies, cookie)

	cookieJar.SetCookies(URL, cookies)
	httpClient := http.Client{Jar: cookieJar}

	form := url.Values{}
	form.Add("utf8", "✓")
	form.Add("style", "20429")
	form.Add("size", "42323")
	form.Add("commit", "add to basket")

	req, err := http.NewRequest(http.MethodPost, "http://www.supremenewyork.com/shop/302969/add.json", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Firefox/57.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println(res.Status)

	}
	out, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	println(string(out))

	println(len(cookieJar.Cookies(URL)), "COOKIES")

	for _, cookie := range cookieJar.Cookies(URL) {
		println(cookie.Name, " ", cookie.Value)
	}
	println()

	println()

	resp, err := http.Get("https://www.supremenewyork.com/checkout")
	if err != nil {
		os.Exit(1)
	}

	out, err = ioutil.ReadAll(resp.Body)

	pageOutput := string(out)
	i := strings.Index(pageOutput, "csrf-token")
	chars := pageOutput[i:]
	x := strings.Index(chars, " />")
	temp := chars[:x]
	temp = temp[21:]
	g := len(temp)
	temp = temp[:(g - 1)]
	token := temp

	println(string(out))

	println(token)

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
		Month:             "",               // 01
		Year:              "",             // 2020
		VVal:              "",              //123
		Terms:             true,
		Captcha:           "",
	}
	println("TOKEN IS "+token)

	form = url.Values{}
	form.Set("utf8", "✓")
	form.Set("authenticity_token", token)
	form.Set("order[billing_name]", checkoutinfo.BillingName)
	form.Set("order[email]", checkoutinfo.Email)
	form.Set("order[tel]", checkoutinfo.Tel)
	form.Set("order[billing_address]", checkoutinfo.Address1)
	form.Set("order[billing_address_2]", "")
	form.Set("order[billing_address_3]", "")
	form.Set("order[billing_city]", checkoutinfo.City)
	form.Set("order[billing_zip]", checkoutinfo.Zip)
	form.Set("order[billing_country]", checkoutinfo.Country)
	form.Set("same_as_billing_address", "1")
	form.Set("store_credit_id", "")
	form.Set("credit_card[type]", checkoutinfo.CardType)
	form.Set("credit_card[cnb]", checkoutinfo.CNB)
	form.Set("credit_card[month]", checkoutinfo.Month)
	form.Set("credit_card[year]", checkoutinfo.Year)
	form.Set("credit_card[vval]", checkoutinfo.VVal)
	form.Add("order[terms]", "0")
	form.Add("order[terms]", "1")
	form.Set("hpcvv", "")
	form.Set("cnt", "1")
	//REQUEST TO JS FILE

	req, err = http.NewRequest(http.MethodGet, "https://www.supremenewyork.com/checkout.js", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Firefox/57.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-CSRF-Token", token)
	req.Header.Set("Referer", "https://www.supremenewyork.com/checkout")
	req.Header.Set("Origin", "https://www.supremenewyork.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Requested-Width", "XMLHttpRequest")

	res, err = httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())

	}

	//ACTUAL CHECKOUT

	cform := url.Values{}
	cform.Set("utf8", "✓")
	cform.Set("authenticity_token", token)
	cform.Set("order[billing_name]", checkoutinfo.BillingName)
	cform.Set("order[email]", checkoutinfo.Email)
	cform.Set("order[tel]", checkoutinfo.Tel)
	cform.Set("order[billing_address]", checkoutinfo.Address1)
	cform.Set("order[billing_address_2]", "")
	cform.Set("order[billing_address_3]", "")
	cform.Set("order[billing_city]", checkoutinfo.City)
	cform.Set("order[billing_zip]", checkoutinfo.Zip)
	cform.Set("order[billing_country]", checkoutinfo.Country)
	cform.Set("same_as_billing_address", "1")
	cform.Set("store_credit_id", "")
	cform.Set("credit_card[type]", checkoutinfo.CardType)
	cform.Set("credit_card[cnb]", checkoutinfo.CNB)
	cform.Set("credit_card[month]", checkoutinfo.Month)
	cform.Set("credit_card[year]", checkoutinfo.Year)
	cform.Set("credit_card[vval]", checkoutinfo.VVal)
	cform.Add("order[terms]", "0")
	cform.Add("order[terms]", "1")
	cform.Set("hpcvv", "")
	cform.Set("g-recaptcha-response", checkoutinfo.Captcha)

	nreq, err := http.NewRequest(http.MethodPost, "https://www.supremenewyork.com/checkout.json", strings.NewReader(cform.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nreq.Header.Set("Accept-Encoding", "gzip, deflate, br")
	nreq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:57.0) Gecko/20100101 Firefox/57.0")
	nreq.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	nreq.Header.Set("X-CSRF-Token", token)
	nreq.Header.Set("Referer", "https://www.supremenewyork.com/checkout")
	nreq.Header.Set("Origin", "https://www.supremenewyork.com")


	for _, cookie := range cookieJar.Cookies(URL) {
		req.AddCookie(cookie)
	}

	time.Sleep(1 * time.Second)

	res, err = httpClient.Do(nreq)
	if err != nil {
		panic(err)

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println(res.Status)

	}
	out, err = ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	println("Response is" + string(out))

	return
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
