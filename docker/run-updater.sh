#!/bin/bash
clear && go install ./shared && go build -o /go/bin/updater ./updater && \
DB_HOST=postgres DB_USER=steamtracker DB_PASS=steamtracker DB=steamtracker /go/bin/updater