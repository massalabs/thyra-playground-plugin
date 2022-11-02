#!/bin/bash
set -ex

PLUGIN=playground-plugin

mkdir -p "build/$PLUGIN/simulator"
# node scripts/get-simulator.js -o build/$PLUGIN
cp simulator/massa-sc-tester* build/$PLUGIN/simulator
go build -o build/$PLUGIN/thyra-playground-plugin cmd/main.go
cp manifest.json build/$PLUGIN/