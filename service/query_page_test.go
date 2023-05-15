package service

import (
	"Gin_In_Action/resposity"
	"Gin_In_Action/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestQueryPageInfo(t *testing.T) {

	type args struct {
		topicId int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "查询页面",
			args: args{
				topicId: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := QueryPageInfo(tt.args.topicId)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}

}

func TestMain(m *testing.M) {
	if err := resposity.Init(); err != nil {
		os.Exit(1)
	}
	if err := util.InitLogger(); err != nil {
		os.Exit(1)
	}
	exitCode := m.Run()

	os.Exit(exitCode)
}
