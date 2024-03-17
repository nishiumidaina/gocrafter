// Package masterAction アクション
//
//go:generate mockgen -source=./master_action_repository.gen.go -destination=./master_action_repository_mock.gen.go -package=masterAction
package masterAction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActionRepository interface {
	Find(ctx context.Context, id int64) (*MasterAction, error)
	FindOrNil(ctx context.Context, id int64) (*MasterAction, error)
	FindByName(ctx context.Context, name string) (*MasterAction, error)
	FindByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterAction, error)
	FinOrNilByName(ctx context.Context, name string) (*MasterAction, error)
	FinOrNilByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterAction, error)
	FindList(ctx context.Context) (MasterActions, error)
	FindListByName(ctx context.Context, name string) (MasterActions, error)
	FindListByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (MasterActions, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterAction) error
}
