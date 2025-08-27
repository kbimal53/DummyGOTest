#!/bin/bash

echo "🚀 Preparing for Vercel deployment..."

# Copy environment file if it exists
if [ -f .env ]; then
    echo "📋 Environment file found"
    echo "⚠️  Remember to set environment variables in Vercel dashboard:"
    echo "   - DATABASE_URL"
    echo "   - POSTGRES_URL"
    echo "   - POSTGRES_USER"
    echo "   - POSTGRES_HOST"
    echo "   - POSTGRES_PASSWORD"
    echo "   - POSTGRES_DATABASE"
else
    echo "⚠️  No .env file found. Make sure to set environment variables in Vercel."
fi

# Install dependencies for api directory
echo "📦 Installing dependencies for serverless function..."
cd api
go mod tidy
cd ..

echo "✅ Ready for deployment!"
echo ""
echo "Next steps:"
echo "1. Push to GitHub: git push origin main"
echo "2. Connect repository to Vercel"
echo "3. Set environment variables in Vercel dashboard"
echo "4. Deploy!"
