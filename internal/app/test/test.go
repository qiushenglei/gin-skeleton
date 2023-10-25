package test

import (
	"github.com/qiushenglei/gin-skeleton/cmd"
	"testing"
)

func TestMain(m *testing.M) {

	cmd.RegistAll("test")

	m.Run()
}
