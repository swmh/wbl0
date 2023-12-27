package service

import "time"

//easyjson:json
type Order struct {
	OrderUID    string `json:"order_uid,required,required"`
	TrackNumber string `json:"track_number,required"`
	Entry       string `json:"entry,required"`
	Delivery    struct {
		Name    string `json:"name,required"`
		Phone   string `json:"phone,required"`
		Zip     string `json:"zip,required"`
		City    string `json:"city,required"`
		Address string `json:"address,required"`
		Region  string `json:"region,required"`
		Email   string `json:"email,required"`
	} `json:"delivery,required"`
	Payment struct {
		Transaction  string `json:"transaction,required"`
		RequestID    string `json:"request_id,required"`
		Currency     string `json:"currency,required"`
		Provider     string `json:"provider,required"`
		Amount       int    `json:"amount,required"`
		PaymentDt    int    `json:"payment_dt,required"`
		Bank         string `json:"bank,required"`
		DeliveryCost int    `json:"delivery_cost,required"`
		GoodsTotal   int    `json:"goods_total,required"`
		CustomFee    int    `json:"custom_fee,required"`
	} `json:"payment,required"`
	Items []struct {
		ChrtID      int    `json:"chrt_id,required"`
		TrackNumber string `json:"track_number,required"`
		Price       int    `json:"price,required"`
		Rid         string `json:"rid,required"`
		Name        string `json:"name,required"`
		Sale        int    `json:"sale,required"`
		Size        string `json:"size,required"`
		TotalPrice  int    `json:"total_price,required"`
		NmID        int    `json:"nm_id,required"`
		Brand       string `json:"brand,required"`
		Status      int    `json:"status,required"`
	} `json:"items,required"`
	Locale            string    `json:"locale,required"`
	InternalSignature string    `json:"internal_signature,required"`
	CustomerID        string    `json:"customer_id,required"`
	DeliveryService   string    `json:"delivery_service,required"`
	Shardkey          string    `json:"shardkey,required"`
	SmID              int       `json:"sm_id,required"`
	DateCreated       time.Time `json:"date_created,required"`
	OofShard          string    `json:"oof_shard,required"`
}
