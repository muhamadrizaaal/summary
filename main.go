package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    _ "github.com/lib/pq"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/robfig/cron/v3"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := os.Getenv("DB_CONN_STRING")
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
}

func summarizeData() {
    query := `
    INSERT INTO summary (total_sales, total_customers, total_products, summary_date)
    SELECT 
        (SELECT SUM(amount) FROM sales) AS total_sales,
        (SELECT COUNT(DISTINCT id) FROM customers) AS total_customers,
        (SELECT COUNT(DISTINCT id) FROM products) AS total_products,
        NOW() AS summary_date;
    `
    _, err := db.Exec(query)
    if err != nil {
        log.Printf("Failed to summarize data: %v", err)
    } else {
        log.Println("Data summarized successfully")
    }
}

func startScheduler() {
    c := cron.New()
    c.AddFunc("@hourly", summarizeData)
    c.Start()
}

func getSummary(c echo.Context) error {
    start := time.Now()
    rows, err := db.Query("SELECT total_sales, total_customers, total_products, summary_date FROM summary ORDER BY summary_date DESC LIMIT 1")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    defer rows.Close()

    var summary struct {
        TotalSales    float64   `json:"total_sales"`
        TotalCustomers int      `json:"total_customers"`
        TotalProducts  int      `json:"total_products"`
        SummaryDate   time.Time `json:"summary_date"`
    }

    if rows.Next() {
        if err := rows.Scan(&summary.TotalSales, &summary.TotalCustomers, &summary.TotalProducts, &summary.SummaryDate); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }
    }

    elapsed := time.Since(start)
    log.Printf("Query completed in %s", elapsed)

    return c.JSON(http.StatusOK, summary)
}

func main() {
    initDB()
    startScheduler()

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/summary", getSummary)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
