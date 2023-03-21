package service

import (
	"log"
	"mini-douyin/config"
	"mini-douyin/dao"
	"sync"
)

type UserServiceImpl struct {
	// 关注服务
	FollowService
	// 点赞服务
	LikeService
	// 视频服务
	VideoService
}

var (
	userServiceImp  *UserServiceImpl
	userServiceOnce sync.Once
)

func GetUserServiceInstance() *UserServiceImpl {
	userServiceOnce.Do(func() {
		userServiceImp = &UserServiceImpl{
			FollowService: GetFollowServiceInstance(),
			LikeService:   GetLikeServImpInstance(),
			VideoService:  GetVideoServiceInstance(),
		}
	})
	return userServiceImp
}

var _ UserService = (*UserServiceImpl)(nil)

// GetUserBasicInfoById implements UserService
func (*UserServiceImpl) GetUserBasicInfoById(id int64) dao.UserBasicInfo {
	user, err := dao.GetUserBasicInfoById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user
	}
	log.Println("Query User Success")
	return user
}

// GetUserBasicInfoByName implements UserService
func (*UserServiceImpl) GetUserBasicInfoByName(name string) dao.UserBasicInfo {
	user, err := dao.GetUserBasicInfoByName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user
	}
	log.Println("Query User Success")
	return user
}

// GetUserLoginInfoById implements UserService
func (*UserServiceImpl) GetUserLoginInfoById(id int64) (user User, err error) {
	plainUser, err := dao.GetUserBasicInfoById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
	}
	user.Id = plainUser.Id
	user.Name = plainUser.Name

	userService := GetUserServiceInstance()
	var wg sync.WaitGroup
	wg.Add(5)
	go func(id int64) {
		followCnt, err := userService.GetFollowingCnt(id)
		if err != nil {
			return
		}
		user.FollowCount = followCnt
		wg.Done()
	}(id)
	go func(id int64) {
		followerCnt, err := userService.GetFollowerCnt(id)
		if err != nil {
			return
		}
		user.FollowerCount = followerCnt
		wg.Done()
	}(id)
	go func(id int64) {
		workCount, err := userService.GetVideoCnt(id)
		if err != nil {
			return
		}
		user.WorkCount = workCount
		wg.Done()
	}(id)
	go func(id int64) {
		// 计算被点赞数, 找出用户被点赞的视频，循环求和:在likeservide实现
		totalFavorited, err := userService.GetUserLikedCnt(id)
		if err != nil {
			return
		}
		user.TotalFavorited = totalFavorited
		wg.Done()
	}(id)
	go func(id int64) {
		// 计算喜欢数量
		favoriteCount, err := userService.GetUserLikeCount(id)
		if err != nil {
			return
		}
		user.FavoriteCount = favoriteCount
		wg.Done()
	}(id)
	wg.Wait()
	return
}

// GetUserLoginInfoByIdWithCurId 登录情况下返回用户信息, 第一个id是视频作者的id，第二个id是我们用户的id
func (*UserServiceImpl) GetUserLoginInfoByIdWithCurId(authorId int64, curId int64) (user User, err error) {
	plainUser, err := dao.GetUserBasicInfoById(authorId)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("USER NOT FOUND")
	}
	user.Id = plainUser.Id
	user.Name = plainUser.Name

	userService := GetUserServiceInstance()
	var wg sync.WaitGroup
	wg.Add(6)
	go func(id int64) {
		followCnt, err := userService.GetFollowingCnt(id)
		if err != nil {
			return
		}
		user.FollowCount = followCnt
		wg.Done()
	}(authorId)
	go func(id int64) {
		followerCnt, err := userService.GetFollowerCnt(id)
		if err != nil {
			return
		}
		user.FollowerCount = followerCnt
		wg.Done()
	}(authorId)
	go func(id, curId int64) {
		//check whether the curId user follow the video author
		isFollowed, err := userService.CheckIsFollowing(curId, id)
		if err != nil {
			return
		}
		user.IsFollow = isFollowed
		wg.Done()
	}(authorId, curId)
	go func(id int64) {
		// 计算作品数
		workCount, err := userService.GetVideoCnt(id)
		if err != nil {
			return
		}
		user.WorkCount = workCount
		wg.Done()
	}(authorId)

	go func(id int64) {
		// 计算被点赞数, 找出用户被点赞的视频，循环求和:在likeservide实现
		totalFavorited, err := userService.GetUserLikedCnt(id)
		if err != nil {
			return
		}
		user.TotalFavorited = totalFavorited
		wg.Done()
	}(authorId)

	go func(id int64) {
		// 计算喜欢数量
		favoriteCount, err := userService.GetUserLikeCount(id)
		if err != nil {
			return
		}
		user.FavoriteCount = favoriteCount
		wg.Done()
	}(authorId)
	wg.Wait()

	user.Avatar = config.CUSTOM_DOMAIN + config.OSS_USER_AVATAR_DIR
	user.BackgroundImage = config.BG_IMAGE
	user.Signature = config.SIGNATURE
	return
}

// InsertUser implements UserService
func (*UserServiceImpl) InsertUser(user *dao.UserBasicInfo) bool {
	flag := dao.InsertUser(user)
	if flag == false {
		log.Println("Insert fail!")
		return false
	}
	return true
}
