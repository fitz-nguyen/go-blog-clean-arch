package user

import (
	"blog/models"
)

// Repository represent the article's repository contract
type Usecase interface {
	Save(a *models.User) error
	// Find(ctx context.Context, a *models.User) error
}
