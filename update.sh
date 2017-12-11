#!/usr/bin/env bash

git pull
cd web && npm run build
cd ../backend/ && sh update.sh