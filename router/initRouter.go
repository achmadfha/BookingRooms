package router

import (
	"BookingRoom/src/transactions/transactionsDelivery"
	"BookingRoom/src/transactions/transactionsRepository"
	"BookingRoom/src/transactions/transactionsUseCase"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	transactionsRepo := transactionsRepository.NewTransactionsRepository(db)

	// usecase
	transactionUC := transactionsUseCase.NewTransactionsUseCase(transactionsRepo)

	// delivery
	transactionsDelivery.NewTransactionsDelivery(v1Group, transactionUC)
}
