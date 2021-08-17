package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"lxtkj.cn/go-micro-member/proto/member/user"
	"lxtkj.cn/go-micro-member/services"

	"log"
)

//type UserService struct {
//	client client2.Client
//}
//
//func(this *UserService) Test(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
//	rsp.Ret = "users" + req.Id
//	//服务之间调用
//	c := course.NewCourseService("go.micro.api.course",this.client)
//	course_rsp,_ := c.ListForTop(ctx,&course.ListRequest{Size:10})
//	log.Println(course_rsp.Result)
//	return nil
//}
//
//func NewUserService(c client2.Client) *UserService {
//	return &UserService{client:c}
//}

type logWrapper struct {
	client.Client
}

func(this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx,req,rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	fmt.Println("进入NewLogWrapper")
	return &logWrapper{c}
}

func main(){
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		micro.WrapClient(NewLogWrapper),//装饰器失效
		)
	service.Init()

	//获取member项目的ServiceHandler
	userServiceHandler := services.NewUserService("userService",service.Client())
	err := user.RegisterUserServiceHandler(service.Server(), userServiceHandler)
	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
