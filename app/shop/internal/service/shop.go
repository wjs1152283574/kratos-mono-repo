package service

import (
	"context"
	"strconv"

	"casso/app/shop/internal/biz"
	pb "casso/gen/shop/v1"
	"casso/pkg/contextx"
	"casso/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

// MockUID 默认用户 ID（开发用）
var MockUID int64 = 2333

// ShopService BFF 服务实现
type ShopService struct {
	pb.UnimplementedShopServer

	sc  *biz.ShopUseCase
	log *log.Helper
}

// NewShopService 创建 ShopService
func NewShopService(sc *biz.ShopUseCase, logger log.Logger) *ShopService {
	return &ShopService{
		sc:  sc,
		log: log.NewHelper(log.With(logger, "module", "service/shop")),
	}
}

// GetUserID 从上下文中获取 userid
func (s *ShopService) GetUserID(ctx context.Context) int64 {
	interID := ctx.Value(contextx.UserID)
	if interID == nil {
		return MockUID
	}
	id, err := strconv.Atoi(interID.(string))
	if err != nil {
		return MockUID
	}
	return int64(id)
}

// Register 注册
func (s *ShopService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if req.Mobile == "" {
		return &pb.RegisterReply{}, errors.InvalidParams
	}
	return s.sc.Register(ctx, req)
}

// Login 登录
func (s *ShopService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	if req.Mobile == "" {
		return &pb.LoginReply{}, errors.InvalidParams
	}
	return s.sc.Login(ctx, req)
}

// GetUser 获取用户信息
func (s *ShopService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return s.sc.GetUser(ctx, s.GetUserID(ctx))
}
