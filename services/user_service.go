package services

import (
	"context"
	client2 "github.com/micro/go-micro/v2/client"
	"log"
	"lxtkj.cn/go-micro-member/proto/course/course"
	"lxtkj.cn/go-micro-member/proto/member/user"
)

type UserInterfaceService interface {
	Test(ctx context.Context, in *user.UserRequest, out *user.UserResponse) error
}

//初始化对象函数
func NewUserService(c client2.Client) UserInterfaceService {
	return &UserService{
		client:c,
	}
}

type UserService struct {
	client client2.Client
}

func(this *UserService) Test(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	rsp.Ret = "users" + req.Id
	//服务之间调用
	c := course.NewCourseService("go.micro.api.course",this.client)
	course_rsp,_ := c.ListForTop(ctx,&course.ListRequest{Size:10})
	log.Println(course_rsp.Result)
	return nil
}
