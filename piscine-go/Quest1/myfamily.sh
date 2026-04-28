#!/bin/bash

# Print the relatives of the hero identified by HERO_ID (no quotes)
URL="https://01.gritlab.ax/assets/superhero/all.json"

if [ -z "$HERO_ID" ]; then
	echo ""
	exit 1
fi

curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select((.id | tostring) == ($id)) | .connections.relatives | gsub("\n"; "\\n")'