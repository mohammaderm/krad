package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	dto "github.com/mohammaderm/krad/internal/dto/user"
	"github.com/mohammaderm/krad/internal/models"
	"github.com/mohammaderm/krad/log"
)

var (
	CreateUser        = "INSERT INTO user (username, email, password) VALUES (?,?,?);"
	GetUserbyEmail    = "SELECT * FROM user WHERE email=?;"
	GetUserbyusername = "SELECT * FROM user WHERE username=?;"

	CreateComment = "INSERT INTO comment (userid, productid, text, createdat) VALUES (?,?,?,?);"
	GetAllComment = "SELECT * FROM comment WHERE productid = ? LIMIT ? OFFSET ?;"
	limit         = 5
)

type (
	repository struct {
		logger log.Logger
		db     *sqlx.DB
	}

	UserRepository interface {
		// user interfaces
		Create_User(ctx context.Context, email, username, password string) error
		GetbyEmail_User(ctx context.Context, email string) (*models.User, error)
		GetByUserName_User(ctx context.Context, username string) (*models.User, error)

		// comment interfaces
		CreateComment(ctx context.Context, comment models.CreateComment) error
		GetAllComments(ctx context.Context, productid, offset int) (*[]dto.GetAllComment, error)
	}
)

func NewRepository(con *sqlx.DB, logger log.Logger) UserRepository {
	return &repository{
		logger: logger,
		db:     con,
	}
}

func (r *repository) Create_User(ctx context.Context, email, username, password string) error {
	_, err := r.db.ExecContext(ctx, CreateUser, username, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetbyEmail_User(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, GetUserbyEmail, email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetByUserName_User(ctx context.Context, username string) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, GetUserbyEmail, username)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) CreateComment(ctx context.Context, comment models.CreateComment) error {
	_, err := r.db.ExecContext(ctx, CreateComment, comment.UserId, comment.ProductId, comment.Text, comment.Createdat)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllComments(ctx context.Context, productid, offset int) (*[]dto.GetAllComment, error) {
	var result []dto.GetAllComment
	err := r.db.SelectContext(ctx, &result, GetAllComment, productid, limit, offset)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
