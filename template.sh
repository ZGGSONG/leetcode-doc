#!/bin/bash

# 中文名
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m填入中文名(121.买卖股票的最佳时机)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read INDEX CNNAME
if [[ -z $CNNAME || -z $INDEX ]]; then
  INDEX="000."
  CNNAME="默认题目"
fi
echo -e "\033[42;37m当前文件名: ${INDEX} ${CNNAME} \033[0m"

if [ ! -d "./code/${INDEX}${CNNAME}" ]; then
  mkdir ./code/${INDEX}${CNNAME}
  echo > ./code/${INDEX}${CNNAME}/main.go
fi

# 分组
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m填入分组名(array)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read FILEPATH
if [[ -z $FILEPATH ]]; then
  FILEPATH="default"
fi
echo -e "\033[42;37m当前分组: ${FILEPATH} \033[0m"

if [ ! -d "./docs/${FILEPATH}" ]; then
  mkdir ./docs/${FILEPATH}
fi

# 英文名
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m填入文件名(0118.pascals-triangle.md)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read FILENAME
if [[ -z $FILENAME ]]; then
  FILENAME="0000.default.md"
fi
echo -e "\033[42;37m当前文件名: ${FILENAME} \033[0m"

# 追加信息到_sidebar.md
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m是否追加到侧边栏(y/n)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read isSideBar
if [[ $isSideBar == "Y" || $isSideBar == "y" ]]; then
  echo >> _sidebar.md
  echo "  * [${INDEX} ${CNNAME}](docs/${FILEPATH}/${FILENAME})" | awk '{printf $0}' >> _sidebar.md
  echo -e "\033[42;37m已添加至侧边栏目录 \033[0m"
else
	echo -e "\033[41;36m取消添加至侧边栏 \033[0m"
fi

# 创建新文件
file_text=$(< ./template.md)
echo "${file_text}" > ./docs/${FILEPATH}/${FILENAME}

echo "Successfully created docs/${FILEPATH}/${FILENAME}"

exit 0