# Execute
- create a .docker/mysql directory
```
docker compose up -d --build
```


Create a transaction ```./goapp/api/client.http```. 
```
POST http://localhost:8080/transactions HTTP/1.1
Content-Type: application/json

{
    "account_id_from": "0e96d032-86fd-11ec-8b22-9a5ce86758a4",
    "account_id_to": "534b6b56-a988-11ec-b7e0-2b8e9696da41",
    "amount": 5
}
```
Make a request ```./consumer/api/client.http```.
```
 GET http://localhost:3003/balances/0e96d032-86fd-11ec-8b22-9a5ce86758a4

```

