package apilib

import (
	"encoding/json"
)

//Data stores sensor data and timestamps
type Data struct {
	Sensors []string `json:"sensors"`   //Sensors holds the sensor data
	Time    string   `json:"timestamp"` //Time holds the time stamp used for sorting
}

//ParsedData stores the sorted data for the frontend
type ParsedData struct {
	Temp     [30]string `json:"temp"`     //Temp holds the parsed tempreture data list
	Humidity [30]string `json:"humidity"` //Humidity holds the parsed humidity data list
	Light    [30]string `json:"light"`    //Light holds the parsed light data list
}

//ParseSensorData parses the data into usable json for the front-end
func ParseSensorData(data [30][]byte) []byte {
	var dataList [30]Data
	for _, newData := range data {
		_ = json.Unmarshal(newData, &dataList)
	}
	var temp [30]string
	var hum [30]string
	var light [30]string
	for i, newData := range dataList {
		for _, list := range newData.Sensors {
			hum[i] = string(list[0])
			temp[i] = string(list[1])
			light[i] = string(list[2])
		}
	}
	parsed := &ParsedData{Temp: temp, Humidity: hum, Light: light}
	encjson, _ := json.Marshal(parsed)
	return encjson
}
