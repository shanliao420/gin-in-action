package service

import (
	"Gin_In_Action/resposity"
	"errors"
	"sync"
)

type TopicInfo struct {
	Topic *resposity.Topic
	User  *resposity.User
}

type PostInfo struct {
	Post *resposity.Post
	User *resposity.User
}

type PageInfo struct {
	TopicInfo *TopicInfo
	PostList  []*PostInfo
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{topicId: topicId}
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic   *resposity.Topic
	posts   []*resposity.Post
	userMap map[int64]*resposity.User
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	var topicErr, postErr error
	// 获取topic信息
	go func() {
		defer waitGroup.Done()
		topic, err := resposity.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()

	// 获取post列表
	go func() {
		defer waitGroup.Done()
		posts, err := resposity.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	waitGroup.Wait()

	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}

	// 获取user信息
	uids := []int64{f.topic.UserId}
	for _, post := range f.posts {
		uids = append(uids, post.UserId)
	}
	userMap, err := resposity.NewUserDaoInstance().BenchQueryUsersByIds(uids)
	if err != nil {
		return err
	}
	f.userMap = userMap

	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	userMap := f.userMap
	topicUser, ok := userMap[f.topic.UserId]
	if !ok {
		return errors.New("has no topic user info")
	}

	postList := make([]*PostInfo, 0)
	for _, post := range f.posts {
		postUser, ok := userMap[post.UserId]
		if !ok {
			return errors.New("has no post user info")
		}
		postList = append(postList, &PostInfo{post, postUser})
	}

	f.pageInfo = &PageInfo{
		TopicInfo: &TopicInfo{Topic: f.topic, User: topicUser},
		PostList:  postList,
	}
	return nil
}
