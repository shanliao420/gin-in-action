package resposity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopicDao_QueryTopicById(t *testing.T) {
	output, err := NewTopicDaoInstance().QueryTopicById(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, output.Title, "青训营开课啦")
}
