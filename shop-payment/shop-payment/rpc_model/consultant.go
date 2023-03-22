package rpc_model


type AAA struct {
	Name string `json:"name"`
}

type ConsultantResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type GetConsultantRequest struct {
	Consns []int64    `json:"conSns"`
	Fields []string `json:"fields"`
}

type GetConsultantResponse struct {
	Basiccname  string `json:"basicCname"`
	Nationality string `json:"nationality"`
	Conname     string `json:"conName"`
	Basiclname  string `json:"basicLname"`
	Location    string `json:"location"`
	Consn       int64    `json:"conSn"`
	Basicfname  string `json:"basicFname"`
	Workfor     string `json:"workFor"`
	Status      string `json:"status"`
}
