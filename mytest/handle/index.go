package handle

import (
	"context"
	api "github.com/micro/go-api/proto"
	"hm-core/core"
)

type User struct{}

func (e *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {




	//post := core.MultipartPost(req)
	//log.Println("post",post)

	//return core.Success(rsp,post)
	core.Success(rsp,"444666")


	return nil

}