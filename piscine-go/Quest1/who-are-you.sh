#!/bin/sh
#read input

#if [ "$input" = '"name"' ]; then
    curl -s https://01.gritlab.ax/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
#else
#    printf '%s\n' '"Wrong Batman, maybe robin?"'
#fi