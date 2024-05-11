grep -oe "[a-zA-Z0-9._]\+@[a-zA-Z]\+.[a-zA-Z]\+" filename.txt.json.md

put to txt file
grep -oe "[a-zA-Z0-9._]\+@[a-zA-Z]\+.[a-zA-Z]\+" filename.txt.json.md > users.txt

sort
cat users.txt | sort -u > sortedusers.txt

example
curl https://abc.com/api/lists/of/passwords > pass.txt


echo jwt_1st_keys | base64 -d  => gives the decoded msg  (can also be viewd in browser jwt page)

################get words after comma
sudo grep -o "[^,]*$" creds.csv

To extract the emails:
cut -d "," -f 1 creds.csv > emails.txt

To extract the passwords:
cut -d "," -f 2 creds.csv > passwords.txt
cut -d "," -f 2 creds.csv | wc -l     // to see the lines count that is L

