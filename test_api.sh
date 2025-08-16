#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"

echo "🚀 Testing Financial Transaction System API"
echo "============================================="

# Test health endpoint
echo "📊 Testing health endpoint..."
curl -s "$BASE_URL/../health" | jq '.'
echo ""

# Test user registration
echo "👤 Testing user registration..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword123",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890"
  }')

echo "$REGISTER_RESPONSE" | jq '.'

# Extract access token
ACCESS_TOKEN=$(echo "$REGISTER_RESPONSE" | jq -r '.access_token')
echo "🔑 Access Token: ${ACCESS_TOKEN:0:50}..."
echo ""

# Test user login
echo "🔐 Testing user login..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword123"
  }')

echo "$LOGIN_RESPONSE" | jq '.'
echo ""

# Test get user profile
echo "👥 Testing get user profile..."
curl -s -X GET "$BASE_URL/users/profile" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
echo ""

# Test update user profile
echo "✏️ Testing update user profile..."
curl -s -X PUT "$BASE_URL/users/profile" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "first_name": "Johnny",
    "address": "123 Main St, New York, NY"
  }' | jq '.'
echo ""

# Test change password
echo "🔒 Testing change password..."
curl -s -X POST "$BASE_URL/users/change-password" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "current_password": "securepassword123",
    "new_password": "newsecurepassword456"
  }' | jq '.'
echo ""

# Test refresh token
REFRESH_TOKEN=$(echo "$REGISTER_RESPONSE" | jq -r '.refresh_token')
echo "🔄 Testing refresh token..."
curl -s -X POST "$BASE_URL/auth/refresh" \
  -H "Content-Type: application/json" \
  -d "{\"refresh_token\": \"$REFRESH_TOKEN\"}" | jq '.'
echo ""

# Test unauthorized access
echo "🚫 Testing unauthorized access..."
curl -s -X GET "$BASE_URL/users/profile" | jq '.'
echo ""

echo "✅ API testing completed!" 