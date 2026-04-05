#!/bin/bash

setsid kitty sh -c "cd ~/Projects/mme/ManagementMadeEazy/server && go run main.go; exec bash" &
setsid kitty sh -c "cd ~/Projects/mme/ManagementMadeEazy/client && npm run dev; exec bash" &
