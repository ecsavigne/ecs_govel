#!/bin/bash
echo "Updating Whatsapp Service"

# git add . && git commit -m "a" && git pull origin develop
# git add . && git commit -m "a" && go build main.go
mkdir binaries

echo "go build main.go"
go build main.go
sleep 1

# exit

main_file="main"

# 1 até 20
# for i in "13301" "13302" "13303" "13304" "13305" "13306" "13307" "13308" "13309" "13310" "13311" "13312" "13313" "13314" "13315" "13316" "13317" "13318" "13319" "13320" 
# do
#     echo "killing $main_file$i"
#     sudo pkill "$main_file$i"
#     echo "copying $main_file to binaries/$main_file$i"
#     sudo cp $main_file binaries/$main_file$i
# done

# #21 até 40
# for i in "13321" "13322" "13323" "13324" "13325" "13326" "13327" "13328" "13329" "13330" "13331" "13332" "13333" "13334" "13335" "13336" "13337" "13338" "13339" "13340"
# do
#     echo "killing $main_file$i"
#     sudo pkill "$main_file$i"
#     echo "copying $main_file to binaries/$main_file$i"
#     sudo cp $main_file binaries/$main_file$i
# done

# #41 até 60
#  for i in "13341" "13342" "13343" "13344" "13345" "13346" "13347" "13348" "13349" "13350" "13351" "13352" "13353" "13354" "13355" "13356" "13357" "13358" "13359" "13360"
# do
#     echo "killing $main_file$i"
#     sudo pkill "$main_file$i"
#     echo "copying $main_file to binaries/$main_file$i"
#     sudo cp $main_file binaries/$main_file$i
# done
