#!/bin/bash
grep "Failed" /var/log/auth.log | grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}' | \
sort | uniq -c | sort -n | awk '{if ($1 > 100) print $2}' | xargs -I {} \
iptables -A INPUT -s {} -j DROP