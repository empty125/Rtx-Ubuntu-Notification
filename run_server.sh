#!/bin/bash
flock -n /tmp/rtx-ubuntu-notiserversi -c "python3 -m server"
