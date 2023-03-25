#!/bin/bash
git pull
STATUS=$(git status)
UPTODATE="nothing to commit, working tree clean"
if [[ $STATUS == *"$UPTODATE"* ]]; then
	echo "All data is updated"
fi

