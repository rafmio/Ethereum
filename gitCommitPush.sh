#!/bin/bash
git status
git add
HOMEIP="87.117.185.196"
CURRENTIP=$(wget -qO- eth0.me)
echo $IP
if["$HOMEIP" -eq "$CURRENTIP"];then
  echo "IP address againg: $IP"
fi
