## Golang Project
### 
* Gin Framework
    * 建立 REST API
* GORM
    * 處理資料庫與CRUD
* swagger
    * 建立API文件
    * 只要有變動，需 執行`swag init`
* GCP (CloudSQL, Cloud Run)
    * 部屬API至雲端服務
    * 連線至雲端資料庫
* 必需 cmd
``` 
    go get 套件名稱
    swag init
    go run main.go
```
## 部署至GCP
參考網址 : <https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service>

須建立Dockerfile
```
FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

CMD ["/app/server"]
```

開啟 GCloud SDK Shell 並在專案目錄下執行 `gcloud run deploy`
## 遭遇之問題 (問題序號對照解決方法序號)
1. 部屬完後須轉換為https
2. CloudSQL無法正常連線
3. Request欄位無法新增至DB，Ex.
    * Response body回傳空字串、資料庫只有id新增成功(因id為auto_increments)，其餘欄位在資料庫為""
    ```
    {
        "ShopId": 14,
        "ShopTitle": "",
        "ShopPhone": "",
        "ShopAddress": ""
    }
    ```
## 解決方法
1. 
* https建立方法，產生私鑰key.pem和證書cert.pem
```
go run  "C:\Program Files\Go\src\crypto\tls\generate_cert.go" --host="localhost"
```
* Swagger 設定需更改( // @host 與 ginSwagger.URL(<span>"http://localhost:8080/swagger/doc.json"</span>))， `swag init`->更新設定檔

2.  需使用另一種連線方式 
    * 匯入 <"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql">
    ```
    var dbConnectionString = dbUserName + ":" + dbPassword + "@cloudsql(" + Project + ":" + Region + ":" + SqlInstanceId + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dbConnectionString)
    ```

3. Request body的model
需加入 **binding** :
```
type 名稱 struct {
    欄位 型態 `json:"ShopPhone" binding:"required"`
}
```
    
## API線上網址(已部屬至GCP)
<https://golangapi-njqijcw6la-df.a.run.app/swagger/index.html>




