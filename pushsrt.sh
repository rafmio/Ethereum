#!/bin/bash
git pull > /dev/null
STATUS=$(git status)
UPTODATE="nothing to commit, working tree clean"
if [[ $STATUS == *"$UPTODATE"* ]]; then
	echo "All data is updated" > /dev/null
else
	git add .
	DATETIME=$(date+'%d.%m.%Y %H:%M')
	STRINGTOCOMMIT="Home PC $DATETIME"
	git commit -m "$STRINGTOCOMMIT"
	git push
fi

