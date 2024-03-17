// Package userIdleBonus ユーザー放置ボーナス
package userIdleBonus

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetUserIdleBonus(userId string, masterIdleBonusId int64, receivedAt *timestamppb.Timestamp) *UserIdleBonus {
	return &UserIdleBonus{
		UserId:            userId,
		MasterIdleBonusId: masterIdleBonusId,
		ReceivedAt:        receivedAt,
	}
}
