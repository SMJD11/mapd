#!/bin/bash

while IFS= read -r line; do
  if [[ -n "$line" ]]; then  # Check if line is not empty
    echo "$line" > test_params/d/LastGPSPosition # Or /dev/shm/params/d/... if on Linux
    echo "GPS Data Sent: $line"
    sleep 1 # Send GPS data every 1 second
  fi
done < gps_route.txt

echo "GPS Simulation Finished"