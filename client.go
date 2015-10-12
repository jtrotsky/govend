// Package main handles interactions with the Vend API.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Client contains API authentication details.
type Client struct {
	Token        string
	DomainPrefix string
	TimeZone     string
}

// NewClient is called to pass authentication details to the manager.
func NewClient(token, domainPrefix, tz string) Client {
	return Client{token, domainPrefix, tz}
}

// Registers gets all registers from a store.
func (c Client) Registers() (*[]Register, error) {

	// Build the URL for the register page.
	url := urlFactory(0, c.DomainPrefix, "", "registers")

	body, err := urlGet(c.Token, url)
	if err != nil {
		fmt.Printf("Error getting resource: %s", err)
	}

	// Decode the JSON into our defined register object.
	response := RegisterPayload{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("\nError unmarshalling Vend sale payload: %s", err)
		return &[]Register{}, err
	}

	// Data is an array of register objects.
	data := response.Data

	// Do not expect more than one page of registers.
	// TODO: Consider including check for multiple pages.
	// version = response.Version["max"]

	return &data, err
}

// Users gets all users from a store.
func (c Client) Users() (*[]User, error) {

	// Build the URL for the register page.
	url := urlFactory(0, c.DomainPrefix, "", "users")

	body, err := urlGet(c.Token, url)
	if err != nil {
		fmt.Printf("Error getting resource: %s", err)
	}

	// Decode the JSON into our defined product object.
	response := UserPayload{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("\nError unmarshalling Vend sale payload: %s", err)
		return &[]User{}, err
	}

	// Data is an array of user objects.
	data := response.Data

	// Do not expect more than one page of users.
	// TODO: Consider including check for multiple pages.
	// version = response.Version["max"]

	return &data, err
}

// Customers grabs and collates all customers in pages of 10,000.
func (c Client) Customers() (*[]Customer, error) {

	customers := []Customer{}
	cp := []Customer{}
	var v int64

	// v is a version that is used to get customers by page.
	// Here we get the first page.
	data, v, err := resourcePage(0, c.DomainPrefix, c.Token, "customers")

	// Unmarshal payload into sales object.
	err = json.Unmarshal(data, &cp)

	customers = append(customers, cp...)

	for len(cp) > 0 {
		cp = []Customer{}

		// Continue grabbing pages until we receive an empty one.
		data, v, err = resourcePage(v, c.DomainPrefix, c.Token, "customers")
		if err != nil {
			return nil, err
		}

		// Unmarshal payload into customer object.
		err = json.Unmarshal(data, &cp)

		// Append customer page to list of customers.
		customers = append(customers, cp...)
	}

	return &customers, err
}

// Products grabs and collates all products in pages of 10,000.
func (c Client) Products() (*[]Product, error) {

	products := []Product{}
	p := []Product{}
	data := []byte{}
	var v int64

	// v is a version that is used to get products by page.
	// Here we get the first page.
	data, v, err := resourcePage(0, c.DomainPrefix, c.Token, "products")

	// Unmarshal payload into sales object.
	err = json.Unmarshal(data, &p)

	products = append(products, p...)

	for len(p) > 0 {
		p = []Product{}

		// Continue grabbing pages until we receive an empty one.
		data, v, err = resourcePage(v, c.DomainPrefix, c.Token, "products")
		if err != nil {
			return nil, err
		}

		// Unmarshal payload into product object.
		err = json.Unmarshal(data, &p)

		// Append page to list.
		products = append(products, p...)
	}

	return &products, err
}

