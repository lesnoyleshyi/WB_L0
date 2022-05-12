package domain

import "time"

type Order struct {
	Id              string    `json:"order_uid"`
	TrackNum        string    `json:"track_number"`
	Entry           string    `json:"entry"`
	Delivery        Delivery  `json:"delivery"`
	Payment         Payment   `json:"payment"`
	Items           []Item    `json:"items"`
	Locale          string    `json:"locale"`
	IntSign         string    `json:"internal_signature"`
	CustomerId      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	ShardKey        string    `json:"shardkey"`
	SmId            uint64    `json:"sm_id"`
	DateCreated     time.Time `json:"date_created"`
	OofShard        string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       uint32 `json:"amount"`
	PaymentDt    uint64 `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint32 `json:"delivery_cost"`
	GoodsTotal   uint32 `json:"goods_total"`
	CustomFee    uint32 `json:"custom_fee"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
