# cab-management

#1. Register Cab (POST /cab/register)
bash
Copy code
curl -X POST http://localhost:8080/cab/register \
-H "Content-Type: application/json" \
-d '{
  "LicensePlate": "ABCD1234",
  "CurrentCityID": 1
}'

#2. Add City (POST /city/add)
bash
Copy code
curl -X POST http://localhost:8080/city/add \
-H "Content-Type: application/json" \
-d '{
  "Name": "New York"
}'

#3. Change Cab Location (PUT /cab/change-location)
bash
Copy code
curl -X PUT http://localhost:8080/cab/change-location \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1,
  "city_id": 2
}'

#4. Change Cab State (PUT /cab/change-state)
bash
Copy code
curl -X PUT http://localhost:8080/cab/change-state \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1,
  "state": "ON_TRIP"
}'

#5. Book Cab (POST /cab/book)
bash
Copy code
curl -X POST http://localhost:8080/cab/book \
-H "Content-Type: application/json" \
-d '{
  "city_id": 1
}'

#6. Cab Idle Time (POST /cab/idle-time)
bash
Copy code
curl -X POST http://localhost:8080/cab/idle-time \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1
}'

#7. Cab History (POST /cab/history)
bash
Copy code
curl -X POST http://localhost:8080/cab/history \
-H "Content-Type: application/json" \
-d '{
  "cab_id": 1
}'