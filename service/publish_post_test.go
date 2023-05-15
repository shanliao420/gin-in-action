package service

import (
	"Gin_In_Action/util"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestPublishPost(t *testing.T) {

	type args struct {
		topicId int64
		userId  int64
		content string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发布回帖",
			args: args{
				topicId: 1,
				userId:  2,
				content: "再次发帖！",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := PublishPost(tt.args.topicId, tt.args.userId, tt.args.content)
			assert.Equal(t, err != nil, tt.wantErr)
			util.Logger.Info(strconv.Itoa(int(output)))
		})
	}

}
