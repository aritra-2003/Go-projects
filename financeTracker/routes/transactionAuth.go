package routes
 


import (
"example.com/financetracker/events"
	"github.com/gin-gonic/gin"
	
)


func TransactionRoutes(r *gin.Engine) {
	r.POST("/transactions", events.CreateTransaction)
	r.GET("/transactions", events.GetTransaction)
    r.POST("/transactions/:id",events.DeleteTransaction)
	r.GET("/transactions/?month=YYYY-MM",events.GetTransactionByMonth)
	r.PUT("transactions/{id}",events.UpdateTransaction)
	r.GET("/transactions/category", events.GetTransactionByCategory) // Fetch transactions by category (query param: ?category=food)
	r.GET("/transactions/summary", events.GetMonthlySummary) // Get monthly summary (query param: ?month=YYYY-MM)
	r.GET("/transactions/limit", events.GetMonthlyLimitStatus) 
	

	
}
