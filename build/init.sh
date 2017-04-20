#!/bin/bash

# Clone the ksched repo, (for now it's a mirror repo because of access issues with private coreos repo)
cd /root/go-workspace/src/github.com/oncecloud && git clone https://github.com/oncecloud/sirius sirius
cd sirius/proto
./genproto.sh
cd ..
go build ./cmd/k8sscheduler
go build ./cmd/podgen
