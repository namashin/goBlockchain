# goBlockchain

blockchainの基礎を実装

## Description

トランザクションは送金のログ
ブロックはトランザクションをまとめたもの
ブロックチェーンはブロックのハッシュを次のブロック内に
持つことでチェーンのように連結されたデータ構造



###blockchain_server側
// blockの中身表示
127.0.0.1:5000/　

// poolに溜まった取引記録表示
127.0.0.1:5000/transactions　　

// mining by hand
127.0.0.1:5000/mine

// automaticaly mining 
127.0.0.1:5000/mine/start　　　

// get bc amount 
127.0.0.1:5000/amount　　　　　 


###wallet_server側
// parse Template/index.html
127.0.0.1:8080/             

// create new wallet for POST
127.0.0.1:8080/wallet        

// get wallet amount of blockchain address
127.0.0.1:8080/wallet/amount

// making transaction request for POST
127.0.0.1:8080/transaction
