package user

import (
	"context"

	"github.com/mohammaderm/krad/log"

	dto "github.com/mohammaderm/krad/internal/dto/user"
	"github.com/mohammaderm/krad/internal/models"
	"github.com/mohammaderm/krad/internal/repository/user"
)

type (
	Service struct {
		logger         log.Logger
		userRepository user.UserRepository
	}
	UserServiceContracts interface {
		// user services
		Create_User(ctx context.Context, req dto.CreateUserReq) error
		GetbyEmail_User(ctx context.Context, req dto.GetByEmailReq) (dto.GetByEmailRes, error)
		GetByUserName_User(ctx context.Context, req dto.GetByUsernameReq) (dto.GetByUsernameRes, error)

		// comment services
		CreateComment(ctx context.Context, req dto.CreateCommentReq) error
	}
)

func NewService(logger log.Logger, userrepository user.UserRepository) UserServiceContracts {
	return &Service{
		logger:         logger,
		userRepository: userrepository,
	}
}

func (s *Service) Create_User(ctx context.Context, req dto.CreateUserReq) error {
	err := s.userRepository.Create_User(ctx, req.Email, req.UserName, req.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetbyEmail_User(ctx context.Context, req dto.GetByEmailReq) (dto.GetByEmailRes, error) {
	user, err := s.userRepository.GetbyEmail_User(ctx, req.Email)
	if err != nil {
		return dto.GetByEmailRes{}, err
	}
	return dto.GetByEmailRes{
		User: user,
	}, nil
}

func (s *Service) GetByUserName_User(ctx context.Context, req dto.GetByUsernameReq) (dto.GetByUsernameRes, error) {
	user, err := s.userRepository.GetByUserName_User(ctx, req.UserName)
	if err != nil {
		return dto.GetByUsernameRes{}, err
	}
	return dto.GetByUsernameRes{
		User: user,
	}, nil
}

func (s *Service) CreateComment(ctx context.Context, req dto.CreateCommentReq) error {
	err := s.userRepository.CreateComment(ctx, models.CreateComment{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Createdat: req.Createdat,
		Text:      req.Text,
	})
	if err != nil {
		return err
	}
	return nil
}
