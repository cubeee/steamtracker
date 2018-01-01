#!/bin/bash
go install ./shared && go build -o /go/bin/web ./web && chmod +x /go/bin/web && /go/bin/web