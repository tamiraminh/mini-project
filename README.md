# mini-project
bootcamp evermos REST API

## How To Run
1. get all dependencies defined in go.mod file in terminal
```
go mod download
```
2. Import database, you can find database in ./internal/database/booking.sql. import using this command to your mysql server 
```
mysql -u username -p new_database < ./internal/database/bootcamp.sql
```
before import it, create new database in your mysql. bootcamp.sql contains 2 table, there is booking and material. in this project, only use booking table.       

3. Configure Database Connecion, you can find this on internal/database/service.go. you will findout connection function like this
```
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:<mysqlserverport>)/<databasename>")
	if err != nil {
		return nil, err
	}

	return db, nil
}
```
best practice is using .env file. ensure your confidential didnt get public   

4. run the server with in the root directory
```
go run main.go
```


## Documentation 
### Postman Collection Run 
![alt text](./TestCollectionMiniProject.png)

to import documentation you can use 
1. Mini Project Env.postman_environment.json
2. Mini Project Evermos.postman_collection.json

on this root project and import it to postman 

to see swagger documentation you can 
1. update swagger 
```
swag init
```
2. run server
```
go run main.go
```
3. access swagger URL in you local browser 
```
http://localhost:8080/swagger/
```

made by : Muhammad Tamiramin Hayat Suhendar