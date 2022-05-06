# https://cups.online/ru/tasks/1224

#! /bin sh
read str
dates=($(echo $str | tr " " "\n"))

d1=$(date -d ${dates[0]} +%s)
d2=$(date -d ${dates[1]} +%s)
echo $((($d2-$d1)/86400))