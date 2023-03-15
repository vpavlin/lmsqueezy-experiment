package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/sirupsen/logrus"
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
	CustomData CustomData `json:"customer_data"`
}

type Attributes struct {
	Identifier string `json:"identifier"`
	UserEmail  string `json:"user_email"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
}

type Data struct {
	Attributes Attributes
}

func main() {
	pb := pocketbase.NewWithConfig(&pocketbase.Config{})

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

				record, err := pb.App.Dao().FindAuthRecordByUsername("users", metadata.CustomData.CustomerId)
				if err != nil {
					logrus.Errorf("Failed to find user: %s", err)
					return nil
				}

				logrus.Info("Payment %s confirmed (status: %s) for user %s: %", data.Attributes.Identifier, data.Attributes.Status, record.Email())
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
