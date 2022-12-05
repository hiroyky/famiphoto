#!/bin/sh

curl -XDELETE http://elasticsearch:9200/photo
curl -H "Content-Type: application/json" -XPUT http://elasticsearch:9200/photo/?pretty --data "@photo_mapping.json"
curl -XGET http://elasticsearch:9200/photo/_mapping?pretty