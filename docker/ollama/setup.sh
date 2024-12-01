#!/bin/sh

ollama serve &

echo "Waiting for Ollama server to be active..."
while [ "$(ollama list | grep 'NAME')" == "" ]; do
  sleep 1
done

ollama pull llava:13b
ollama pull 7shi/llama-translate:8b-q4_K_M
