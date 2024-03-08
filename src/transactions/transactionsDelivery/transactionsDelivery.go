package transactionsDelivery

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/transactionsDto"
	"BookingRoom/pkg/middleware"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/transactions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		transactionsGroup.GET("", middleware.JWTAuth("ADMIN"), handler.RetrieveAllTransactions)
		transactionsGroup.GET("/:id", middleware.JWTAuth("ADMIN"), handler.RetrieveTransactionsByID)
		transactionsGroup.POST("", middleware.JWTAuth("ADMIN"), handler.CreateTransactions)
	}

	transactionsGroupLogs := transactionsGroup.Group("/logs")
	{
		transactionsGroupLogs.POST("/:id", middleware.JWTAuth("GA"), handler.UpdateTransactions)
		transactionsGroupLogs.GET("/:id", middleware.JWTAuth("GA"), handler.RetrieveTrxLogBydID)
		transactionsGroupLogs.GET("", middleware.JWTAuth("GA"), handler.RetrieveAllTrxLogs)
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

		if err.Error() == "03" {
			json.NewResponseBadRequest(ctx, nil, "Employee ID not found on our records", "03", "03")
			return
		}

		if err.Error() == "04" {
			json.NewResponseBadRequest(ctx, nil, "Rooms ID not found on our records", "03", "03")
			return
		}

		if err.Error() == "05" {
			json.NewResponseBadRequest(ctx, nil, "Rooms are not available right now", "03", "03")
			return
		}

		json.NewResponseError(ctx, err.Error(), "03", "03")
		return
	}

	json.NewResponseCreatedSuccess(ctx, "Transactions created successfully", "03", "03")
}

func (c transactionsDelivery) UpdateTransactions(ctx *gin.Context) {
	var req transactionsDto.TransactionLogRequest
	trxIDStr := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "04")
		return
	}

	trxID, err := uuid.Parse(trxIDStr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "04")
		return
	}

	trxData := transactionsDto.TransactionLog{
		ApprovedBy:       req.ApprovedBy,
		ApprovalStatus:   req.ApprovalStatus,
		Descriptions:     req.Descriptions,
		TransactionLogID: trxID,
	}

	validationError := utils.ValidationUpdateTrxReq(trxData)
	if len(validationError) > 0 {
		json.NewResponseBadRequest(ctx, validationError, "failed", "03", "04")
		return
	}

	err = c.transactionsUC.UpdateTrxLog(trxData)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "Employee ID not found on our records", "03", "04")
			return
		}

		if err.Error() == "02" {
			json.NewResponseBadRequest(ctx, nil, "Transaction logs ID not found on our records", "03", "04")
			return
		}

		json.NewResponseError(ctx, err.Error(), "03", "04")
		return
	}

	json.NewResponseSuccess(ctx, "Transaction log updated successfully", nil, "success", "03", "04")
}

func (c transactionsDelivery) RetrieveTrxLogBydID(ctx *gin.Context) {
	trxLogID := ctx.Param("id")

	data, err := c.transactionsUC.RetrieveTrxLogByID(trxLogID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "Transaction Logs ID not found on our records", "03", "05")
			return
		}
		json.NewResponseError(ctx, err.Error(), "03", "05")
		return
	}

	json.NewResponseSuccess(ctx, data, nil, "success", "03", "05")
}

func (c transactionsDelivery) RetrieveAllTrxLogs(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("size"))
	startDate, _ := time.Parse("2006-01-02", ctx.Query("startDate"))
	endDate, _ := time.Parse("2006-01-02", ctx.Query("endDate"))

	startDateFormatted := startDate.Format("2006-01-02")
	endDateFormatted := endDate.Format("2006-01-02")

	data, pagination, err := c.transactionsUC.RetrieveAllTrxLog(page, pageSize, startDateFormatted, endDateFormatted)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "06")
		return
	}

	json.NewResponseSuccess(ctx, data, pagination, "success", "03", "06")
}
