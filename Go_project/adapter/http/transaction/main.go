package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Strke12i/Go_project/Go_project/model/transaction"
	"github.com/Strke12i/Go_project/Go_project/util"
)

//GetTransactions GEt
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salario",
			Amount:    1500.0,
			Type:      0,
			CreatedAt: util.StringToTime("2019-02-12T18:43:23"),
		},
	}

	_ = json.NewEncoder(w).Encode((transactions))
}

//CreateATransaction Post
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader((http.StatusMethodNotAllowed))
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
