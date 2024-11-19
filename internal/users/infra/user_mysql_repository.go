package infra

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/infra/metrics"
)

type UserMysqlRepository struct {
	db *gorm.DB
}

func NewUserMysqlRepository(db *gorm.DB) UserMysqlRepository {
	return UserMysqlRepository{
		db: db,
	}
}

// All retrieves a list with all users
func (r UserMysqlRepository) All(ctx context.Context) ([]user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	var usersSql []UserMysql
	if err := r.db.WithContext(ctx).Find(&usersSql).Error; err != nil {
		return nil, fmt.Errorf("[UserMysqlRepository:All][err:%w]", err)
	}

	users := make([]user.User, len(usersSql))
	for i := range usersSql {
		users[i] = usersSql[i].ToDomainModel()
	}

	return users, nil
}

// Save creates a new user
func (r UserMysqlRepository) Save(ctx context.Context, u user.User) error {
	defer metrics.TrackDBQueryDuration(time.Now())

	if err := r.db.WithContext(ctx).Create(&UserMysql{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return fmt.Errorf("[UserMysqlRepository:Save][err: %w]", err)
	}

	return nil
}

// FindByEmail search a user by email
func (r UserMysqlRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	var u UserMysql
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return user.User{}, fmt.Errorf("[UserMysqlRepository:FindByEmail][err: %w]", err)
	}

	return u.ToDomainModel(), nil
}
