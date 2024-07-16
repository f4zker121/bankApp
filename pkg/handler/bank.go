package handler

import (
	bankapp "bankApp"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Request struct {
	UserId int
	Amount float64
	Result chan<- Result
}

type Result struct {
	Balance float64
	Err     error
}

func (h *Handler) createAcc(c *gin.Context) {
	start := time.Now()
	id, err := h.accounts.BankAccount.CreateUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), logrus.Fields{
			"duration": time.Since(start),
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"duration": time.Since(start),
		"userId":   id,
	}).Info("createAcc successful")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deposit(c *gin.Context) {
	start := time.Now()
	idParam := c.Param("id")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id", logrus.Fields{
			"duration": time.Since(start),
		})
		return
	}

	var input bankapp.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), logrus.Fields{
			"duration": time.Since(start),
			"userId":   userId,
		})
		return
	}

	resultChan := make(chan Result)
	defer close(resultChan)
	go func() {
		err := h.accounts.BankAccount.Deposit(userId, input.Balance)
		resultChan <- Result{Err: err}
	}()

	result := <-resultChan
	if result.Err != nil {
		newErrorResponse(c, http.StatusInternalServerError, result.Err.Error(), logrus.Fields{
			"duration": time.Since(start),
			"userId":   userId,
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"duration": time.Since(start),
		"userId":   userId,
	}).Info("deposit successful")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) withdraw(c *gin.Context) {
	start := time.Now()
	idParam := c.Param("id")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id", logrus.Fields{
			"duration": time.Since(start),
		})
		return
	}

	var input bankapp.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), logrus.Fields{
			"duration": time.Since(start),
			"userId":   userId,
		})
		return
	}

	resultChan := make(chan Result)
	defer close(resultChan)
	go func() {
		err := h.accounts.BankAccount.Withdraw(userId, input.Balance)
		resultChan <- Result{Err: err}
	}()

	result := <-resultChan
	if result.Err != nil {
		newErrorResponse(c, http.StatusInternalServerError, result.Err.Error(), logrus.Fields{
			"duration": time.Since(start),
			"userId":   userId,
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"duration": time.Since(start),
		"userId":   userId,
	}).Info("withdraw successful")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getBalance(c *gin.Context) {
	start := time.Now()
	idParam := c.Param("id")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id", logrus.Fields{
			"duration": time.Since(start),
		})
		return
	}

	resultChan := make(chan Result)
	defer close(resultChan)
	go func() {
		balance, err := h.accounts.BankAccount.GetBalance(userId)
		resultChan <- Result{Balance: balance, Err: err}
	}()

	result := <-resultChan
	if result.Err != nil {
		newErrorResponse(c, http.StatusBadRequest, result.Err.Error(), logrus.Fields{
			"duration": time.Since(start),
			"userId":   userId,
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"duration": time.Since(start),
		"userId":   userId,
	}).Info("getBalance successful")

	c.JSON(http.StatusOK, result.Balance)
}
