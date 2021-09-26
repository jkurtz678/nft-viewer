#!/usr/bin/env sh

set -e

npm run build

cd dist

git init
git add -A
git commit -m "New Deployment"
git push -f git@github.com:jkurtz678/nft-viewer.git master:gh-pages
cd -