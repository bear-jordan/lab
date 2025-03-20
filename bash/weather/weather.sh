#!/usr/bin/env bash

command=$1
zone_mappings="$(git rev-parse --show-toplevel)/data/zone_mapping.txt"

if [[ ! $command == "help" && ! $command == "check" && ! $command == "list" ]]; then
    echo "Invalid command. Use [help|check|list]"
    exit 1
fi

if [[ $command == "help" ]]; then
    echo "usage: weather [help] [check <state> <code>] [list]"
    echo ""
    echo "Interact with a search to find bad weather warnings"
    echo "    list"
    echo ""
    echo "Provide a state and code to find bad weather warnings"
    echo "    check <state> <code>"
    exit 0
fi

if [[ $command == "list" ]]; then
    valid_states=$(awk -F '|' '{print $1}' "$zone_mappings" | sort -u | grep -v STATE)
    declare -u state
    read -r -p "Search for a state: " state
    if [[ $valid_states == "*$state*" ]]; then
        echo "Invalid state"
        exit 1
    fi
    location=$(grep "$state" "$zone_mappings" | awk -F '|' '{print $4}' | sort | fzf)
    full_code=$(grep "$state" "$zone_mappings" | grep "$location" | awk -F '|' '{print $1 "Z" $2}' | head -n1)
    response=$(curl -s -X GET "https://api.weather.gov/alerts/active/zone/$full_code" -H "accept: application/geo+json")
    features=$(echo "$response" | jq '.features')
    if [[ $features == "[]" ]]; then
        echo "No warnings found"
    else
        echo "$response" | jq '.title'
        echo "$features" | jq | less
    fi
    exit 0
fi

state=$2
code=$3
full_code="${state}Z${code}"
valid_codes=$(awk -F '|' '{print $1 "Z" $2}' "$zone_mappings")

if [[ ! $valid_codes == *$full_code* ]]; then
    echo "Invalid code"
    exit 1
fi

response=$(curl -s -X GET "https://api.weather.gov/alerts/active/zone/$full_code" -H "accept: application/geo+json")
features=$(echo "$response" | jq '.features')
if [[ $features == "[]" ]]; then
    echo "No warnings found"
else
    echo "$response" | jq '.title'
    echo "$features" | jq | less
fi
exit 0
