#!/bin/bash

set -uex

CURRENT=$(cd "$(dirname "$0")" && pwd)
VERSION=$1
MAJOR=$(echo "$VERSION" | cut -d. -f1)
MINOR=$(echo "$VERSION" | cut -d. -f2)
PATCH=$(echo "$VERSION" | cut -d. -f3)

git commit -m "bump up v$MAJOR.$MINOR.$PATCH"
git tag "v$MAJOR.$MINOR.$PATCH"
git push --tags
