package services

import (
	"context"
	"fmt"
	client2 "github.com/micro/go-micro/v2/client"
	"log"
	"lxtkj.cn/go-micro-member/proto/course/course"
	"lxtkj.cn/go-micro-member/proto/member/user"
	"lxtkj.cn/go-micro-member/repositories"
)

type UserInterfaceService interface {
	Test(ctx context.Context, in *user.UserRequest, out *user.UserResponse) error
	GetUserInfo(context.Context, *user.GetUserByIdRequest, *user.GetUserByIdResponse) error
}

//初始化对象函数
func NewUserService(name string,c client2.Client) UserInterfaceService {
	return &UserService{
		client:c,
		name:name,
		shopMemberService:       repositories.NewBbsDiscuss(),
		bbsPostService:          repositories.NewBbsPost(),
	}
}

type UserService struct {
	client client2.Client
	name string
	shopMemberService 			repositories.BbsDiscussInterface     //商城会员服务
	bbsPostService 			    repositories.BbsPostInterface        //社区帖子服务
}

func(this *UserService) Test(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	rsp.Ret = "users" + req.Id
	//服务之间调用
	c := course.NewCourseService("go.micro.api.course",this.client)
	course_rsp,_ := c.ListForTop(ctx,&course.ListRequest{Size:10})
	log.Println(course_rsp.Result)
	return nil
}

func(this *UserService) GetUserInfo(ctx context.Context, req *user.GetUserByIdRequest, rsp *user.GetUserByIdResponse) error{
	fmt.Println("进入GetUserInfo方法")
	userArr := make([]*user.User, 0)
	var userInfo user.User
	userInfo.Id = req.UserId
	userInfo.UserName = "雷小天"
	userInfo.UserAge = "25"
	userArr = append(userArr,&userInfo)
	rsp.Result = userArr
	return nil
}