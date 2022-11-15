#!/bin/bash
set -ex

PLUGIN=playground-plugin

mkdir -p "build/$PLUGIN/simulator"
cp simulator/massa-sc-tester* build/$PLUGIN/simulator
go build -o build/$PLUGIN/thyra-playground-plugin cmd/main.go
cp manifest.json build/$PLUGIN/