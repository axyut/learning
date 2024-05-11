#!/usr/bin/bash
#echo Hello World!

#variable
#myName=achyut
#echo Hello $myName

# asking for input 
#echo What is your Name?
#read firstName
#read lastName
#echo Hello $firstName $lastName

# positional argumnets
#echo What is your name ?
#read first second
#echo Welcome Mr.$second $first


read -p "Enter a number: " first
if [[ $first -lt 10 ]]; then
     echo "good"
elif [[ $first -gt 10 ]]; then
     echo "bad"
else
     echo "get out"
fi
