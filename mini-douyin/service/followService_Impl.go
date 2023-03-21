package service

import (
	"log"
	"math/rand"
	"mini-douyin/config"
	"mini-douyin/dao"
	"strconv"
	"sync"
	"time"
)

type FollowServiceImpl struct {
	UserService
}

var (
	followService *FollowServiceImpl
	followOnce    sync.Once
)

func GetFollowServiceInstance() *FollowServiceImpl {
	followOnce.Do(func() {
		followService = &FollowServiceImpl{
			UserService: &UserServiceImpl{},
		}
	})
	return followService
}

func CacheTimeGenerator() time.Duration {
	// 先设置随机数 - 这里比较重要
	rand.Seed(time.Now().Unix())
	// 再设置缓存时间
	// 10 + [0~20) 分钟的随机时间
	return time.Duration((10 + rand.Int63n(20)) * int64(time.Minute))
}

func convertToInt64Array(strArr []string) ([]int64, error) {
	int64Arr := make([]int64, len(strArr))
	for i, str := range strArr {
		int64Val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		int64Arr[i] = int64Val
	}
	return int64Arr, nil
}

var _ FollowService = (*FollowServiceImpl)(nil)

// CancelFollowAction implements FollowService
func (*FollowServiceImpl) CancelFollowAction(userId int64, targetId int64) (bool, error) {
	followDao := dao.NewFollowDaoInstance()
	follow, err := followDao.FindEverFollowing(userId, targetId)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	if follow != nil {
		//ever followed, cancel the relation
		_, err = followDao.UpdateFollowRelation(userId, targetId, 0)
		if err != nil {
			log.Println(err.Error())
		}
		return true, nil
	}
	return false, nil
}

// CheckIsFollowing implements FollowService
func (*FollowServiceImpl) CheckIsFollowing(userId int64, targetId int64) (bool, error) {
	followDao := dao.NewFollowDaoInstance()
	isFollow, err := followDao.FindFollowRelation(userId, targetId)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return isFollow, nil
}

// FollowAction implements FollowService
func (*FollowServiceImpl) FollowAction(userId int64, targetId int64) (bool, error) {
	followDao := dao.NewFollowDaoInstance()
	follow, err := followDao.FindEverFollowing(userId, targetId)
	if err != nil {
		return false, err
	}
	if follow != nil {
		//if ever followed, update the relation
		_, err := followDao.UpdateFollowRelation(userId, targetId, 1)
		if err != nil {
			log.Println(err.Error())
			return false, err
		}
		return true, nil
	}
	//follow for the first time
	_, err = followDao.InsertFollowRelation(userId, targetId)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

// GetFollowerCnt implements FollowService
func (*FollowServiceImpl) GetFollowerCnt(userId int64) (int64, error) {
	followDao := dao.NewFollowDaoInstance()
	ids, err := followDao.GetFollowerCnt(userId)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	return int64(ids), nil
}

// GetFollowers implements FollowService
func (*FollowServiceImpl) GetFollowers(userId int64) ([]User, error) {
	followDao := dao.NewFollowDaoInstance()
	userFollowersId, userFollowersCnt, err := followDao.GetFollowersInfo(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	userFollowers := make([]User, userFollowersCnt)
	buildUser(userId, userFollowers, userFollowersId, 1)
	return userFollowers, nil
}

// GetFollowingCnt implements FollowService
func (*FollowServiceImpl) GetFollowingCnt(userId int64) (int64, error) {
	followDao := dao.NewFollowDaoInstance()
	ids, err := followDao.GetFollowingCnt(userId)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	return int64(ids), nil
}

// GetFollowings implements FollowService
func (*FollowServiceImpl) GetFollowings(userId int64) ([]User, error) {
	followDao := dao.NewFollowDaoInstance()
	followingIds, followingCnt, err := followDao.GetFollowingsInfo(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	userFollowings := make([]User, followingCnt)
	buildUser(userId, userFollowings, followingIds, 0)
	return userFollowings, nil

}

// GetFriends implements FollowService
func (*FollowServiceImpl) GetFriends(userId int64) ([]FriendUser, error) {
	followDao := dao.NewFollowDaoInstance()
	userFriendsId, userFriendsCnt, err := followDao.GetFriendsInfo(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	userFriends := make([]FriendUser, userFriendsCnt)
	err = buildFriendUser(userId, userFriends, userFriendsId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return userFriends, nil
}

func buildFriendUser(userId int64, friendUsers []FriendUser, ids []int64) error {

	//获取好友聊天记录
	msi := GetMessageServiceInstance()
	followDao := dao.NewFollowDaoInstance()

	// 遍历传入的好友id，组装好友user结构体
	for i := 0; i < len(ids); i++ {

		// 好友id赋值
		friendUsers[i].Id = ids[i]

		// 好友name赋值
		var err1 error
		friendUsers[i].Name, err1 = followDao.GetUserName(ids[i])
		if nil != err1 {
			log.Println(err1)
			return err1
		}

		// 好友关注数赋值
		var err2 error
		friendUsers[i].FollowCount, err2 = followService.GetFollowingCnt(ids[i])
		if nil != err2 {
			log.Println(err2.Error())
			return err2
		}

		// 好友粉丝数赋值
		var err3 error
		friendUsers[i].FollowerCount, err3 = followService.GetFollowerCnt(ids[i])
		if nil != err3 {
			log.Println(err3.Error())
			return err3
		}

		// 好友其他属性赋值
		friendUsers[i].IsFollow = true
		friendUsers[i].Avatar = config.CUSTOM_DOMAIN + config.OSS_USER_AVATAR_DIR

		// 调用message模块获取聊天记录
		messageInfo, err := msi.LatestMessage(userId, ids[i])

		//在根据id获取不到最新一条消息时，需要返回对应的id
		if err != nil {

			continue
		}

		friendUsers[i].Message = messageInfo.message
		friendUsers[i].MsgType = messageInfo.msgType
	}

	// 将空数组内属性构建完成即可，不用特意返回数组
	return nil
}

// BuildUser 根据传入的id列表和空user数组，构建业务所需user数组并返回
func buildUser(userId int64, users []User, ids []int64, buildtype int) error {
	folowDao := dao.NewFollowDaoInstance()

	// 遍历传入的用户id，组成user结构体
	for i := 0; i < len(ids); i++ {

		// 用户id赋值
		users[i].Id = ids[i]

		// 用户name赋值
		var err1 error
		users[i].Name, err1 = folowDao.GetUserName(ids[i])
		if nil != err1 {
			log.Println(err1)
			return err1
		}

		// 用户关注数赋值
		var err2 error
		users[i].FollowCount, err2 = followService.GetFollowingCnt(ids[i])
		if nil != err2 {
			log.Println(err2.Error())
			return err2
		}

		// 用户粉丝数赋值
		var err3 error
		users[i].FollowerCount, err3 = followService.GetFollowerCnt(ids[i])
		if nil != err3 {
			log.Println(err3.Error())
			return err3
		}

		// 根据传入的buildtype决定是哪种业务的user构建
		if buildtype == 1 {
			// 粉丝用户的isfollow属性需要调用接口再确认一下
			users[i].IsFollow, _ = followService.CheckIsFollowing(userId, ids[i])
		} else {
			// 关注用户的isfollow属性确定是true
			users[i].IsFollow = true
		}

	}
	return nil
}