// Sales grabs and collates all sales.
func (c Client) Sales() (*[]Sale, error) {

	var sales []Sale
	var s []Sale
	var v int64

	// v is a version that is used to objects by page.
	// Here we get the first page.
	data, v, err := resourcePage(0, c.DomainPrefix, c.Token, "sales")

	// Unmarshal payload into sales object.
	err = json.Unmarshal(data, &s)

	// Append page to list.
	sales = append(sales, s...)

	// NOTE: Turns out empty response is 2bytes.
	for len(data) > 2 {
		s = []Sale{}

		// Continue grabbing pages until we receive an empty one.
		data, v, err = resourcePage(v, c.DomainPrefix, c.Token, "sales")
		if err != nil {
			return nil, err
		}

		// Unmarshal payload into sales object.
		err = json.Unmarshal(data, &s)

		// Append sale page to list of sales.
		sales = append(sales, s...)
	}

	return &sales, err
}

func resourcePage(version int64, domainPrefix, key,
	resource string) ([]byte, int64, error) {

	// Build the URL for the resource page.
	url := urlFactory(version, domainPrefix, "", resource)

	body, err := urlGet(key, url)
	if err != nil {
		fmt.Printf("Error getting resource: %s", err)
	}

	// Decode the raw JSON.
	response := Payload{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("\nError unmarshalling payload: %s", err)
		return nil, 0, err
	}

	// Data is the resource body.
	data := response.Data

	// Version contains the maximum version number of the resources.
	version = response.Version["max"]

	return data, version, err
}

// urlGet performs a get request on a Vend API endpoint.
func urlGet(key, url string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("\nError creating http request: %s", err)
		return nil, err
	}

	// Using personal token authentication.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
	// TODO: Set user agent?

	log.Printf("Grabbing: %s\n", url)
	// Doing the request.
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nError performing request: %s", err)
		return nil, err
	}
	// Make sure response body is closed at end.
	defer res.Body.Close()

	// Check for invalid status codes.
	responseCheck(res.StatusCode)

	// Read what we got back.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("\nError while reading response body: %s\n", err)
		return nil, err
	}

	return body, err
}

// responseCheck checks the HTTP status codes of responses.
func responseCheck(statusCode int) {

	// Check HTTP response status codes.
	switch statusCode {
	case 200:
	// 	Response is bueno.
	case 401:
		fmt.Printf("\nAccess denied - check personal API token. Status: %d",
			statusCode)
		os.Exit(0)
	case 404:
		fmt.Printf("\nURL not found - check domain prefix. Status: %d",
			statusCode)
		os.Exit(0)
	case 429:
		fmt.Printf("\nRate limited by the Vend API :S Status: %d",
			statusCode)
	default:
		fmt.Printf("\nGot an unknown status code - Google it. Status: %d",
			statusCode)
		os.Exit(0)
	}
}

// urlFactory creates a Vend API 2.0 URL based on a resource.
func urlFactory(version int64, domainPrefix, objectID, resource string) string {
	// Page size is capped at ten thousand for all endpoints except sales
	// which it is capped at five hundred.
	const (
		pageSize = 10000
		deleted  = true
	)

	// Using 2.x Endpoint.
	address := fmt.Sprintf("https://%s.vendhq.com/api/2.0/", domainPrefix)
	query := url.Values{}
	query.Add("after", fmt.Sprintf("%d", version))
	query.Add("page_size", fmt.Sprintf("%d", pageSize))
	query.Add("deleted", fmt.Sprintf("%t", deleted))

	if objectID != "" {
		address += fmt.Sprintf("%s/%s/products?%s", resource, objectID, query.Encode())
	} else {
		address += fmt.Sprintf("%s?%s", resource, query.Encode())
	}

	return address
}

// ParseVendDT converts the default Vend timestamp string into a
// go Time.time value.
func ParseVendDT(dt, tz string) time.Time {

	// Load store's timezone as location.
	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("Error loading timezone as location: %s", err)
	}

	// Default Vend timedate layout.
	const longForm = "2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(longForm, dt)
	if err != nil {
		log.Fatalf("Error parsing time into deafult timestamp: %s", err)
	}

	// Time in retailer's timezone.
	dtWithTimezone := t.In(loc)

	return dtWithTimezone

	// Time string with timezone removed.
	// timeStr := timeLoc.String()[0:19]
}
