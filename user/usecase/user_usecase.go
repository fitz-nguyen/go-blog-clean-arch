package usecase

import (
	"blog/models"
	"blog/user"
	"time"
)

type userUsecase struct {
	articleRepo    user.Repository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewUserUsecase(a user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Save(a *models.User) error {
	var err = u.Save(a)
	if err != nil {
		return err
	}
	return nil
}
