package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewUserRepository(db *gorm.DB, cache *redis.Client) *UserRepository {
	return &UserRepository{
		db:    db,
		cache: cache,
	}
}

func (r *UserRepository) FindAll(pagination *model.UserPagination) error {
	err := GetCache(r.cache, "users", pagination)
	if err != nil {
		if err != redis.Nil {
			return err
		}

		err = r.db.Scopes(Pagination(&pagination.Items, &pagination.Meta, r.db)).Find(&pagination.Items).Count(&pagination.Meta.ItemCount).Error
		if err != nil {
			return err
		}

		return SaveCache(r.cache, "users", pagination)
	}

	return nil
}

func (r *UserRepository) FindOne(id uint, user *model.User) error {
	return r.db.Preload(clause.Associations).First(&user, id).Error
}

func (r *UserRepository) FindMulti(ids []uint32, users *[]*model.User) error {
	err := GetCache(r.cache, "users", users)
	if err != nil {
		if err != redis.Nil {
			return err
		}

		err = r.db.Where("id IN ?", ids).Find(&users).Error
		if err != nil {
			return err
		}

		return SaveCache(r.cache, "users", users)
	}

	return nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Update(id uint, user *model.User) error {
	return r.db.Where(id).Updates(&user).First(&user).Error
}

func (r *UserRepository) Delete(id uint, user *model.User) error {
	return r.db.First(&user, id).Delete(&model.User{}).Error
}
