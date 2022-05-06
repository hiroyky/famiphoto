#!/bin/sh

#curl -XDELETE http://localhost:9200/photo
curl -H "Content-Type: application/json" -XPUT http://localhost:9200/photo/?pretty --data "@photo_mapping.json"
curl -XGET http://localhost:9200/photo/_mapping?pretty