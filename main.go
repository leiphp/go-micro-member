package main

import (
	"github.com/micro/go-micro/v2"
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

func main(){
	service := micro.NewService(
		micro.Name("go.micro.api.user"))
	service.Init()

	//获取member项目的ServiceHandler
	userServiceHandler := services.NewUserService(service.Client())
	err := user.RegisterUserServiceHandler(service.Server(), userServiceHandler)
	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
