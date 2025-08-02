#!/bin/bash

set -e

# Load env vars
export $(cat .env | xargs)

# Optional: path override
TEMPLATE_FILE=${1:-nginx.conf.template}
OUTPUT_FILE=${2:-nginx.conf}

echo "🛠️  Generating $OUTPUT_FILE from $TEMPLATE_FILE..."

# Only substitute your intended variables (leave $host, $remote_addr, etc. intact)
envsubst '$DOMAIN $CERTBOT_EMAIL' < "$TEMPLATE_FILE" > "$OUTPUT_FILE"

echo "✅ Done: $OUTPUT_FILE created."
