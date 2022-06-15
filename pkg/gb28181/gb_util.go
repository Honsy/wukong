package gb28181

import (
	"fmt"
)

// Error Error
type Error struct {
	err    error
	params []interface{}
}

func (err *Error) Error() string {
	if err == nil {
		return "<nil>"
	}
	str := fmt.Sprint(err.params...)
	if err.err != nil {
		str += fmt.Sprintf(" err:%s", err.err.Error())
	}
	return str
}

// NewError NewError
func NewError(err error, params ...interface{}) error {
	return &Error{err, params}
}

// func sipResponse(tx *sip.Transaction) (*sip.Response, error) {
// 	response := tx.GetResponse()
// 	if response == nil {
// 		return nil, NewError(nil, "response timeout", "tx key:", tx.Key())
// 	}
// 	if response.StatusCode() != http.StatusOK {
// 		return response, NewError(nil, "response fail", response.StatusCode(), response.Reason(), "tx key:", tx.Key())
// 	}
// 	return response, nil
// }
