package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetJsonData(url string, v interface{}, headers map[string]string) (err error) {
	log := "url: " + url
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
	if err != nil {
		return errors.New(log + " action: RequestNew err: " + err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := http.DefaultClient
	client.Timeout = time.Duration(time.Second * 60)

	timeStart := time.Now()
	resp, err := client.Do(req)
	log += " time: " + strconv.Itoa(int(time.Since(timeStart)/time.Millisecond)) + "ms"
	if err != nil {
		return errors.New(log + " action: RequestDo err: " + err.Error())
	}
	defer resp.Body.Close()

	answerStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(log + " action: ResponseRead err: " + err.Error())
	}

	err = json.Unmarshal(answerStr, &v)
	if err != nil {
		return errors.New(log + " action: ResponseUnmarshal statusCode: " + strconv.Itoa(resp.StatusCode) + " body: " +
			strings.Replace(string(answerStr), "\n", "\\n", -1) + " err: " + err.Error())
	}

	return nil
}

func PostJsonData(url string, request interface{}, v interface{}, headers map[string]string) (err error) {
	log := "url: " + url
	requestByte, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestByte))
	if err != nil {
		return errors.New(log + " action: RequestNew err: " + err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := http.DefaultClient
	client.Timeout = time.Duration(time.Second * 60)

	timeStart := time.Now()
	resp, err := client.Do(req)
	log += " time: " + strconv.Itoa(int(time.Since(timeStart)/time.Millisecond)) + "ms"
	if err != nil {
		return errors.New(log + " action: RequestDo err: " + err.Error())
	}
	defer resp.Body.Close()

	answerStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(log + " action: ResponseRead err: " + err.Error())
	}

	/*
		if strings.Contains(url, `yachtReservation`) {
			f, _ := os.Create(`/Users/madredix/Project/src/github.com/Madredix/theboats/testing/mock/gds/yachtReservation.json`)
			f.Write(answerStr)
			f.Close()
		}
	*/

	err = json.Unmarshal(answerStr, &v)
	if err != nil {
		return errors.New(log + " action: ResponseUnmarshal statusCode: " + strconv.Itoa(resp.StatusCode) + " body: " +
			strings.Replace(string(answerStr), "\n", "\\n", -1) + " err: " + err.Error())
	}

	return nil
}
