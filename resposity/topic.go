package resposity

import (
	"Gin_In_Action/util"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"colum:id"`
	UserId     int64     `gorm:"colum:user_d"`
	Title      string    `gorm:"colum:title"`
	Content    string    `gorm:"colum:content"`
	CreateTime time.Time `gorm:"colum:create_time"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		util.Logger.Error("find topic by id error:" + err.Error())
		return nil, err
	}
	return &topic, nil
}
