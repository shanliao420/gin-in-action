package resposity

import (
	"Gin_In_Action/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	assert.Equal(t, err, nil)
}

func TestMain(m *testing.M) {
	err := Init()
	if err != nil {
		return
	}

	err = util.InitLogger()
	if err != nil {
		return
	}

	exitCode := m.Run()

	//fmt.Printf("%#v\n", db)

	os.Exit(exitCode)
}
