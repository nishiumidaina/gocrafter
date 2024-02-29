package friend

import (
	"context"

	friendServer "github.com/game-core/gocrafter/api/game/presentation/server/friend"
	"github.com/game-core/gocrafter/internal/errors"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type FriendUsecase interface {
	Send(ctx context.Context, req *friendServer.FriendSendRequest) (*friendServer.FriendSendResponse, error)
	Approve(ctx context.Context, req *friendServer.FriendApproveRequest) (*friendServer.FriendApproveResponse, error)
}

type friendUsecase struct {
	friendService      friendService.FriendService
	transactionService transactionService.TransactionService
}

func NewFriendUsecase(
	friendService friendService.FriendService,
	transactionService transactionService.TransactionService,
) FriendUsecase {
	return &friendUsecase{
		friendService:      friendService,
		transactionService: transactionService,
	}
}

// Send フレンド申請を送信する
func (s *friendUsecase) Send(ctx context.Context, req *friendServer.FriendSendRequest) (*friendServer.FriendSendResponse, error) {
	// transaction
	txs, err := s.transactionService.MultiUserBegin(ctx, []string{req.UserId, req.FriendUserId})
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.MultiUserBegin", err)
	}
	defer func() {
		s.transactionService.MultiUserEnd(ctx, txs, err)
	}()

	result, err := s.friendService.Send(ctx, txs, friendService.SetFriendSendRequest(req.UserId, req.FriendUserId))
	if err != nil {
		return nil, errors.NewMethodError("s.friendService.Send", err)
	}

	return friendServer.SetFriendSendResponse(
		friendServer.SetUserFriend(
			result.UserFriend.UserId,
			result.UserFriend.FriendUserId,
			friendServer.FriendType(result.UserFriend.FriendType),
		),
	), nil
}

// Approve フレンド申請を承諾する
func (s *friendUsecase) Approve(ctx context.Context, req *friendServer.FriendApproveRequest) (*friendServer.FriendApproveResponse, error) {
	// transaction
	txs, err := s.transactionService.MultiUserBegin(ctx, []string{req.UserId, req.FriendUserId})
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.MultiUserBegin", err)
	}
	defer func() {
		s.transactionService.MultiUserEnd(ctx, txs, err)
	}()

	result, err := s.friendService.Approve(ctx, txs, friendService.SetFriendApproveRequest(req.UserId, req.FriendUserId))
	if err != nil {
		return nil, errors.NewMethodError("s.friendService.Approve", err)
	}

	return friendServer.SetFriendApproveResponse(
		friendServer.SetUserFriend(
			result.UserFriend.UserId,
			result.UserFriend.FriendUserId,
			friendServer.FriendType(result.UserFriend.FriendType),
		),
	), nil
}
