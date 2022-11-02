#!/bin/bash
set -ex

# Init playground submodule
git submodule init
git submodule update

# Copy static files
mkdir -p cmd/static
pushd massa-sc-playground
cp -r libs favicons index.html index.css module.js ../cmd/static
popd

# Download simulator
node scripts/get-simulator.js
