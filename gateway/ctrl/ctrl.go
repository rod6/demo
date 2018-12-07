package ctrl

import (
	"context"
	"demo/gateway/rpc"
	"demo/misc"
	"demo/pb"
	"net/http"

	"github.com/arcplus/go-lib/router"
)

func UserMethodX(rw http.ResponseWriter, r *http.Request, params router.Params) {
	// get req
	id := params.ByName("id")
	resp, err := rpc.UserClient.MethodX(context.Background(), &pb.OK{})
	if err != nil {
		misc.Response(rw, r, err)
	}

	result := map[string]interface{}{
		"id":            id,
		"user_method_x": resp,
	}

	misc.Response(rw, r, result)
}

func ThingsMethodY(rw http.ResponseWriter, r *http.Request) {
	resp, err := rpc.ThingsClient.MethodY(context.Background(), &pb.OK{})
	misc.Response(rw, r, resp, err)
}

func TestPanic(rw http.ResponseWriter, r *http.Request) {
	panic("woops")
}
