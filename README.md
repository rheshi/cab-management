# cab-management

# Dependencies:
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite 

gnorm sqlite used hence has a dependeicy on c++ while building:
set CGO_ENABLED=1
# Run your Go app again
go run main.go 

# build exe syntax:
go build -o cab-management-server.exe maincode.go

run exe directly to launch the server

## REST API CURL

# 1. Register Cab (POST /cab/register)

curl -X POST http://localhost:8080/cab/register \
-H "Content-Type: application/json" \
-d '{
  "LicensePlate": "ABCD1234",
  "CurrentCityID": 1
}'

# 2. Add City (POST /city/add)

curl -X POST http://localhost:8080/city/add \
-H "Content-Type: application/json" \
-d '{
  "Name": "New York"
}'

# 3. Change Cab Location (PUT /cab/change-location)

curl -X PUT http://localhost:8080/cab/change-location \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1,
  "city_id": 2
}'

# 4. Change Cab State (PUT /cab/change-state)

curl -X PUT http://localhost:8080/cab/change-state \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1,
  "state": "ON_TRIP"
}'

# 5. Book Cab (POST /cab/book)

curl -X POST http://localhost:8080/cab/book \
-H "Content-Type: application/json" \
-d '{
  "city_id": 1
}'

# 6. Cab Idle Time (POST /cab/idle-time)

curl -X POST http://localhost:8080/cab/idle-time \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1
}'

# 7. Cab History (POST /cab/history)
curl -X POST http://localhost:8080/cab/history \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1
}'