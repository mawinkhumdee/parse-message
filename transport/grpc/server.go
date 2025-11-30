package grpc

import (
	"context"
	"fmt"
	"net"
	"parse-message/model"
	pb "parse-message/proto"
	"parse-message/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMessageServiceServer
	service service.Service
}

func NewServer(service service.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) InsertMessage(ctx context.Context, req *pb.InsertMessageRequest) (*pb.InsertMessageResponse, error) {
	message := model.Message{
		UserID:  req.UserId,
		Content: req.Content,
		Source:  req.Source,
	}

	if err := s.service.InsertMessage(ctx, message); err != nil {
		return &pb.InsertMessageResponse{
			Success: false,
		}, fmt.Errorf("failed to insert message: %w", err)
	}

	return &pb.InsertMessageResponse{
		Success: true,
		Id:      "pending",
	}, nil
}

func Start(port int, service service.Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, NewServer(service))
	reflection.Register(s)

	fmt.Printf("gRPC server listening on port %d\n", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
