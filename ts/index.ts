const id = 'send_money_button';
// idで指定して、elementを取得
const e = document.getElementById(id);
// イベントリスナー登録
// クリックされたら、コールバック関数発動
e.addEventListener('click', () => {
    // 生のjavaScriptでAjaxを行うためには、XMLHttpRequestを使用する

    // XMLHttpRequestインスタンスの生成
    const httpRequest = new XMLHttpRequest();

    // openメソッド : リクエストを初期化する
    // 第1引数で、GETかPOSTのリクエストメソッドを指定
    // 第2引数で、"リクエスト先のURL"
    // 第3引数で、同期か非同期かの選択（非同期の場合は、true）
    httpRequest.open("GET", "/wallet", true)

    // readyStateプロパティ : XMLHttpRequestオブジェクト(HTTPリクエスト)のサーバとの通信の状態
    // この値がサーバとの通信状態によって変化する
    // readyStateプロパティの値が変化するたびにonreadystatechangeイベントが発生する
    // このonreadystatechangeイベントが発生したら、readyStateプロパティをチェックすれば、現在の通信状態を知ることができる

    console.log(httpRequest.readyState);

    // XMLHttpRequestオブジェクト(HTTPリクエスト)のonreadystatechangeイベントが発生したら、=で与えた関数が発動する
    httpRequest.onreadystatechange = () => {

        // 通信が完了したならば
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {
            // jsonをtextデータとして受け取る
            const jsonText = httpRequest.responseText

            // textデータとして受け取ったjsonをパース
            const json = JSON.parse(jsonText);

            // 通信が完了（受信）したら、行う処理を記載
            e.innerHTML = `
            <p>名前: ${json["Name"]}</p>
            <p>型　: ${json["LangType"]}</p>
            `
        }
    }
    // sendメソッド : openで作成したHTTPリクエストをサーバへ送信する
    // POSTの場合は、送信するデータを引数で与える
    // GETの場合は、nullを引数で与える
    httpRequest.send(null);

    // 非同期のテスト
    // 非同期で通信しているので、HTMLの<p>が変化するよりもこちらの方が先に呼び出される
    window.alert('非同期');

})

