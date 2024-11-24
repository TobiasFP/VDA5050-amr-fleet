#!/bin/bash
#
# mqtt_data.sh - send data to MQTT broker
#
# sudo apt install mosquitto-clients
#
# Get Data values
idle=$(vmstat | awk '{ if (NR==3) print $15}')
used=$(df | awk '{if (NR==4) printf "%.f\n", $5 }')
space=$(df | awk '{if (NR==4) printf "%.1f\n", $4/1000 }')
 
echo " $idle $used $space"
 
# Publish Data
server="localhost"
pause=2
 
mosquitto_pub -h $server -t rtr_idle -m $idle
sleep $pause
mosquitto_pub -h $server -t rtr_used -m $used
sleep $pause
mosquitto_pub -h $server -t rtr_space -m $space