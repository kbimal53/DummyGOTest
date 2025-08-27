#!/bin/bash

echo "🚀 Preparing for Vercel deployment..."

# Check if we're in a git repository
if [ ! -d .git ]; then
    echo "❌ Not a git repository. Please run 'git init' first."
    exit 1
fi

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

# Check for uncommitted changes
if [ -n "$(git status --porcelain)" ]; then
    echo "📝 Found uncommitted changes. Committing them..."
    git add .
    read -p "Enter commit message (or press Enter for default): " commit_msg
    if [ -z "$commit_msg" ]; then
        commit_msg="Deploy: Update API for Vercel deployment $(date '+%Y-%m-%d %H:%M')"
    fi
    git commit -m "$commit_msg"
fi

# Check if we're ahead of remote
if [ -n "$(git log origin/main..HEAD 2>/dev/null)" ]; then
    echo "🔄 Pushing changes to GitHub..."
    git push origin main
    echo "✅ Changes pushed to GitHub!"
elif [ $? -ne 0 ]; then
    echo "⚠️  No remote repository configured. Please set up GitHub remote first:"
    echo "   git remote add origin https://github.com/yourusername/yourrepo.git"
    echo "   git push -u origin main"
else
    echo "✅ Repository is up to date with remote!"
fi

echo ""
echo "🎉 Deployment preparation complete!"
echo ""
echo "📋 Next steps:"
echo "1. ✅ Code pushed to GitHub"
echo "2. 🌐 Go to https://vercel.com/dashboard"
echo "3. 🔗 Import your GitHub repository (if not already connected)"
echo "4. ⚙️  Set environment variables in Vercel dashboard"
echo "5. 🚀 Deploy!"
echo ""
echo "🔄 For future deployments, just run:"
echo "   git add . && git commit -m 'your message' && git push origin main"
echo ""
echo "💡 Tip: Vercel will automatically deploy when you push to main branch!"
