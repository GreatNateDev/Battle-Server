#!/bin/bash
echo Building Container
docker build -t pokebackend .
echo Finished Build
echo Starting Server
docker run --rm -p 8080:8080 pokebackend
echo Finished Server