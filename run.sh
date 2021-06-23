#!/bin/bash

echo "Enter name:"
read authorName

echo ""
echo "Starting server..."
PORT=19132 AUTHOR_NAME=$authorName ./server