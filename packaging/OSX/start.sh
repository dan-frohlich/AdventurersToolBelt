#!/bin/bash

mkdir -p ~/.adventurers_toolbelt 2>>/dev/null

set  \
  > ~/.adventurers_toolbelt/out.log \
  2>> ~/.adventurers_toolbelt/out.log

cd "$(dirname "$0")"

echo "$(date) app - $BASENAME" \
  >> ~/.adventurers_toolbelt/out.log \
  2>> ~/.adventurers_toolbelt/out.log

echo $(date) pwd - $(pwd) \
  >> ~/.adventurers_toolbelt/out.log \
  2>> ~/.adventurers_toolbelt/out.log


./adventurers_tools \
  >> ~/.adventurers_toolbelt/out.log \
  2>> ~/.adventurers_toolbelt/out.log &

sleep 1 

open "http://localhost:8080/"
