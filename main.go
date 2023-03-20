package main

import (
	"encoding/json"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/sirupsen/logrus"

	_ "github.com/vpavlin/lmsqueezy-experiment/migrations"
)

type Payload struct {
	Meta Metadata `json:"meta"`
	Data Data     `json:"data"`
}
type CustomData struct {
	CustomerId string `json:"customer_id"`
}
type Metadata struct {
	EventName  string     `json:"event_name"`
	CustomData CustomData `json:"custom_data"`
}

type Attributes struct {
	Identifier string  `json:"identifier"`
	UserEmail  string  `json:"user_email"`
	Total      float64 `json:"total"`
	Status     string  `json:"status"`
}

type Data struct {
	Attributes Attributes
}

type Order struct {
	NFTId int `json:"nft_id"`
}
type OrderCreated struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

func main() {
	pb := pocketbase.NewWithConfig(&pocketbase.Config{})
	migratecmd.MustRegister(pb, pb.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/lms/webhook", func(c echo.Context) error {
			event := c.Request().Header.Get("X-Event-Name")
			var payload Payload
			err := c.Bind(&payload)
			if err != nil {
				logrus.Errorf("Failed to bind payload: %s", err)
				return nil
			}

			switch event {
			case "order_created":
				logrus.Infof("ORDER CREATED event: ")

				metadata := payload.Meta
				data := payload.Data

				d, _ := json.MarshalIndent(payload, "", "  ")
				logrus.Info(string(d))

				record, err := pb.App.Dao().FindAuthRecordByUsername("users", metadata.CustomData.CustomerId)
				if err != nil {
					logrus.Errorf("Failed to find user: %s", err)
					return nil
				}

				logrus.Infof("Payment %s confirmed (status: %s) for user %s: %", data.Attributes.Identifier, data.Attributes.Status, record.Email())
			default:
				return nil
			}

			return nil

		})

		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
