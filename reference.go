package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sutra/service/sys/time"
	t "time"
)

type configuration struct {
	Http      string
	ServiceIP string
	Type      string
	Sensor    string
	StartTime string
	EndTime   string
	Device    string
}

// GetMockData http get uri data
func GetMockData(uri string) []byte {
	res, err := http.Get(uri)
	if err != nil {
		fmt.Println("http.Get device list Err: ", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read response body err: ", err.Error())
	}
	return body
}

func main() {
	uris := InitMapper()

	file, err := os.Open("conf.json")
	if err != nil {
		panic("os.Open conf json err")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("decoder.Decode Error: ", err)
	}

	for k, v := range uris {
		if k == "/api/device-mgt/device/config/service" {
			k += "?ip=" + conf.ServiceIP
		} else if k == "/api/device-mgt/device/sensor/history" {
			k += "?device=" + conf.Device + "&type=" + conf.Type + "&sensor=" + conf.Sensor + "&time_end=" + time.TimeToString(t.Now())
		} else if k == "/api/device-mgt/device/sensor/history/statics" ||
			k == "/api/device-mgt/device/sensor" {
			k += "?device=" + conf.Device + "&type=" + conf.Type + "&sensor=" + conf.Sensor
		} else {
			k += "?device=" + conf.Device
		}

		resp, err := http.Get(conf.Http + k)
		if err != nil {
			fmt.Println("http.Get Err: ", err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read response body err: ", err.Error())
		}

		pathSlice := strings.Split(v, "/")
		var path string
		for i := 0; i < len(pathSlice)-1; i++ {
			path += pathSlice[i] + "/"
		}
		_, err = os.Stat(path)
		if err != nil {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				panic("os.MkdirAll err: " + path)
			}
		}

		file, err := os.Create(v)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = file.Write(body)
		if err != nil {
			panic(err.Error())
		}
	}

}
