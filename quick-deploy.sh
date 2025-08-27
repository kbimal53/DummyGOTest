#!/bin/bash

# Quick deployment script for updating Vercel production
echo "🚀 Quick Deploy to Vercel Production"
echo ""

# Check for changes
if [ -z "$(git status --porcelain)" ]; then
    echo "⚠️  No changes detected. Nothing to deploy."
    exit 0
fi

# Show current changes
echo "📝 Current changes:"
git status --short

echo ""
read -p "💬 Enter commit message: " commit_msg

if [ -z "$commit_msg" ]; then
    echo "❌ Commit message is required!"
    exit 1
fi

# Add, commit, and push
echo "📦 Adding changes..."
git add .

echo "💾 Committing changes..."
git commit -m "$commit_msg"

echo "🔄 Pushing to GitHub (this will trigger Vercel deployment)..."
git push origin main

echo ""
echo "✅ Deployment initiated!"
echo "🌐 Check your Vercel dashboard for deployment progress:"
echo "   https://vercel.com/dashboard"
echo ""
echo "⏱️  Deployment usually takes 1-2 minutes."
echo "🎉 Your API will be live at your Vercel URL once deployment completes!"
