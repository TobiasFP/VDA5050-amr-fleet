#!/bin/bash
#
# mqtt_data.sh - send data to MQTT broker
#
# sudo apt install mosquitto-clients
#
 
# Publish Data
server="localhost"
 
mosquitto_sub -h $server -t "vda5050/+/+/state"
