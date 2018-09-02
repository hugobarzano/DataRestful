#!/usr/bin/env bash

HOST=localhost
PORT=8080


function data_index() {
  TITLE=$1
  DATA=$2
  curl -X GET  \
    "http://$HOST:$PORT/" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function data_addDataset() {
  TITLE=$1
  DATA=$2
  curl -X POST  \
    "http://$HOST:$PORT/AddDataset" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
      "title": "'$TITLE'",
      "data": "'$DATA'"
    }'
}

function data_getDataset() {
  ID=$1
  curl -X GET  \
    "http://$HOST:$PORT/datasets/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function data_deleteDataset() {
  ID=$1
  curl -X DELETE  \
    "http://$HOST:$PORT/deleteDataset/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function data_updateDataset() {
  ID=$1
  TITLE=$2
  DATA=$3
  curl -X PUT  \
    "http://$HOST:$PORT/UpdateDataset" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
       "_id": '$ID',
      "title": "'$TITLE'",
      "data": "'$DATA'"
    }'
}
