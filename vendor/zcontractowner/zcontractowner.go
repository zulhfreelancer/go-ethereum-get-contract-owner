package zcontractowner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetContractOwner will return the owner of a contract in string format.
// E.g. 0x463cf5545ea6da915cf37483a48a5f36bb7f7845
func GetContractOwner(rpcServer, contractAddress string) string {
	type rpcParam struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	type rpcPayload struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}

	type rpcResponse struct {
		Jsonrpc string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"`
	}

	var rpcRes rpcResponse

	param := rpcParam{
		To:   contractAddress,
		Data: "0x8da5cb5b", // owner getter function signature
	}

	params := make([]interface{}, 2)
	params[0] = param
	params[1] = "latest"

	data := rpcPayload{
		Jsonrpc: "2.0",
		Method:  "eth_call",
		Params:  params,
		ID:      1,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", rpcServer, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resBody, &rpcRes)
	if err != nil {
		panic(err)
	}

	address := rpcRes.Result[len(rpcRes.Result)-40:]
	withZeroX := fmt.Sprintf("0x%v", address)
	return withZeroX
}
