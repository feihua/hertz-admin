package http

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

func Post(url string, data []byte) ([]byte, error) {
	newClient, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	req := &protocol.Request{}
	res := &protocol.Response{}

	req.Reset()
	req.Header.SetMethod(consts.MethodPost)
	//req.SetHeader("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6IjEiLCJ1c2VybmFtZSI6Iui2hee6p-euoeeQhuWRmCIsInBlcm1pc3Npb25zIjpbIiIsIiIsIi9hcGkvc3lzdGVtL3VzZXIvbGlzdCIsIi9hcGkvc3lzdGVtL3JvbGUvbGlzdCIsIi9hcGkvc3lzdGVtL21lbnUvbGlzdCIsIi9hcGkvc3lzdGVtL3VzZXIvdXBkYXRlU3RhdHVzIiwiL2FwaS9zeXN0ZW0vdXNlci9zYXZlIiwiL2FwaS9zeXN0ZW0vdXNlci9kZWxldGUiLCIvYXBpL3N5c3RlbS91c2VyL3ZpZXciLCIvYXBpL3N5c3RlbS91c2VyL3VwZGF0ZSIsIi9hcGkvc3lzdGVtL3VzZXIvcGFzc3dvcmQiLCIvYXBpL3N5c3RlbS91c2VyL3VwZGF0ZVBhc3N3b3JkIiwiL2FwaS9zeXN0ZW0vcm9sZS9saXN0IiwiL2FwaS9zeXN0ZW0vcm9sZS91c2VyUm9sZVNhdmUiLCIvYXBpL3N5c3RlbS9yb2xlL3VzZXJSb2xlTGlzdCIsIi9hcGkvdXNlci9wYy91c2VyL2xpc3QiLCIvYXBpL3N5c3RlbS9yb2xlL3VwZGF0ZVN0YXR1cyIsIi9hcGkvc3lzdGVtL3JvbGUvc2F2ZSIsIi9hcGkvc3lzdGVtL3JvbGUvZGVsZXRlIiwiL2FwaS9zeXN0ZW0vcm9sZS92aWV3IiwiL2FwaS9zeXN0ZW0vcm9sZS91cGRhdGUiLCIvYXBpL3N5c3RlbS9tZW51L2xpc3QiLCIvYXBpL3N5c3RlbS9tZW51L3JvbGVNZW51TGlzdCIsIi9hcGkvc3lzdGVtL21lbnUvcm9sZU1lbnVTYXZlIiwiL2FwaS9zeXN0ZW0vbWVudS91cGRhdGVTdGF0dXMiLCIvYXBpL3N5c3RlbS9tZW51L3NhdmUiLCIvYXBpL3N5c3RlbS9tZW51L2RlbGV0ZSIsIi9hcGkvc3lzdGVtL21lbnUvdmlldyIsIi9hcGkvc3lzdGVtL21lbnUvdXBkYXRlIiwiL2FwaS9yb2xlX2xpc3QiLCIvYXBpL3N5c3RlbS9lbnVtL2xpc3QiLCIvYXBpL3VzZXJfbGlzdCIsIi9hdXRoL3BpbmciXSwiYXVkIjoicnVzdF9hZG1pbiIsImV4cCI6MTY3MjAwNjUyOSwiaWF0IjoxNjcwMjA2NTI5LCJpc3MiOiJrb29iZSIsIm5iZiI6MTY3MDIwNjUyOSwic3ViIjoicnVzdF9hZG1pbiIsImp0aSI6Imlnbm9yZSJ9.v7S1xFXIrg6seCc1-gfMhIETKAptXqM0o2B9DMvhc1s")
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI(url)

	req.SetBody(data)
	newClient.Use(MyMiddleware)
	err = newClient.Do(context.Background(), req, res)
	return res.Body(), err
}

func MyMiddleware(next client.Endpoint) client.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		// pre-handle
		log.Printf("pre-clientHandle url: %s, req_data: %s", req.Path(), string(req.Body()))
		err = next(ctx, req, resp)
		if err != nil {
			return
		}
		// post-handle
		log.Printf("post-clientHandle status: %d, resp_data: %s", resp.StatusCode(), string(resp.Body()))
		return nil
	}
}
