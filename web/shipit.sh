#!/bin/bash
cd ng2-app && ng build ---prod --aot && cd .. && gcloud app deploy --project pfl-order-demo --version 1 --no-promote --quiet && say "deploy complete" || say "deploy failed"
