package dtos

type KrlResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Station `json:"data"`
}

type Station struct {
	StaID    string `json:"sta_id"`
	StaName  string `json:"sta_name"`
	GroupWil int    `json:"group_wil"`
	FgEnable int    `json:"fg_enable"`
}

type ScheduleResponse struct {
	Status int        `json:"status"`
	Data   []Schedule `json:"data"`
}

type Schedule struct {
	TrainID   string `json:"train_id"`
	KaName    string `json:"ka_name"`
	RouteName string `json:"route_name"`
	Dest      string `json:"dest"`
	TimeEst   string `json:"time_est"`
	Color     string `json:"color"`
	DestTime  string `json:"dest_time"`
}

type ScheduleTrainResponse struct {
	Status int             `json:"status"`
	Data   []ScheduleTrain `json:"data"`
}

type ScheduleTrain struct {
	TrainID        string      `json:"train_id"`
	KaName         string      `json:"ka_name"`
	StationID      string      `json:"station_id"`
	StationName    string      `json:"station_name"`
	TimeEst        string      `json:"time_est"`
	TransitStation bool        `json:"transit_station"`
	Color          string      `json:"color"`
	Transit        interface{} `json:"transit"` // bisa []string atau string kosong
}
