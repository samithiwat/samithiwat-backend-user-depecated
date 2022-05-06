package repository

import (
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(meta *proto.PaginationMetadata, users *[]*model.User) error {
	return r.db.Scopes(Pagination(users, meta)).Find(&users).Count(&meta.ItemCount).Error
}

func (r *UserRepository) FindOne(id uint, user *model.User) error {
	return r.db.Preload(clause.Associations).First(&user, id).Error
}

func (r *UserRepository) FindMulti(ids []uint32, users *[]*model.User) error {
	return r.db.Where("id IN ?", ids).Find(&users).Error
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
