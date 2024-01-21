package logger

// LogTDRModel or Transaction Data Record
type LogTDRModel struct {
	XTime         string `json:"xtime"`
	AppName       string `json:"app"`
	Runtime       string `json:"runtime"`
	Env           string `json:"env"`
	AppVersion    string `json:"app_version"`
	CorrelationID string `json:"correlation_id"`

	Path         string `json:"path"`
	Method       string `json:"method"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	SrcIP        string `json:"src_ip"`
	RespTime     int64  `json:"rt"`
	ResponseCode int    `json:"rc"`

	Header   interface{} `json:"header"` // better to pass data here as is, don't cast it to string. use map or array
	Request  interface{} `json:"req"`
	Response interface{} `json:"resp"`
	Error    string      `json:"error"`
	TraceID  string      `json:"trace_id"`

	AdditionalData interface{} `json:"additional_data"`
}
