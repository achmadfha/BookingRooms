package transactionsDelivery

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/transactionsDto"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/transactions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type transactionsDelivery struct {
	transactionsUC transactions.TransactionsUseCase
}

func NewTransactionsDelivery(v1Group *gin.RouterGroup, transactionsUC transactions.TransactionsUseCase) {
	handler := transactionsDelivery{
		transactionsUC: transactionsUC,
	}

	transactionsGroup := v1Group.Group("/transactions")
	{
		transactionsGroup.GET("", handler.RetrieveAllTransactions)
		transactionsGroup.GET("/:id", handler.RetrieveTransactionsByID)
		transactionsGroup.POST("", handler.CreateTransactions)
	}
}

func (c transactionsDelivery) RetrieveAllTransactions(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("size"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("startDate"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("endDate"))

	startDateFormatted := startDate.Format("2006-01-02")
	endDateFormatted := endDate.Format("2006-01-02")

	transactionData, pagination, err := c.transactionsUC.RetrieveAllTransactions(page, pageSize, startDateFormatted, endDateFormatted)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "01")
		return
	}

	json.NewResponseSuccess(ctx, transactionData, pagination, "success", "03", "01")
}

func (c transactionsDelivery) RetrieveTransactionsByID(ctx *gin.Context) {
	trxID := ctx.Param("id")

	data, err := c.transactionsUC.RetrieveTransactionsByID(trxID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "Transaction ID not found on our records", "03", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "03", "02")
		return
	}

	json.NewResponseSuccess(ctx, data, nil, "success", "03", "02")
}

func (c transactionsDelivery) CreateTransactions(ctx *gin.Context) {
	var req transactionsDto.TransactionsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "03")
		return
	}

	validationError := utils.ValidationTrxReq(req)
	if len(validationError) > 0 {
		json.NewResponseBadRequest(ctx, validationError, "failed", "03", "03")
		return
	}

	err := c.transactionsUC.CreateTransactions(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseError(ctx, "err", "03", "03")
			return
		}

		if err.Error() == "02" {
			json.NewResponseError(ctx, "err", "03", "03")
			return
		}

		json.NewResponseError(ctx, err.Error(), "03", "03")
		return
	}

	json.NewResponseCreatedSuccess(ctx, "Transactions created successfully", "03", "03")
}
