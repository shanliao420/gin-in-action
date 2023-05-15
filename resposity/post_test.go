package resposity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostDao_QueryPostById(t *testing.T) {
	output, err := NewPostDaoInstance().QueryPostById(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, output.Content, "举手报名！")
}

func TestPostDao_QueryPostByParentId(t *testing.T) {
	output, err := NewPostDaoInstance().QueryPostByParentId(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(output), 2)
}

func TestPostDao_CreatePost(t *testing.T) {
	err := NewPostDaoInstance().CreatePost(&Post{ParentId: 2, UserId: 1, Content: "举手报名+2", DiggCount: 30, CreateTime: time.Now()})
	assert.Equal(t, err, nil)
}
