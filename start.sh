#!/bin/bash

chmod +x hertz-admin
pkill -f hertz-admin
nohup ./hertz-admin > /dev/null 2>&1 &