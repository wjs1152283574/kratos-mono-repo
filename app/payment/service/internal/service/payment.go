package service

import (
    "context"

    pb "casso/api/payment/service/v1"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentReply, error) {
	return &pb.CreatePaymentReply{}, nil
}
func (s *PaymentService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentReply, error) {
	return &pb.UpdatePaymentReply{}, nil
}
func (s *PaymentService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentReply, error) {
	return &pb.DeletePaymentReply{}, nil
}
func (s *PaymentService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentReply, error) {
	return &pb.GetPaymentReply{}, nil
}
func (s *PaymentService) ListPayment(ctx context.Context, req *pb.ListPaymentRequest) (*pb.ListPaymentReply, error) {
	return &pb.ListPaymentReply{}, nil
}
