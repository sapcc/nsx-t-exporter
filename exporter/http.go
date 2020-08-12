package exporter

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/time/rate"

	nsxv3config "github.com/sapcc/nsx-t-exporter/config"
	log "github.com/sirupsen/logrus"
)

const pathCreateSession = "/api/session/create"
const httpHeaderAcceptJSON = "application/json"
const httpHarderContentTypeJSON = "application/json"

// Nsxv3Client represents connection to NSXc3 Manger
type Nsxv3Client struct {
	config  nsxv3config.NSXv3Configuration
	client  http.Client
	cookie  string
	token   string
	limiter *rate.Limiter
	context context.Context
}

// Nsxv3ResourceKind represents Nsxv3Resource type
type Nsxv3ResourceKind string

// Nsxv3Resource resource kinds
const (
	ManagementCluster           Nsxv3ResourceKind = "ManagementCluster"
	ManagementClusterNodes      Nsxv3ResourceKind = "ManagementClusterNodes"
	ManagementClusterDatabase   Nsxv3ResourceKind = "ManagementClusterDatabase"
	ManagerNodeFirewall         Nsxv3ResourceKind = "ManagerNodeFirewall"
	ManagerNodeFirewallSections Nsxv3ResourceKind = "ManagerNodeFirewallSections"
	TransportNode               Nsxv3ResourceKind = "TransportNode"
	TransportNodes              Nsxv3ResourceKind = "TransportNodes"
	LogicalSwitchAdmin          Nsxv3ResourceKind = "LogicalSwitchAdmin"
	LogicalSwitch               Nsxv3ResourceKind = "LogicalSwitch"
	LogicalPort                 Nsxv3ResourceKind = "LogicalPort"
	ActivityFrameworkStatistics Nsxv3ResourceKind = "ActivityFrameworkStatistics"
)

// Nsxv3Resource represents endpoint status snapshot
type Nsxv3Resource struct {
	request  *http.Request
	response *http.Response
	kind     Nsxv3ResourceKind
	state    map[string]interface{} // JSON
	err      error
}

// GetClient initialize NSXv3 http client
func GetClient(c nsxv3config.NSXv3Configuration) Nsxv3Client {
	timeout := time.Duration(c.RequestTimeout) * time.Second

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).Dial,
		TLSHandshakeTimeout: timeout,
		IdleConnTimeout:     timeout,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: c.SuppressSslWarnings},
		MaxIdleConns:        c.RequestsConnPoolSize,
		MaxIdleConnsPerHost: c.RequestsConnPoolSize,
	}

	return Nsxv3Client{
		config: c,
		client: http.Client{
			Timeout:   timeout,
			Transport: netTransport,
		},
		limiter: rate.NewLimiter(rate.Limit(c.RequestsPerSecond), 1),
		context: context.Background(),
	}
}

// login to NSXv3 manager
func (c *Nsxv3Client) login(force bool) error {

	if !force && (c.cookie != "" || c.token != "") {
		return nil
	}

	requestBody := url.Values{}
	requestBody.Set("j_username", c.config.LoginUser)
	requestBody.Set("j_password", c.config.LoginPassword)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://%s%s", c.config.LoginHost, pathCreateSession),
		strings.NewReader(requestBody.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.executeRequest(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.cookie = resp.Header.Get("Set-Cookie")
	c.token = resp.Header.Get("X-XSRF-TOKEN")
	return nil
}

// AsyncGetRequest executes http get requests in an async mode
func (c *Nsxv3Client) updateEndpointStatus(status *Nsxv3Resource) {
	c.login(false)

	if status.request.URL.Host == "" {
		status.request.URL.Host = c.config.LoginHost
	}
	status.request.URL.Scheme = "https"
	status.request.Header = http.Header{}
	status.request.Header.Set("Accept", httpHeaderAcceptJSON)
	status.request.Header.Set("Content-Type", httpHarderContentTypeJSON)
	status.request.Header.Set("Cookie", c.cookie)
	status.request.Header.Set("X-XSRF-TOKEN", c.token)

	status.response, status.err = c.executeRequest(status.request)

	if status.err != nil {
		return
	}

	if !(status.response.StatusCode >= 200 && status.response.StatusCode <= 299) {
		status.err = fmt.Errorf("Request to endpoint %v failed with %d", status.kind, status.response.StatusCode)
		return
	}

	defer status.response.Body.Close()

	var body []byte

	body, status.err = ioutil.ReadAll(status.response.Body)

	if status.err != nil {
		return
	}

	status.err = json.Unmarshal(body, &status.state)
}

func (c *Nsxv3Client) executeRequest(req *http.Request) (*http.Response, error) {
	var cancel context.CancelFunc
	var childContext context.Context

	if c.limiter.Allow() == false {
		childContext, cancel = context.WithTimeout(
			c.context,
			time.Duration(c.config.RequestsPerSecondTimeout)*time.Second)

		err := c.limiter.Wait(childContext)
		if err != nil && cancel != nil {
			log.Error(err)
			cancel()
			return nil, err
		}
	}

	resp, err := c.client.Do(req)

	if err != nil && cancel != nil {
		cancel()
	}
	return resp, err
}
