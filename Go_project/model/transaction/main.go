package transaction

import "time"

//Transaction retorna um struct
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

//Transactions retorna um[]
type Transactions []Transaction
