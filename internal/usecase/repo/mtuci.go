package repo

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/mysql"
)

type MtuciRepo struct {
	*mysql.MySQL
	l *logger.Logger
}

func NewMtuciRepo(msq *mysql.MySQL, l *logger.Logger) *MtuciRepo {
	return &MtuciRepo{msq, l}
}

func (r *MtuciRepo) GetUserByLoginWithPassword(ctx context.Context, user entity.UserLogin) (*entity.User, error) {
	var (
		foundUser *entity.User
		group     *entity.Group
	)

	if err := r.DB.Table("users").Where("users.email = ?", user.Login).Where("users.pass_text = ?", user.Password).First(&foundUser); err.Error != nil {
		return nil, err.Error
	}
	if err := r.DB.Table("groups").Where("id = ?", foundUser.GroupID).First(&group); err.Error != nil {
		return nil, err.Error
	}

	foundUser.Group = group

	return foundUser, nil
}

func (r *MtuciRepo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var (
		foundUser  *entity.User
		foundGroup *entity.Group
	)
	if err := r.DB.Table("users").Where("id = ?", id).First(&foundUser); err.Error != nil {
		return nil, err.Error
	}

	if err := r.DB.Table("groups").Where("id = ?", foundUser.GroupID).First(&foundGroup); err.Error != nil {
		return nil, err.Error
	}

	foundUser.Group = foundGroup

	return foundUser, nil
}

func (r *MtuciRepo) GetAllGroups(ctx context.Context) ([]*entity.Group, error) {
	var groups []*entity.Group
	if err := r.DB.Table("groups").Find(&groups); err.Error != nil {
		return nil, err.Error
	}

	return groups, nil
}

func (r *MtuciRepo) GetGroupById(ctx context.Context, groupId int) (*entity.Group, error) {
	var group *entity.Group
	if err := r.DB.Table("groups").First(&group, groupId); err.Error != nil {
		return nil, err.Error
	}

	return group, nil
}
