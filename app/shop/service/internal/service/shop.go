package service

import (
	pb "casso/api/shop/service/v1"
	"casso/app/shop/service/internal/biz"
	"casso/pkg/errors/normal"
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

var MockUID int64 = 2233

type ShopService struct {
	pb.UnimplementedShopServer

	sc  *biz.ShopUseCase
	log *log.Helper
}

func NewShopService(sc *biz.ShopUseCase, logger log.Logger) *ShopService {
	return &ShopService{
		sc:  sc,
		log: log.NewHelper(log.With(logger, "module", "service/shop"))}
}

// GetUserID 从上下文中获取userid；取出的为interface，需要断言
func (s *ShopService) GetUserID(ctx context.Context) int64 {
	inter_id := ctx.Value("userid")
	if inter_id == nil {
		return MockUID
	}
	id, err := strconv.Atoi(inter_id.(string))
	if err != nil {
		return MockUID
	}
	return int64(id)
}

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	// 数据校验
	if req.Mobile == "" {
		return &pb.RegisterReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.sc.Register(ctx, req)
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	// 数据校验
	if req.Mobile == "" {
		return &pb.LoginReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.sc.Login(ctx, req)
}

func (s *ShopService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	// 调用业务用例
	return s.sc.GetUser(ctx, s.GetUserID(ctx))
}
