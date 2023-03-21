# go-nosql

## Configuration
Install database:
```
$> brew install mongodb
```
Install client database
```
$> brew install --cask robo-3t
```
## Running this application
Start database
```
$> make db_up
```
Start Application
```
$> go run app/main.go
```
Stop database
```
$> make db_down
```

## Execute application
```
$> curl --location 'localhost:3000/routes' \
--header 'Content-Type: application/json' \
--data '{
    "total_routes": 5000
}'
```