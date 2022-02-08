package dao

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/snxq/snxq.cc/internal/configs"
	"github.com/snxq/snxq.cc/internal/model"
	"github.com/snxq/snxq.cc/internal/server"
)

// Dao impl server.Dao
type Dao struct {
	db *gorm.DB
}

// check impl
var _ server.Dao = (*Dao)(nil)

// New return an real dao(data access object)
func New(cfg configs.MySQL) (*Dao, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN), nil)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)

	if err := db.AutoMigrate(&model.Article{},
		&model.People{}, &model.Tag{}); err != nil {

		return nil, err
	}

	return &Dao{db}, nil
}

// ArticleGet return article object by
func (d *Dao) ArticleGet(ctx context.Context, query map[string]interface{}) (
	*model.Article, error) {

	article := &model.Article{}
	if err := d.db.WithContext(ctx).Where(query).First(article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

// ArticleQuery return a list of article
func (d *Dao) ArticleQuery(ctx context.Context, query map[string]interface{}) (
	*[]*model.Article, error) {

	articles := &[]*model.Article{}
	if err := d.db.WithContext(ctx).Where(query).Find(articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
