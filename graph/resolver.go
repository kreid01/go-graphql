package graph

import (
	"kreid.com/graphl-go/graph/model"
	"github.com/go-pg/pg/v10"
)

type Resolver struct{
	DB *pg.DB
	ChatMessages []*model.Message
	ChatObservers map[string]chan []*model.Message	
}
