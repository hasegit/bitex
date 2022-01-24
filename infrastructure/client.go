package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const (
	subscribe = "subscribe"
	endpoint  = "wss://ws.lightstream.bitflyer.com/json-rpc"
)

type BitFlyerClient struct {
	conn *websocket.Conn
}

func NewBitFlyerClient() BitFlyerClient {

	// 接続の確立
	c, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	return BitFlyerClient{conn: c}
}

func (c BitFlyerClient) subscribe(channel string) (err error) {

	// Subscribeリクエストの送信
	w := NewJsonRequest(channel, subscribe)
	if err = c.conn.WriteJSON(w); err != nil {
		return err
	}

	// 結果の確認
	// 単純にエラー発生したらリターン
	message, err := c.ReadMessage()
	if err != nil {
		return err
	}

	// レスポンスがfalseの場合もエラーリターン
	var r JsonResponse
	json.Unmarshal(message, &r)
	if !r.Result {
		return fmt.Errorf("subscribe failed")
	}
	return nil
}

func (c BitFlyerClient) SubscribeExecutions() error {
	// https://bf-lightning-api.readme.io/docs/realtime-executions
	return c.subscribe("lightning_executions_FX_BTC_JPY")
}

func (c BitFlyerClient) Close() {
	c.conn.Close()
}

func (c BitFlyerClient) ReadMessage() ([]byte, error) {

	// wsからデータを受け取る
	_, message, err := c.conn.ReadMessage()
	return message, err
}
