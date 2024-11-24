#!/bin/sh


for i in {30..0}; do
    if curl elasticsearch:9200; then
        curl -H "Content-Type: application/json" -XPUT http://elasticsearch:9200/photo/?pretty --data "@elasticsearch/photo_mapping.json"
        break;
    fi
    sleep 2
done

#curl -XDELETE http://elasticsearch:9200/photo
curl -H "Content-Type: application/json" -XPUT http://elasticsearch:9200/photo/?pretty --data "@elasticsearch/photo_mapping.json"
curl -XGET http://elasticsearch:9200/photo/_mapping?pretty