#!/bin/bash

exportFilePath="/home/knightchaser/random.txt"

# Remove the file if it exists
if [ -f "$exportFilePath" ]; then
    rm "$exportFilePath"
fi
# Generate 10 random hex numbers, each 100 bytes long
for ((i=1; i<=10; i++))
do
    hex_number=$(openssl rand -hex 50)
    echo "$hex_number" >> "$exportFilePath"
done