package husgrpc

import (
	"golang.org/x/net/context"
	"grpc-getting-started/proto"
	"time"
)

type HusService struct {
}

type ReqData struct {
	FullName string    `json:"full_name"`
	BirthDay time.Time `json:"birth_year"`
	Point    int       `json:"point"`
}

func (s *HusService) GetMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	result := &proto.MessageResponse{
		Message: "Hello" + req.Text + ", I'm Khoa",
	}

	return result, nil
}

func (s *HusService) GetAccount(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	now := time.Now().String()
	result := &proto.UserResponse{
		FullName: "Huynh Tan Khoa",
		BirthDay: now,
		Point:    300,
	}

	return result, nil
}
