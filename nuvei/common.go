package nuvei

import (
	"bytes"
	"encoding/json"
	"github.com/Peller-Tech-Inc/go-nuvei/nuvei/model"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

type nuveiProfile struct {
	merchantId     string
	merchantSiteId string
	secret         string
	server         string
	isDebug        bool
}

var lock = &sync.Mutex{}
var instance *nuveiProfile
var version = "1"

func Initialize(merchantId string, merchantSiteId string, secret string, isLive bool, isDebug bool) {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		var server string
		if isLive {
			server = "https://secure.safecharge.com/ppp/api/v" + version + "/"
		} else {
			server = "https://ppp-test.safecharge.com/ppp/api/v" + version + "/"
		}
		instance = &nuveiProfile{
			merchantId:     merchantId,
			merchantSiteId: merchantSiteId,
			secret:         secret,
			server:         server,
			isDebug:        isDebug,
		}
	}
}

func CreateRequest(body []byte, action model.NuveiAction, responseObject interface{}) (interface{}, error) {
	if instance.isDebug {
		log.Println("Creating request to " + string(action) + " with body: " + string(body))
	}
	url := instance.server + string(action)
	resp, err := request(http.MethodPost, url, body)
	if err != nil {
		log.Printf("Can't send request to Nuvei: %v", err)
		return nil, err
	}
	printResponse(resp)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Can't close request body after reading: %v", err)
		}
	}(resp.Body)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseObject)
	if err != nil {
		log.Printf("Can't read response body: %v", err)
		return nil, err
	}
	return responseObject, nil
}

func printResponse(resp *http.Response) {
	if instance.isDebug {
		res, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Printf("Response body: %s", res)
		}
	}
}

func request(method string, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(req)
}
