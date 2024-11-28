package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra/metrics"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserPostgresRepository {
	return UserPostgresRepository{
		db: db,
	}
}

// All retrieves a list with all users
func (r UserPostgresRepository) All(ctx context.Context) ([]user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	var usersSql []UserPostgres
	if err := r.db.WithContext(ctx).Find(&usersSql).Error; err != nil {
		return nil, fmt.Errorf("[UserPostgresRepository:All][err:%w]", err)
	}

	users := make([]user.User, len(usersSql))
	for i := range usersSql {
		users[i] = usersSql[i].ToDomainModel()
	}

	return users, nil
}

// Save creates a new user
func (r UserPostgresRepository) Save(ctx context.Context, u user.User) error {
	defer metrics.TrackDBQueryDuration(time.Now())

	if err := r.db.WithContext(ctx).Create(&UserPostgres{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return fmt.Errorf("[UserPostgresRepository:Save][err: %w]", err)
	}

	return nil
}

// FindByEmail search a user by email
func (r UserPostgresRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	var u UserPostgres
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return user.User{}, fmt.Errorf("[UserPostgresRepository:FindByEmail][err: %w]", err)
	}

	return u.ToDomainModel(), nil
}
