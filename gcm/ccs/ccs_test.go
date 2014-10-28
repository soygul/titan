package ccs

import (
	"os"
	"testing"
)

// GCM environment variables
var senderID = os.Getenv("GCM_SENDER_ID")
var regID = os.Getenv("GCM_REG_ID")
var ccsEndpoint = os.Getenv("GCM_CCS_ENDPOINT")
var apiKey = os.Getenv("GOOGLE_API_KEY")

func TestConnect(t *testing.T) {
}

func TestDisconnect(t *testing.T) {
}

// recv message
// send message
// gcm message types
// message data fields match
// documentation descriptions match

func getConn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	// config := GetConfig()
	// ccsConn, err := New(config.GCM.CCSEndpoint, config.GCM.SenderID, config.GCM.APIKey, config.App.Debug)
	// if err != nil {
	// 	t.Fatalf("Connection to CCS failed with error: %+v", err)
	// }
	// t.Log("CCS connection established.")
	//
	// msgCh := make(chan map[string]interface{})
	// errCh := make(chan error)
	//
	// go ccsConn.Listen(msgCh, errCh)
	//
	// ccsMessage := ccs.NewMessage(config.GCM.RegID)
	// ccsMessage.SetData("hello", "world")
	// ccsConn.Send(ccsMessage)
	//
	// for {
	// 	select {
	// 	case err := <-errCh:
	// 		fmt.Println("err:", err)
	// 		return
	// 	case msg := <-msgCh:
	// 		fmt.Println("msg:", msg)
	// 		return
	// 	}
	// }
}
