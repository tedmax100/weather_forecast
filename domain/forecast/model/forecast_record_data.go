package model

type RecordData struct {
	Location       string
	Element        string
	ParameterName  string
	ParameterUnit  *string
	ParameterValue *int
	StartTime      string
	EndTime        string
}

type Records struct {
	Location []Location `json:"location"`
}

type Location struct {
	LocationName   string           `json:"locationName"`
	WeatherElement []WeatherElement `json:"weatherElement"`
}

type WeatherElement struct {
	ElementName string       `json:"elementName"`
	Time        []TimeRecord `json:"time"`
}

type TimeRecord struct {
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Parameter Parameter `json:"parameter"`
}

type Parameter struct {
	ParameterName  string  `json:"parameterName"`
	ParameterUnit  *string `json:"parameterUnit,omitempty"`
	ParameterValue *int    `json:"parameterValue,omitempty"`
}
