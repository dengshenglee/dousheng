package service

import (
	"log"
	"mini-douyin/dao"
	"sync"
	"time"
)

type LikeServiceImpl struct {
	VideoService
}

var (
	likeServiceImpl *LikeServiceImpl
	likeServiceOnce sync.Once
)

func GetLikeServImpInstance() *LikeServiceImpl {
	likeServiceOnce.Do(func() {
		likeServiceImpl = &LikeServiceImpl{
			VideoService: GetVideoServiceInstance(),
		}
	})
	return likeServiceImpl
}

var _ LikeService = (*LikeServiceImpl)(nil)

// FavoriteAction implements LikeService
func (*LikeServiceImpl) FavoriteAction(userId int64, videoId int64, actionType int32) error {
	isLike, err := dao.IsLikedByUser(userId, videoId)
	if err != nil {
		log.Println("database err")
		return err
	}
	if isLike == false {
		//user didn't like this video
		likeInfo := dao.Like{
			Id:        userId,
			UserId:    userId,
			VideoId:   videoId,
			Liked:     int8(actionType),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		err = dao.InsertLikeInfo(likeInfo)
		if err != nil {
			return err
		}
	} else {
		err = dao.UpdateLikeInfo(userId, videoId, int8(actionType))
		if err != nil {
			return err
		}
	}
	return nil
}

// GetLikesList implements LikeService
func (*LikeServiceImpl) GetLikesList(userId int64) ([]Video, error) {
	likeVideos, err := dao.GetLikeListByUserId(userId)
	if err != nil {
		log.Println("GetLikeListByUserId:", err)
		return nil, err
	}
	likeService := GetLikeServImpInstance()
	likedVideoInfoList, err := likeService.GetVideoListById(likeVideos, userId)
	if err != nil {
		log.Println("Get videoList failed")
	}
	return likedVideoInfoList, nil
}

// GetUserLikeCount implements LikeService
func (*LikeServiceImpl) GetUserLikeCount(userId int64) (int64, error) {
	likeList, err := dao.GetLikeListByUserId(userId)
	if err != nil {
		log.Println("GetLikeListByUserId:", err)
		return -1, err
	}
	return int64(len(likeList)), nil
}

// GetUserLikedCnt implements LikeService
func (*LikeServiceImpl) GetUserLikedCnt(userId int64) (int64, error) {
	likedList, err := dao.GetUserVideoLikedTotalCount(userId)
	if err != nil {
		log.Println("GetUserVideoLikedTotalCount:", err)
		return -1, err
	}
	return likedList, nil
}

// GetVideoLikedCount implements LikeService
func (*LikeServiceImpl) GetVideoLikedCount(videoId int64) (int64, error) {
	likedCnt, err := dao.VideoLikedCount(videoId)
	if err != nil {
		log.Println("VideoLikedCount:", err)
		return -1, err
	}
	return likedCnt, nil
}

// IsLikedByUser implements LikeService
func (*LikeServiceImpl) IsLikedByUser(userId int64, videoId int64) (bool, error) {
	liked, err := dao.IsLikedByUser(userId, videoId)
	if err != nil {
		log.Println("IsLikedByUser:", err)
		return false, err
	}
	return liked, nil
}
