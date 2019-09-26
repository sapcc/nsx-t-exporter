package exporter

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/time/rate"

	nsxv3config "github.com/sapcc/nsx-t-exporter/config"
)

const pathCreateSession = "/api/session/create"
const httpHeaderAcceptJSON = "application/json"
const httpHarderContentTypeJSON = "application/json"

// Nsxv3Client representes connection to NSXc3 Manger
type Nsxv3Client struct {
	config  nsxv3config.NSXv3Configuration
	client  http.Client
	cookie  string
	token   string
	limiter *rate.Limiter
	context context.Context
}

// Nsxv3Response http.Response wrapper including the error and extracted response content bytes
type Nsxv3Response struct {
	path     string
	response *http.Response
	body     []byte
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
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: c.SuppressSslWornings},
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

	c.cookie = resp.Header.Get("Set-Cookie")
	c.token = resp.Header.Get("X-XSRF-TOKEN")
	return nil
}

// AsyncGetRequest executes http get requests in an asych mode
func (c *Nsxv3Client) AsyncGetRequest(path string, ch chan<- *Nsxv3Response) error {
	c.login(false)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s%s", c.config.LoginHost, path), nil)
	if err != nil {
		ch <- &Nsxv3Response{path, nil, []byte{}, err}
	}

	req.Header.Set("Accept", httpHeaderAcceptJSON)
	req.Header.Set("Content-Type", httpHarderContentTypeJSON)
	req.Header.Set("Cookie", c.cookie)
	req.Header.Set("X-XSRF-TOKEN", c.token)

	resp, err := c.executeRequest(req)
	if err != nil {
		ch <- &Nsxv3Response{path, nil, []byte{}, err}
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- &Nsxv3Response{path, nil, []byte{}, err}
		return err
	}

	ch <- &Nsxv3Response{path, resp, bodyBytes, err}
	return nil
}

func (c *Nsxv3Client) executeRequest(req *http.Request) (*http.Response, error) {
	var cancel context.CancelFunc
	var childContext context.Context

	if c.limiter.Allow() == false {
		childContext, cancel = context.WithTimeout(c.context, 1*time.Second)
		c.limiter.Wait(childContext)
	}

	resp, err := c.client.Do(req)

	if err != nil && cancel != nil {
		cancel()
	}
	return resp, err
}
