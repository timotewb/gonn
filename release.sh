#!/bin/bash

# Define the version number
VERSION="v0.0.6"

# Commit the changes with the version number
git commit -am "$VERSION"

# Tag the commit with the version number
git tag $VERSION

# Push the tag to the remote repository
git push origin $VERSION

# Push the changes to the main branch
git push origin main
