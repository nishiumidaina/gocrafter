package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gocrafter/api/game/di"
	"github.com/game-core/gocrafter/api/game/presentation/server/account"
	"github.com/game-core/gocrafter/api/game/presentation/server/friend"
	"github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
)

func Router(lis net.Listener) {
	// DI
	accountHandler := di.InitializeAccountHandler()
	friendHandler := di.InitializeFriendHandler()
	authInterceptor := di.InitializeAuthInterceptor()
	loginBonusHandler := di.InitializeLoginBonusHandler()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	account.RegisterAccountServer(s, accountHandler)
	friend.RegisterFriendServer(s, friendHandler)
	loginBonus.RegisterLoginBonusServer(s, loginBonusHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
