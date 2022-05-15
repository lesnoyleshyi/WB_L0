package main

import (
	"L0/internal/domain"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const JSONTemplate = `{
  "order_uid": {{.Id}},
  "track_number": {{.TrackNum}},
  "entry": {{.Entry}},
  "delivery": {
    "name": {{.Delivery.Name}},
    "phone": {{.Delivery.Phone}},
    "zip": {{.Delivery.Zip}},
    "city": {{.Delivery.City}},
    "address": {{.Delivery.Address}},
    "region": {{.Delivery.Region}},
    "email": {{.Delivery.Email}}
  },
  "payment": {
    "transaction": {{.Payment.Transaction}},
    "request_id": {{.Payment.RequestId}},
    "currency": {{.Payment.Currency}},
    "provider": {{.Payment.Provider}},
    "amount": {{.Payment.Amount}},
    "payment_dt": {{.Payment.PaymentDt}},
    "bank": {{.Payment.Bank}},
    "delivery_cost": {{.Payment.DeliveryCost}},
    "goods_total": {{.Payment.GoodsTotal}},
    "custom_fee": {{.Payment.CustomFee}}
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": {{.Locale}},
  "internal_signature": {{.IntSign}},
  "customer_id": {{.CustomerId}},
  "delivery_service": {{.DeliveryService}},
  "shardkey": {{.ShardKey}},
  "sm_id": {{.SmId}},
  "date_created": {{with $time := .DateCreated.Format "2006-01-02T15:04:05Z07:00" }}{{ printf "\"%s\"" $time}}{{end}},
  "oof_shard": {{.OofShard}}
}`

var all []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
var capLetters []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers []rune = []rune("123456789")
var names []string = []string{"James", "Mary", "Robert", "Patricia", "John", "Jennifer",
	"Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan",
	"Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen", "Max"}
var surenames []string = []string{"Adams", "Wilson", "Burton", "Harris", "Stevens", "Robinson",
	"Lewis", "Walker", "Payne", "Baker"}
var cities []string = []string{"NY", "Los-Angeles", "Chicago", "Boston", "Washington", "Seattle"}
var streets []string = []string{"Court Street", "Creek Street", "Congress Street", "Dickson Street",
	"Lombard Street", "Larimer Square", "Chapel Street", "Second Street", "Ocean Drive"}
var currency []string = []string{"USD", "EUR", "RUB", "AUD", "CAD"}
var provider []string = []string{"wbpay", "applpay", "samspay", "nalik", "crypto"}
var bank []string = []string{"alpha", "sber", "tinek", "vtb", "abb"}
var locale []string = []string{"en", "ru", "gb", "uk", "can", "zimb"}
var delivery []string = []string{"usps", "rupost", "banderolka", "chinamail", "sdek", "dhl"}

const defaultMsgCount = 10
const defaultPause = 1000

func main() {
	var msgCount int = defaultMsgCount
	var sleepMs int = defaultPause
	var err error

	args := os.Args
	if len(args) > 2 {
		msgCount, err = strconv.Atoi(args[1])
		if err != nil || args[1] == "" {
			log.Println("default msg count", err)
		}
		sleepMs, err = strconv.Atoi(args[2])
		if err != nil || args[2] == "" {
			log.Println("default pause", err)
		}
	}

	stanClusterId := "test-cluster"
	clientId := "publisher"
	url := stan.NatsURL("nats://localhost:4222")

	conn, err := stan.Connect(stanClusterId, clientId, url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	for i := 0; i < msgCount; i++ {
		time.Sleep(time.Millisecond * time.Duration(sleepMs))
		data := generateJson()
		if err := conn.Publish("foo", []byte(data)); err != nil {
			log.Println("Something goes wrong:", err)
		}
	}
}

func generateJson() string {
	buf := new(strings.Builder)

	order := domain.Order{
		Id:       fmt.Sprintf("\"%s\"", randString(19, all)),
		TrackNum: fmt.Sprintf("\"%s\"", randString(14, capLetters)),
		Entry:    fmt.Sprintf("\"%s\"", randString(4, capLetters)),
		Delivery: domain.Delivery{
			Name:    fmt.Sprintf("\"%s %s\"", randChoice(names), randChoice(surenames)),
			Phone:   fmt.Sprintf("\"+%s\"", randString(11, numbers)),
			Zip:     fmt.Sprintf("\"%s\"", randString(6, numbers)),
			City:    fmt.Sprintf("\"%s\"", randChoice(cities)),
			Address: fmt.Sprintf("\"%s, %s, %d\"", randChoice(cities), randChoice(streets), rand.Intn(200)),
			Region:  fmt.Sprintf("\"%s\"", randChoice(cities)),
			Email:   fmt.Sprintf("\"%s%s@gmail.com\"", randChoice(names), randChoice(surenames)),
		},
		Payment: domain.Payment{
			Transaction:  fmt.Sprintf("\"%s\"", randString(19, all)),
			RequestId:    fmt.Sprintf("\"%s\"", randString(5, all)),
			Currency:     fmt.Sprintf("\"%s\"", randChoice(currency)),
			Provider:     fmt.Sprintf("\"%s\"", randChoice(provider)),
			Amount:       uint32(rand.Intn(1000)),
			PaymentDt:    uint64(rand.Intn(9999999999)),
			Bank:         fmt.Sprintf("\"%s\"", randChoice(bank)),
			DeliveryCost: uint32(rand.Intn(100000)),
			GoodsTotal:   uint32(rand.Intn(1000)),
			CustomFee:    uint32(rand.Intn(100)),
		},
		Items:           []domain.Item{},
		Locale:          fmt.Sprintf("\"%s\"", randChoice(locale)),
		IntSign:         fmt.Sprintf("\" \""),
		CustomerId:      fmt.Sprintf("\"%s\"", randString(10, all)),
		DeliveryService: fmt.Sprintf("\"%s\"", randChoice(delivery)),
		ShardKey:        fmt.Sprintf("\"%s\"", fmt.Sprint(rand.Intn(100))),
		SmId:            rand.Uint64(),
		DateCreated:     time.Now(),
		OofShard:        fmt.Sprintf("\"%s\"", fmt.Sprint(rand.Intn(10))),
	}
	templ, err := template.New("order").Parse(JSONTemplate)
	if err != nil {
		log.Println("error parsing template:", err)
		return ""
	}
	if err := templ.Execute(buf, order); err != nil {
		log.Println("error executing template:", err)
	}
	fmt.Println(order)
	return buf.String()
}

func randString(length int, vocabulary []rune) string {
	id := make([]rune, length)
	rand.Seed(time.Now().UnixMilli() / 10)
	for i := 0; i < length; i++ {
		id[i] = vocabulary[rand.Intn(len(vocabulary))]
	}
	return string(id)
}

func randChoice(source []string) string {
	rand.Seed(time.Now().UnixNano())
	return source[rand.Intn(len(source))]
}
