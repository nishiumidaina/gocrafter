// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/api/game/presentation/handler/account"
	"github.com/game-core/gocrafter/api/game/presentation/handler/friend"
	"github.com/game-core/gocrafter/api/game/presentation/handler/loginBonus"
	"github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	account2 "github.com/game-core/gocrafter/api/game/usecase/account"
	friend2 "github.com/game-core/gocrafter/api/game/usecase/friend"
	loginBonus2 "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	"github.com/game-core/gocrafter/configs/database"
	account3 "github.com/game-core/gocrafter/pkg/domain/model/account"
	friend3 "github.com/game-core/gocrafter/pkg/domain/model/friend"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonus3 "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userFriend"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userItemBox"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userLoginBonus"
	masterTransaction2 "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
)

// Injectors from wire.go:

func InitializeAuthInterceptor() auth.AuthInterceptor {
	authInterceptor := auth.NewAuthInterceptor()
	return authInterceptor
}

func InitializeAccountHandler() account.AccountHandler {
	accountUsecase := InitializeAccountUsecase()
	accountHandler := account.NewAccountHandler(accountUsecase)
	return accountHandler
}

func InitializeFriendHandler() friend.FriendHandler {
	friendUsecase := InitializeFriendUsecase()
	friendHandler := friend.NewFriendHandler(friendUsecase)
	return friendHandler
}

func InitializeLoginBonusHandler() loginBonus.LoginBonusHandler {
	loginBonusUsecase := InitializeLoginBonusUsecase()
	loginBonusHandler := loginBonus.NewLoginBonusHandler(loginBonusUsecase)
	return loginBonusHandler
}

func InitializeAccountUsecase() account2.AccountUsecase {
	accountService := InitializeAccountService()
	transactionService := InitializeTransactionService()
	accountUsecase := account2.NewAccountUsecase(accountService, transactionService)
	return accountUsecase
}

func InitializeFriendUsecase() friend2.FriendUsecase {
	friendService := InitializeFriendService()
	transactionService := InitializeTransactionService()
	friendUsecase := friend2.NewFriendUsecase(friendService, transactionService)
	return friendUsecase
}

func InitializeLoginBonusUsecase() loginBonus2.LoginBonusUsecase {
	loginBonusService := InitializeLoginBonusService()
	transactionService := InitializeTransactionService()
	loginBonusUsecase := loginBonus2.NewLoginBonusUsecase(loginBonusService, transactionService)
	return loginBonusUsecase
}

func InitializeAccountService() account3.AccountService {
	shardService := InitializeShardService()
	sqlHandler := database.NewDB()
	userAccountRepository := userAccount.NewUserAccountDao(sqlHandler)
	accountService := account3.NewAccountService(shardService, userAccountRepository)
	return accountService
}

func InitializeFriendService() friend3.FriendService {
	accountService := InitializeAccountService()
	sqlHandler := database.NewDB()
	userFriendRepository := userFriend.NewUserFriendDao(sqlHandler)
	friendService := friend3.NewFriendService(accountService, userFriendRepository)
	return friendService
}

func InitializeLoginBonusService() loginBonus3.LoginBonusService {
	itemService := InitializeItemService()
	sqlHandler := database.NewDB()
	userLoginBonusRepository := userLoginBonus.NewUserLoginBonusDao(sqlHandler)
	masterLoginBonusRepository := masterLoginBonus.NewMasterLoginBonusDao(sqlHandler)
	masterLoginBonusEventRepository := masterLoginBonusEvent.NewMasterLoginBonusEventDao(sqlHandler)
	masterLoginBonusItemRepository := masterLoginBonusItem.NewMasterLoginBonusItemDao(sqlHandler)
	masterLoginBonusScheduleRepository := masterLoginBonusSchedule.NewMasterLoginBonusScheduleDao(sqlHandler)
	loginBonusService := loginBonus3.NewLoginBonusService(itemService, userLoginBonusRepository, masterLoginBonusRepository, masterLoginBonusEventRepository, masterLoginBonusItemRepository, masterLoginBonusScheduleRepository)
	return loginBonusService
}

func InitializeItemService() item.ItemService {
	sqlHandler := database.NewDB()
	userItemBoxRepository := userItemBox.NewUserItemBoxDao(sqlHandler)
	masterItemRepository := masterItem.NewMasterItemDao(sqlHandler)
	itemService := item.NewItemService(userItemBoxRepository, masterItemRepository)
	return itemService
}

func InitializeShardService() shard.ShardService {
	sqlHandler := database.NewDB()
	commonShardRepository := commonShard.NewCommonShardDao(sqlHandler)
	shardService := shard.NewShardService(commonShardRepository)
	return shardService
}

func InitializeTransactionService() transaction.TransactionService {
	sqlHandler := database.NewDB()
	commonTransactionRepository := commonTransaction.NewCommonTransactionDao(sqlHandler)
	masterTransactionRepository := masterTransaction.NewMasterTransactionDao(sqlHandler)
	userTransactionRepository := masterTransaction2.NewUserTransactionDao(sqlHandler)
	transactionService := transaction.NewTransactionService(commonTransactionRepository, masterTransactionRepository, userTransactionRepository)
	return transactionService
}
