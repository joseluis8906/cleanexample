package order

type (
	Order struct {
		Customer string `json:"customer" bson:"customer"`
		Courier  string `json:"courier" bson:"courier"`
		Address  string `json:"address" bson:"address"`
		Amount   int    `json:"amount" bson:"amount"`
	}
)
