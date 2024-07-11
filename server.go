package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type CreateIntentResponse struct {
  ClientSecret string `json:"client_secret"`
}

func CreateIntent(c echo.Context) error {
  params := &stripe.PaymentIntentParams{
    Amount: stripe.Int64(1099),
    Currency: stripe.String(string(stripe.CurrencyUSD)),
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
