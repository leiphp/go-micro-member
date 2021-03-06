package repositories

import (
	"lxtkj.cn/go-micro-member/datamodels"
)


/*
	操作bbs_post表的接口定义

	作者名称：leixiaotian 创建时间：20210412
*/

type BbsPostInterface interface {
	SelectInfo(id int64) (datamodels.BbsPost, error) //获得帖子信息
}

//返回结构体对象
func NewBbsPost() BbsPostInterface {
	return &bbsPost{}
}

//bbsPost构体
type bbsPost struct {
}

//获得主播信息
func (this *bbsPost) SelectInfo(id int64) (datamodels.BbsPost, error) {

	var bbsPostInfo datamodels.BbsPost
	//redis礼物key
	//jmfMemberKey := ReturnRedisKey(API_CACHE_JMF_MEMBER, userId)
	//result, err := initialize.RedisCluster.Get(jmfMemberKey).Bytes()


	return bbsPostInfo, nil
}
