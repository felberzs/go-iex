package goiex

import (
	"net/http"
	"net/url"
	"time"
)

// DataAPI struct to interface with DataAPI endpoints
type DataAPI struct {
	iex
}

// DataPoint struct
type DataPoint struct {
	Key         string    `json:"key"`
	Weight      int       `json:"weight"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// NewDataAPI return new DataAPI
func NewDataAPI(token, version string, base *url.URL, httpClient *http.Client) *DataAPI {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}
	return &DataAPI{
		iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// Token return token string
func (d *DataAPI) Token() string {
	return d.token
}

// Version return version string
func (d *DataAPI) Version() string {
	return d.version
}

// URL return URL base
func (d *DataAPI) URL() *url.URL {
	return d.url
}

// APIURL return APIURL
func (d *DataAPI) APIURL() *url.URL {
	return d.apiurl
}

// Client return HTTP client
func (d *DataAPI) Client() *http.Client {
	return d.client
}

// DataPoints GET /data-points/{symbol}
func (d *DataAPI) DataPoints(symbol string) (datapoints []*DataPoint, err error) {
	err = get(d, &datapoints, "data-points/"+symbol, nil)
	return
}
