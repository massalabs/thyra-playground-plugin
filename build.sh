#!/bin/bash
set -ex

mkdir -p build/simulator
node scripts/get-simulator.js -o build
go build -o build/thyra-playground-plugin cmd/main.go
cp manifest.json build/