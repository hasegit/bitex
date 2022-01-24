package adapter

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hasegit/bitex/domain"
)

type BitFlyerRepository struct {
	client WSClient
}

func NewBitFlyerRepository(client WSClient) BitFlyerRepository {
	return BitFlyerRepository{client}
}

func (r BitFlyerRepository) ReadExecutions(ch chan domain.MarketData) {

	// 約定データの購読を開始
	r.client.SubscribeExecutions()

	// 受信データをドメインオブジェクトに変換し、
	// channel経由でinteractorに返す
	done := make(chan struct{})

	for {
		// データを受け取る
		message, err := r.client.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		// データを扱いやすいように一旦構造体に入れる
		var e BitFlyerExecution
		err = json.Unmarshal(message, &e)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e)

		// データから必要な部分を抽出する
		for _, data := range e.Params.Message {
			ch <- domain.MarketData{
				Date:  data.ExecDate,
				Price: data.Price,
			}
		}
	}

	<-done
}

/*

// メッセージを受信し、構造体に格納
message, err := client.ReadMessage()
if err != nil {
	log.Println("read:", err)
	os.Exit(1)
}
var e domain.Execution
err = json.Unmarshal(message, &e)
if err != nil {
	log.Fatal(err)
}

for _, m := range e.Params.Message {
	fmt.Println(m.Price, m.ExecDate)
}
*/
