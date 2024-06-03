#! /bin/bash

cd api/pkg/frontend
npm install
npm run build
cd ../../../

go build api/main.go
