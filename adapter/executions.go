package adapter

import "time"

type BitFlyerExecution struct {
	Params struct {
		Message []struct {
			ID                         int64     `json:"id"`
			Side                       string    `json:"side"`
			Price                      float64   `json:"price"`
			Size                       float64   `json:"size"`
			ExecDate                   time.Time `json:"exec_date"`
			BuyChildOrderAcceptanceID  string    `json:"buy_child_order_acceptance_id"`
			SellChildOrderAcceptanceID string    `json:"sell_child_order_acceptance_id"`
		} `json:"message"`
	} `json:"params"`
}
