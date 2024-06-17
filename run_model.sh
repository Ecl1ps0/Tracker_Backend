#!/bin/bash

solution=$SOLUTION

echo "$solution"

echo "Running script with input: $1"
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <file_path>"
    exit 1
fi

echo "Executing Python script..."
python "./ml_model/main.py" "$1"
echo "Python script finished with status: $?"
