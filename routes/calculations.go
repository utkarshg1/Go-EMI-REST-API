package routes

import (
	"net/http"

	"example.com/emi-rest-api/models"
	"github.com/gin-gonic/gin"
)

func calculateLoanEMI(context *gin.Context) {
	var input *models.LoanDetails

	err := context.ShouldBindJSON(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	loan, err := models.New(input.Principal, input.Period, input.RateOfInterest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data entered",
			"error":   err.Error(),
		})
		return
	}

	calcs := loan.GetEMICalcs()

	context.JSON(http.StatusOK, gin.H{
		"message": "EMI Calculations successful",
		"data":    calcs,
	})
}
