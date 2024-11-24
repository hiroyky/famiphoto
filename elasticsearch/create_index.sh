#!/bin/sh


for i in {30..0}; do
    if curl localhost:9200; then
        curl -H "Content-Type: application/json" -XPUT http://localhost:9200/photo/?pretty --data "@elasticsearch/photo_mapping.json"
        break;
    fi
    sleep 2
done

#curl -XDELETE http://localhost:9200/photo
curl -H "Content-Type: application/json" -XPUT http://localhost:9200/photo/?pretty --data "@elasticsearch/photo_mapping.json"
curl -XGET http://localhost:9200/photo/_mapping?pretty