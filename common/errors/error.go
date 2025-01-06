package errors

import "encoding/json"

type Err struct {
	Code    int
	Message string
}

const (
	SystemError                   = 1000
	ServerError                   = 1001
	UnknownError                  = 1002
	ParameterError                = 1003
	NetworkAnomaly                = 1004
	OperationTooFrequently        = 1005
	LockCountNotEnough            = 1006
	DataNotFound                  = 1007
	IllegalInvite                 = 1008
	TransferError                 = 1009
	ReachNewestBlock              = 1010
	UriNotFoundOrMethodNotSupport = 1404
)

var ErrCodeText = map[int]string{
	SystemError:                   "System Error",
	ServerError:                   "Server Error",
	UnknownError:                  "Unknown Error",
	ParameterError:                "Parameter Error",
	NetworkAnomaly:                "Network Anomaly",
	OperationTooFrequently:        "Operation too frequently",
	LockCountNotEnough:            "Too much to guaranty",
	DataNotFound:                  "Data not found",
	IllegalInvite:                 "Illegal invite",
	TransferError:                 "Transfer Error",
	UriNotFoundOrMethodNotSupport: "Uri not found or method can not support",
	ReachNewestBlock:              "Filter logs reach newest block",
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int) *Err {
	return &Err{
		Code:    code,
		Message: ErrCodeText[code],
	}
}
