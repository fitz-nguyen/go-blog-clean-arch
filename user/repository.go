package user

import (
	"blog/models"
)

// Repository represent the article's repository contract
type Repository interface {
	Save(a *models.User) error
	// Find(ctx context.Context, a *models.User) error
}
