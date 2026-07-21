#!/bin/bash

curl https://acad.learn2earn.ng/assets/superhero/all.json | jq --arg JACK "$HERO_ID" ' .[] | select(.id == ($JACK | tonumber)) | .connections.relatives' | sed 's/"//g'

