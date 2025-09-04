#!/bin/bash
echo "Как вас зовут?"
read name
mkdir $name 
echo "Привет, $name! Это твоя первая папка" > "$name"/welcome.txt
