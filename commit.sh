
#!/bin/bash

set -e

VERSION=$(cat version)

if [ -z "$VERSION" ]; then
  echo "Version file is empty or missing"
  exit 1
fi

echo "Tagging version: $VERSION"
git tag -a "$VERSION" -m "Release version $VERSION"
git push origin "$VERSION"

