package misc

import (
	"errors"
	"net/http"

	"github.com/arcplus/go-lib/errs"
	"github.com/arcplus/go-lib/json"
	"github.com/arcplus/go-lib/pb"
)

// Response, multi v is accpetable. handler err first.
// rw, r
// rw, r, data
// rw, r, err
// rw, r, data, err
// rw, r, data, code
// rw, r, err, code
// rw, r, data, err, code
func Response(rw http.ResponseWriter, r *http.Request, v ...interface{}) {
	rw.Header().Set("content-type", "application/json; charset=utf-8")

	lv := len(v)

	if lv == 0 {
		rw.Write([]byte("{}"))
		return
	}

	var data []byte
	var err error
	var code int

	switch lv {
	case 1:
		// do nothing
	case 2:
		if v1, ok := v[1].(int); ok {
			code = v1
		} else if v[1] != nil {
			v[0] = v[1]
		}
	case 3:
		if v2, ok := v[2].(int); ok {
			code = v2
		}
		if v[1] != nil {
			v[0] = v[1]
		}
	default:
		v[0] = errors.New("func Response usage error")
	}

	switch v0 := v[0].(type) {
	case nil:
		// do nothing
	case *errs.Error:
		if code == 0 {
			code = http.StatusBadRequest
		}
		data, err = json.Marshal(v0)
	case error:
		err = v0
	case pb.Message:
		data, err = pb.Marshal(v0)
	default:
		data, err = json.Marshal(v0)
	}

	if err != nil {
		Response(rw, r, errs.New(errs.CodeInternal, err.Error()), code)
		return
	}

	if code == 0 {
		code = http.StatusOK
	}

	rw.WriteHeader(code)
	rw.Write(data)
}
