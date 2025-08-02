#!/bin/bash

set -e

cd "$(dirname "$0")"

# Load env vars from .env safely
if [ -f .env ]; then
  echo "🔐 Loading environment variables from .env..."
  set -o allexport
  source .env
  set +o allexport
else
  echo "⚠️  .env file not found — aborting."
  exit 1
fi

SERVICE=${1:-go}
CERT_PATH="/etc/letsencrypt/live/$DOMAIN/fullchain.pem"
KEY_PATH="/etc/letsencrypt/live/$DOMAIN/privkey.pem"

echo "📌 Target: $SERVICE"
echo "🌍 Domain: $DOMAIN"

case "$SERVICE" in
  go)
    echo "📦 Pulling latest Go app image..."
    docker-compose pull go_app

    echo "🚀 Restarting Go app..."
    docker-compose up -d --no-deps go_app
    ;;

  certbot)
    echo "🔧 Generating bootstrap nginx.conf for Certbot challenge..."
    ./generate_nginx_config.sh nginx.bootstrap.conf nginx.conf

    echo "🧼 Shutting down all running containers..."
    docker-compose down

    echo "🚀 Starting temporary Nginx on port 80..."
    docker-compose -f docker-compose.cert-init.yml up -d nginx

    echo "🔐 Running Certbot..."
    docker-compose -f docker-compose.cert-init.yml run --rm certbot || echo "⚠️  Certbot may have already succeeded"

    echo "🛑 Shutting down bootstrap stack..."
    docker-compose -f docker-compose.cert-init.yml down

    if [ -f "$CERT_PATH" ] && [ -f "$KEY_PATH" ]; then
      echo "✅ Certificate found — starting full stack..."
    else
      echo "❌ Certificate not found — aborting"
      exit 1
    fi

    echo "🔧 Generating full nginx.conf with SSL..."
    ./generate_nginx_config.sh nginx.conf.template nginx.conf

    echo "🚀 Starting Go app and Nginx with SSL..."
    docker-compose up -d
    ;;

  *)
    echo "❌ Unknown target: $SERVICE"
    echo "Usage: $0 [go|certbot]"
    exit 1
    ;;
esac

echo "✅ Deployment complete for: $SERVICE"
