#!/usr/bin/bash
read -p "Enter a number: " USER_INPUT
if [[ $USER_INPUT -lt 0 ]]; then
     echo "Not a valid positive number input"
elif [[ $USER_INPUT -gt 0 && $USER_INPUT -lt 10 ]]; then
     echo "Valid 1 digit number entered"
elif [[ $USER_INPUT -gt 9 && $USER_INPUT -lt 100 ]]; then
     echo "Valid 2 digit number entered"
else
     echo "A valid 2 digit number was not entered"
fi
