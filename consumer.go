package gosend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CalculatePrice(request CalculatePriceRequest) (response PriceAndDistanceEstimate, err error) {
	queryParam := "?" + "origin=" + request.OriginLatLong + "&destination=" + request.DestinationLatLong + "&paymentType=" + request.PaymentType
	url := Host + URLCalculatePrice + queryParam

	GSRequest, err := http.NewRequest(MethodGet, url, nil)
	if err != nil {
		log.Println("GoSend calculate price new request error : ", err)
		return
	}

	GSRequest.Header.Add("Client-ID", ClientID)
	GSRequest.Header.Add("Pass-Key", PassKey)
	GSResponse, err := http.DefaultClient.Do(GSRequest)
	if err != nil {
		log.Println("GoSend response error : ", err)
		return
	}

	defer GSResponse.Body.Close()
	if GSResponse.StatusCode != 200 {
		log.Println("GoSend HTTP Response error code : ", GSResponse.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(GSResponse.Body)
	if err != nil {
		log.Println("GoSend read response body error :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("GoSend unmarshal response body error :", err)
		return
	}
	return
}

func Booking(request BookingRequest) (response BookingResponse, err error) {
	url := Host + URLBooking

	marshalRequest, err := json.Marshal(request)
	if err != nil {
		log.Println("Request booking JSON marshal error : ", err)
		return
	}

	payload := strings.NewReader(string(marshalRequest))
	reqs, err := http.NewRequest(MethodPost, url, payload)
	if err != nil {
		log.Println("GoSend booking new request error : ", err)
		return
	}

	reqs.Header.Add("Client-ID", ClientID)
	reqs.Header.Add("Pass-Key", PassKey)
	reqs.Header.Add("Content-Type", "application/json")
	reqs.Header.Add("Accept", "*/*")
	resps, err := http.DefaultClient.Do(reqs)
	if err != nil {
		log.Println("GoSend response error : ", err)
		return
	}

	defer resps.Body.Close()

	if resps.StatusCode != 201 {
		log.Println("GoSend HTTP Response error code : ", resps.StatusCode)
		err = fmt.Errorf("Gosend HTTP Response error")
		return
	}

	body, err := ioutil.ReadAll(resps.Body)
	if err != nil {
		log.Println("GoSend read response body error :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("GoSend unmarshal response body error :", err)
		return
	}
	return
}

func BookingStatus(uniqueType, unique string) (response BookingStatusResponse, err error) {
	url := Host
	if uniqueType == TypeGoSendOrderNumber {
		url += URLTraceGosendOrderNo + unique
	} else if uniqueType == TypeStoreOrderID {
		url += URLTraceGosendStoreOrderID + unique
	} else {
		err = errors.New("Invalid type.")
		return
	}

	GSRequest, err := http.NewRequest(MethodGet, url, nil)
	if err != nil {
		log.Println("GoSend get booking status new request error : ", err)
		return
	}

	GSRequest.Header.Add("Client-ID", ClientID)
	GSRequest.Header.Add("Pass-Key", PassKey)
	GSResponse, err := http.DefaultClient.Do(GSRequest)
	if err != nil {
		log.Println("GoSend response error : ", err)
		return
	}

	defer GSResponse.Body.Close()
	if GSResponse.StatusCode != 200 {
		log.Println("GoSend HTTP Response error code : ", GSResponse.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(GSResponse.Body)
	if err != nil {
		log.Println("GoSend read response body error :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("GoSend unmarshal response body error :", err)
		return
	}
	return
}

func Cancel(request BookingCancelRequest) (response BookingCancelResponse, err error) {
	url := Host + URLCancel

	marshalRequest, err := json.Marshal(request)
	if err != nil {
		log.Println("Request booking JSON marshal error : ", err)
		return
	}

	payload := strings.NewReader(string(marshalRequest))
	GSRequest, err := http.NewRequest(MethodPut, url, payload)
	if err != nil {
		log.Println("GoSend booking new request error : ", err)
		return
	}

	GSRequest.Header.Add("Client-ID", ClientID)
	GSRequest.Header.Add("Pass-Key", PassKey)
	GSRequest.Header.Add("Content-Type", "application/json")
	GSRequest.Header.Add("Accept", "*/*")
	GSResponse, err := http.DefaultClient.Do(GSRequest)
	if err != nil {
		log.Println("GoSend response error : ", err)
		return
	}

	defer GSResponse.Body.Close()

	if GSResponse.StatusCode != 200 {
		log.Println("GoSend HTTP Response error code : ", GSResponse.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(GSResponse.Body)
	if err != nil {
		log.Println("GoSend read response body error :", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("GoSend unmarshal response body error :", err)
		return
	}
	return
}
