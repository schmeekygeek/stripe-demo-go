package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72/refund"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type CreateIntentResponse struct {
  ClientSecret string `json:"client_secret"`
}

func CreateIntent(c echo.Context) error {
  params := &stripe.PaymentIntentParams{
    Amount: stripe.Int64(500),
    Currency: stripe.String(string(stripe.CurrencyINR)),
  };
  result, err := paymentintent.New(params);
  if err != nil {
    return err
  }
  return c.JSON(
    http.StatusOK,
    CreateIntentResponse{ ClientSecret: result.ClientSecret },
  )
}

func ListIntents(c echo.Context) error {
  iter := paymentintent.List(&stripe.PaymentIntentListParams{ })
  return c.JSON(200, iter.List())
}

func CaptureIntent(c echo.Context) error {
  intentId := c.Param("id")
  log.Println(intentId)
  params := &stripe.PaymentIntentCaptureParams{}
  result, err := paymentintent.Capture(intentId, params)
  if err != nil {
    log.Println(err)
  }
  return c.JSON(http.StatusOK, result)
}

func RefundIntent(c echo.Context) error {
  params := &stripe.RefundParams{Charge: stripe.String(c.Param("id"))};
  result, err := refund.New(params);
  if err != nil {
    log.Println(err)
  }
  return c.JSON(200, result)
}
