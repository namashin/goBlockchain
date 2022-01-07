package main

import (
	"bytes"
	"encoding/json"
	"goblockchain/block"
	"goblockchain/utils"
	"goblockchain/wallet"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port: port, gateway: gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join("templates", "index.html"))
		t.Execute(w, "")
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	case http.MethodGet:
		io.WriteString(w, "/walletのURLはPOST専用だよ～,GETしないで")
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTp Method")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Println("ERROR: Invalid HTTP Method", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.ValidDate() {
			log.Println("ERROR: missing field value")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR Parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		value32 := float32(value)
		//fmt.Println(publicKey)
		//fmt.Println(privateKey)
		//fmt.Printf("%v", value32)
		w.Header().Add("Content-Type", "application/json")

		transaction := wallet.NewTransaction(privateKey, publicKey, *t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, value32)
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			t.SenderBlockchainAddress,
			t.RecipientBlockchainAddress,
			t.SenderPublicKey,
			t.Value,
			&signatureStr,
		}
		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)

		resp, _ := http.Post(ws.gateway+"/transactions", "application/json", buf)
		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}
		io.WriteString(w, string(utils.JsonStatus("fail")))

	case http.MethodGet:
		io.WriteString(w, "/transactionのURLはPOST専用だよ～、GETしないで")
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {

	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
