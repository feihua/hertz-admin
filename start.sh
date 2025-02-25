#!/bin/bash

chmod +x hertz-admin
pkill -f hertz-admin
sleep 3
nohup ./hertz-admin > /dev/null 2>&1 &