#!/bin/bash

ulimit -c unlimited
echo /var/server/core.%e.%p.%h.%t > /proc/sys/kernel/core_pattern
env GOTRACEBACK=crash ./Server

