#! /bin/bash

####
# Sample API call to get today's scoreboard JSON result for a specific league.
####

SPORT="${SPORT:-hockey}"
LEAGUE="${LEAGUE:-nhl}"
GAMEDATE="${GAMEDATE:-$(date +"%Y%m%d")}"

URL="https://site.api.espn.com/apis/site/v2/sports/$SPORT/$LEAGUE/scoreboard?dates=$GAMEDATE"

SCRIPT_DIR=$(dirname "$0")


curl $URL | jq > "$SCRIPT_DIR/sample-$GAMEDATE-$SPORT-$LEAGUE.json"
