#!/bin/sh

# Usage: ./gen_auth.sh [merchant_user_id] [secret_key]
# Or export MERCHANT_USER_ID and SECRET_KEY before running

merchant_user_id="${1:-$MERCHANT_USER_ID}"
secret_key="${2:-$SECRET_KEY}"

# Validate input
if [ -z "$merchant_user_id" ] || [ -z "$secret_key" ]; then
  echo "❌ Usage: $0 [merchant_user_id] [secret_key] or set MERCHANT_USER_ID and SECRET_KEY env vars"
  exit 1
fi

timestamp=$(date +%s)
digest=$(echo -n "${timestamp}${secret_key}" | sha1sum | awk '{print $1}')
auth_header="${merchant_user_id}:${digest}:${timestamp}"

echo "Auth: $auth_header"
