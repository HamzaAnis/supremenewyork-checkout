package main

import "github.com/parnurzeal/gorequest"
import "fmt"

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
	reqJsonUM := AddProduct{
		Style: 20429,
		Size:  42323,
		Qty:   1,
	}
	request := gorequest.New()
	req, body, errs := request.Get("https://www.supremenewyork.com").End()
	// fmt.Println("REsp is " + resp.Status)
	if errs != nil {
		fmt.Print("errrs")
		fmt.Print(errs)
	}
	// resp.Body.Close()
	fmt.Println("\n\n\n\n")
	fmt.Println(body)
}
