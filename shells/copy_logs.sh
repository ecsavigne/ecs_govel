#!/bin/bash

#  adcionar ao crontab: 
#  0 0 */3 * *  bash  /home/edson/go/src/github.com/app-socialhub-pro/new_whatsmeow/utils/copy_logs.sh

# caminho e nome da pasta a ser criada
bkpFolder="/var/log/whatsmeow/whatsappx-$(date +%Y-%m-%d)"

# cria a pasta
echo $bkpFolder | xargs mkdir -p  

# copia o conteudo para pasta
echo $bkpFolder | xargs cp -ar /var/log/whatsappx/*

# esvazia os arquivos originais main1337* x1 to x9
for i in "1" "2" "3" "4" "5" "6" "7" "8" "9"
do 
  > /var/log/whatsappx/main1337$i.log
  > /var/log/whatsappx/main1337$i-general-log.log 
done

# esvazia os arquivos originais main133**
for i in "10" "11" "12" "13" "14" "15" "16" "17" "18" "19" "20" "21" "22" "23" "24" "25" "26" "27" "28" "29" "30" "31" "32" "33" "34" "35" "36" "37" "38" "39" "40" "41" "42" "43" "44" "45" "46" "47" "48" "49" "50" "51" "52" "53" "54" "55" "56" "57" "58" "59" "50"
do 
  > /var/log/whatsappx/main133$i.log
  > /var/log/whatsappx/main133$i-general-log.log 
  # truncate -s 0 /var/log/whatsappx/main133$i.log
  # truncate -s 0 /var/log/whatsappx/main133$i-general-log.log 
done