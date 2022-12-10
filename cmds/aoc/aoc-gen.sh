#!/bin/bash
YEARS=`/bin/ls -d ../../y2* | sed -e 's/^.*y/y/'`
echo "YEARS: $YEARS" > tt.go
for y in "$YEARS"; do
    DAYS=`/bin/ls ../../$y/day?.go ../../$y/day??.go`
    echo "DAYS: $DAYS" >> tt.go
done
