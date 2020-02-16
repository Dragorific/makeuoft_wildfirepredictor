package apilib

//Data stores sensor data and timestamps
type Data struct {
	Sensors []string `json:"sensors"`   //Sensors holds the sensor data
	Time    string   `json:"timestamp"` //Time holds the time stamp used for sorting
}

//ParsedData stores the sorted data for the frontend
type ParsedData struct {
	Temp     []float32 `json:"temp"`     //Temp holds the parsed tempreture data list
	Humidity []float32 `json:"humidity"` //Humidity holds the parsed humidity data list
	Light    []int     `json:"light"`    //Light holds the parsed light data list
}

/*//ParseSensorData parses the data into usable json for the front-end
func ParseSensorData(data []byte) (Data, error) {
	var dataList Data
	err := json.Unmarshal(data, &dataList)
	var temp[30] float32
	var hum[30] float32
	var light[30] int
	for i, newData:= range dataList.Sensors{
		newData[]
	}
	return dataList, err
}

//GetSensorData returns the parsed data that the front-end uses
func GetSensorData(Data) (ParsedData, error) {

}*/
