#!/bin/bash

API_URL="http://localhost:6001"
EMAIL="test@example.com"
PASSWORD="secret123"

echo "🔹 1. SIGNUP"
curl -s -X POST $API_URL/signup \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\", \"password\":\"$PASSWORD\"}"
echo -e "\n"

echo "🔹 2. LOGIN"
LOGIN_RESPONSE=$(curl -s -X POST $API_URL/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\", \"password\":\"$PASSWORD\"}")

ACCESS_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.access_token')
REFRESH_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.refresh_token')

echo "✅ Access Token: $ACCESS_TOKEN"
echo "✅ Refresh Token: $REFRESH_TOKEN"
echo

echo "🔹 3. GET /holdings"
curl -s -X GET $API_URL/holdings \
  -H "Authorization: Bearer $ACCESS_TOKEN"
echo -e "\n"

echo "🔹 4. GET /orderbook"
curl -s -X GET $API_URL/orderbook \
  -H "Authorization: Bearer $ACCESS_TOKEN"
echo -e "\n"

echo "🔹 5. GET /positions"
curl -s -X GET $API_URL/positions \
  -H "Authorization: Bearer $ACCESS_TOKEN"
echo -e "\n"

echo "🔹 6. REFRESH TOKEN"
curl -s -X POST $API_URL/refresh \
  -H "Content-Type: application/json" \
  -d "{\"token\":\"$REFRESH_TOKEN\"}"
echo -e "\n"

echo "🔹 7. HEALTH CHECK"
curl -s $API_URL/health
echo -e "\n"
