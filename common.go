package gosend

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	MethodGet                  = "GET"
	MethodPost                 = "POST"
	MethodPut                  = "PUT"
	URLCalculatePrice          = "gokilat/v10/calculate/price"
	URLBooking                 = "gokilat/v10/booking"
	URLTraceGosendOrderNo      = "gokilat/v10/booking/orderno/"
	URLTraceGosendStoreOrderID = "gokilat/v10/booking/storeOrderId/"
	URLCancel                  = "gokilat/v10/booking/cancel"
	TypeGoSendOrderNumber      = "OrderNumber"
	TypeStoreOrderID           = "OrderID"
)

var (
	Host     string
	ClientID string
	PassKey  string
)

func init() {
	_ = godotenv.Load()

	Host = os.Getenv("GOSEND_HOST")
	ClientID = os.Getenv("GOSEND_CLIENT_ID")
	PassKey = os.Getenv("GOSEND_PASS_KEY")
}

func InitDB() (db *sql.DB) {
	conn := os.Getenv("DB_CONNECTION")

	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	return
}
