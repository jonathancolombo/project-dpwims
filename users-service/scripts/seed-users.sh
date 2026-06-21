#!/bin/sh

echo "Seeding default users..."

API_URL="http://localhost:8081/users"

echo "Waiting for users-service to be ready..."
until curl -s -o /dev/null "$API_URL" 2>/dev/null; do
  sleep 1
done

create_user() {
  USERNAME=$1
  PASSWORD=$2
  EMAIL=$3
  FISCAL=$4
  PHONE=$5
  ROLE=$6

  echo "Creating user: $USERNAME"

  HTTP_CODE=$(curl -s -o /tmp/response.json -w "%{http_code}" -X POST "$API_URL" \
    -H "Content-Type: application/json" \
    -d "{
      \"username\": \"$USERNAME\",
      \"password\": \"$PASSWORD\",
      \"email\": \"$EMAIL\",
      \"fiscal_code\": \"$FISCAL\",
      \"telephone\": \"$PHONE\",
      \"role\": \"$ROLE\"
    }")

  echo "Response code: $HTTP_CODE"
  cat /tmp/response.json
  echo ""
}

create_user "admin" "Password123!" "admin@example.com" "RSSMRA80A01H501U" "+390000000002" "admin"
create_user "customer" "Password123!" "customer@example.com" "AAARA80A01H501U" "+390000000001" "customer"

echo ""
echo "Seeding completed."