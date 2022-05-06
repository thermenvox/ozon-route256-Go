# https://cups.online/ru/tasks/1220

#! /bin/bash
read str

length=$(echo "${str}" | tr -cd " " | wc -m)
g="$(echo -e "${length}" | tr -d ' ')"
length=$(($g+1))

echo "Квадрат со стороной $length"
for ((i = 0; i < $length; i++))
do
echo $str
done