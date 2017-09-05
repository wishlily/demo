package main

import (
	"fmt"
	_ "net"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	buf := make([]byte, 516)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	conn := NewMockConn(mockCtrl)
	conn.EXPECT().Read(buf).Return(4, nil)
	b, err := get(conn)
	if err != nil {
		t.Fail()
	}
	fmt.Println(b)
}
