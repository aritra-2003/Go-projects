package events

import (
	"net/http"

	"example.com/financetracker/database"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
	
)

func CreateTransaction(c *gin.Context) {
	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if tx.UserID == 0 || tx.Amount == 0 || tx.Category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields: user_id, amount, or category"})
		return
	}

	query := `
		INSERT INTO transactions (user_id, amount, category, description, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	row := database.DB.QueryRow(query, tx.UserID, tx.Amount, tx.Category, tx.Description, tx.Date)

	if err := row.Scan(&tx.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func GetTransaction(c* gin.Context) {
	var transactions [] models.Transaction
	query:=`SELECT id,user_id,amount,category,description,date from transactions order by date desc`
	rows,err:=database.DB.Query(query)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to fetch transaction"})
		return
	}
	defer rows.Close()
	for rows.Next(){
		var transaction models.Transaction
	
		if err:=rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Category, &transaction.Description, &transaction.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transactions"})
			return
	}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}
func DeleteTransaction(c *gin.Context){
	transactionId:=c.Param("id")
	_,err:=database.DB.Exec("DELETE FROM transactions where id =$1",transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return                         
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func GetTransactionByMonth(c *gin.Context){
	month:=c.Query("month")
	if month==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required"})
		return
	}
	var transactions []models.Transaction
	query:="SELECT id,user_id,amount,category,description,date from transactions where to char(date,'YYYY-MM')=$1 order by date desc"
    rows,err:=database.DB.Query(query)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})

	}
	defer rows.Close()
	for rows.Next(){
		var transaction models.Transaction
	
		if err:=rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Category, &transaction.Description, &transaction.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transactions"})
			return
	}
			transactions = append(transactions, transaction)
		}
		c.JSON(http.StatusOK, transactions)
	}
	func UpdateTransaction(c *gin.Context) {
		var tx models.Transaction
		if err := c.ShouldBindJSON(&tx); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
	
		transactionId := c.Param("id")
		query := `
			UPDATE transactions
			SET amount = $1, category = $2, description = $3, date = $4
			WHERE id = $5
		`
		_, err := database.DB.Exec(query, tx.Amount, tx.Category, tx.Description, tx.Date, transactionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
	}


func GetTransactionByCategory(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category parameter is required"})
		return
	}

	var transactions []models.Transaction
	query := `SELECT id, user_id, amount, category, description, date FROM transactions WHERE category = $1 ORDER BY date DESC`
	rows, err := database.DB.Query(query, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Category, &transaction.Description, &transaction.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transactions"})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}


func GetMonthlySummary(c *gin.Context) {
	month := c.Query("month")
	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required"})
		return
	}

	type Summary struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}

	var summaries []Summary
	query := `SELECT category, SUM(amount) FROM transactions WHERE TO_CHAR(date, 'YYYY-MM') = $1 GROUP BY category`
	rows, err := database.DB.Query(query, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monthly summary"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var summary Summary
		if err := rows.Scan(&summary.Category, &summary.Total); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse summary"})
			return
		}
		summaries = append(summaries, summary)
	}

	c.JSON(http.StatusOK, summaries)
}

func GetMonthlyLimitStatus(c *gin.Context) {
	month := c.Query("month")
	limit := 10000.0 

	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required"})
		return
	}

	var totalSpent float64
	query := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE TO_CHAR(date, 'YYYY-MM') = $1`
	err := database.DB.QueryRow(query, month).Scan(&totalSpent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expense data"})
		return
	}

	remaining := limit - totalSpent
	status := "Within Limit"
	if remaining < 0 {
		status = "Over Limit"
	}

	c.JSON(http.StatusOK, gin.H{
		"total_spent": totalSpent,
		"remaining":   remaining,
		"status":      status,
	})
}
