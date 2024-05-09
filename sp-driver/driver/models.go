package driver

import (
	"fmt"

	"github.com/cockroachdb/apd/v3"
	"github.com/cycle-labs/example-step-driver/sp-driver/api"
)

type Transfer struct {
	FromAccountID int
	ToAccountID   int
	Amount        apd.Decimal
}

func FromMonetaryValue(v api.MonetaryValue) (*apd.Decimal, error) {
	decimal, _, err := apd.NewFromString(v)
	if err != nil {
		return nil, fmt.Errorf("failed to parse monetary value: %w", err)
	}
	return decimal, nil
}
