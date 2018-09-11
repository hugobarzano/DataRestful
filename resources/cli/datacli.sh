#!/usr/bin/env bash
HOST=localhost
PORT=8080

function Datarestful_index() {
  curl -X GET  \
    "http://$HOST:$PORT/" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function Datarestful_addDataset() {
  TITLE=$1
  DATA=$2
  curl -X POST  \
    "http://$HOST:$PORT/AddDataset" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
      "title": "'$TITLE'",
      "data": ["1", "2", "3"]
    }'
}

function Datarestful_getDataset() {
  ID=$1
  curl -X GET  \
    "http://$HOST:$PORT/datasets/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function Datarestful_getDatasetByTitle() {
  ID=$1
  curl -X GET  \
    "http://$HOST:$PORT/Search/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function Datarestful_deleteDataset() {
  ID=$1
  curl -X DELETE  \
    "http://$HOST:$PORT/deleteDataset/$ID" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function Datarestful_updateDataset() {
  ID=$1
  TITLE=$2
  #DATA=[1, 2, 3]
  curl -X PUT  \
    "http://$HOST:$PORT/UpdateDataset" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
       "_id": "'$ID'",
      "title": "'$TITLE'",
      "data": ["1111", "2222", "3333"]
    }'
}

function Datarestful_listservices() {
  curl -X GET  \
    "http://$HOST:$PORT/Services/" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json'
}

function Datarestful_operation() {
  ID=$1
  VALUE=$2
  OPE=$3
  URL=$4
  curl -X POST  \
    "http://$HOST:$PORT/Operation/" \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -d '{
	    "dataset_id": "'$ID'",
	    "value": "'$VALUE'",
	    "operator": "'$OPE'",
	    "service_url":  "'URL'"
    }'
}
