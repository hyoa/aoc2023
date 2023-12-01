#!/bin/bash

# The first argument is the day number
day_number=$1

cp gen/dayX.mock internal/day/day$day_number.go

# Use sed to replace the day number in the Go file
sed -i '' "s/X/"$day_number"/g" internal/day/day$day_number.go

cp -r gen/dayX input/day$day_number