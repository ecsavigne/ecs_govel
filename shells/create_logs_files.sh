#!/bin/bash
chmod -R 777 /var/log/whatsappx

for i in "13301" "13302" "13303" "13304" "13305" "13306" "13307" "13308" "13309" "13310" "13311" "13312" "13313" "13314" "13315" "13316" "13317" "13318" "13319" "13320"
do
    echo $i
    # echo -n  "" > /var/log/whatsappx/main"$i"-general-log.log
    echo -n  "" > /var/log/whatsappx/main"$i".log
    # sudo > /var/log/whatsappx/main"$i"-general-log.log
    # sudo > /var/log/whatsappx/main"$i".log
done

chmod -R 777 /var/log/whatsappx
