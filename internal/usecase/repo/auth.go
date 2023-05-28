package repo

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/mysql"
)

type AuthRepo struct {
	*mysql.MySQL
	l *logger.Logger
}

func NewAuthRepo(msq *mysql.MySQL, l *logger.Logger) *AuthRepo {
	return &AuthRepo{msq, l}
}

func (r *AuthRepo) GetUserByLoginWithPassword(ctx context.Context, user entity.UserLogin) (*entity.User, error) {
	var foundUser *entity.User
	if err := r.DB.Table("users").Where("email = ?", user.Login).Where("pass_text = ?", user.Password).First(&foundUser); err.Error != nil {
		return nil, err.Error
	}

	return foundUser, nil
}

func (r *AuthRepo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var user *entity.User
	if err := r.DB.Table("users").First(&user, id); err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}
