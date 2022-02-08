package server

import (
	"context"

	"github.com/snxq/snxq.cc/internal/model"
)

// Dao data accept object
type Dao interface {
	ArticleGet(context.Context, map[string]interface{}) (
		*model.Article, error)
	ArticleQuery(context.Context, map[string]interface{}) (
		*[]*model.Article, error)
}
