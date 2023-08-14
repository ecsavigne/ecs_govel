#!/bin/bash
echo "Updating Whatsapp Service"

# git add . && git commit -m "a" && git pull origin develop
# git add . && git commit -m "a" && go build main.go
echo "go build main.go"
go build main.go
sleep 1

# exit

main_file="main"

for i in "13399"
do
    echo "killing $main_file$i"
    sudo pkill "$main_file$i"
    echo "copying $main_file to binaries/$main_file$i"
    sudo cp $main_file binaries/$main_file$i
done
