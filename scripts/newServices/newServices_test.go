package newServices_test

import (
	"go-crawler/scripts/newServices"
	"testing"
)

/**
用于生成一项服务
*/
func TestNewServices(t *testing.T) {
	gc := newServices.Bro{
		ServiceName: "test",
		ServiceDir:  "../../services",
	}

	if err := newServices.Run(gc); err != nil {
		t.Fatal(err)
	}
}
