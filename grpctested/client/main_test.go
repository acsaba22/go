package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/acsaba22/go/grpctested/mempb"
	"github.com/acsaba22/go/grpctested/mock_mempb"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
)

type MyMock map[string]string

func (mm MyMock) Get(ctx context.Context, in *mempb.GetRequest, opts ...grpc.CallOption) (*mempb.GetResponse, error) {
	res := mempb.GetResponse{}
	res.Value, res.Exists = mm[in.Key]
	// time.Sleep(time.Second * 2)
	return &res, nil
}

func TestWithOwnMock(t *testing.T) {
	bb := bytes.Buffer{}
	clientDo(MyMock{"42": "life", "300": "sparta"}, "42", &bb)
	fmt.Println("bbstring", bb.String())
	if !strings.Contains(bb.String(), "life") {
		t.Error("Bad response: " + bb.String())
	}
}

func TestWithGoMock1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	bb := bytes.Buffer{}
	mockServer := mock_mempb.NewMockMemServerClient(ctrl)
	mockServer.EXPECT().Get(
		gomock.Any(),
		gomock.Any(), // TODO implement a matcher
	).Return(&mempb.GetResponse{Value: "life", Exists: true}, nil)

	clientDo(mockServer, "42", &bb)
	fmt.Println("bbstring", bb.String())
	if !strings.Contains(bb.String(), "life") {
		t.Error("Bad response: " + bb.String())
	}
}

func TestWithGoMock2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	bb := bytes.Buffer{}
	mockServer := mock_mempb.NewMockMemServerClient(ctrl)
	mockServer.EXPECT().Get(
		gomock.Any(),
		gomock.Any(), // TODO implement a matcher
	).DoAndReturn(func(c context.Context, req *mempb.GetRequest) (*mempb.GetResponse, error) {
		if req.Key != "42" {
			t.Error("bad request key: " + req.Key)
		}
		return &mempb.GetResponse{Value: "life", Exists: true}, nil
	})
	clientDo(mockServer, "42", &bb)
	fmt.Println("bbstring", bb.String())
	if !strings.Contains(bb.String(), "life") {
		t.Error("Bad response: " + bb.String())
	}
}
