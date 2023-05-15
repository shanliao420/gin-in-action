package resposity

import (
	"Gin_In_Action/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDao_QueryUserById(t *testing.T) {
	output, err := NewUserDaoInstance().QueryUserById(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, output.Name, "Jerry")
	util.Logger.Info(output.Name)
}

func TestUserDao_BenchQueryUsersByIds(t *testing.T) {
	output, err := NewUserDaoInstance().BenchQueryUsersByIds([]int64{1, 2})
	assert.Equal(t, err, nil)
	for id, user := range output {
		fmt.Println(id, *user)
	}
	assert.Equal(t, len(output), 2)
}
