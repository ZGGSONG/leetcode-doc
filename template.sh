#!/bin/bash

# func name
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m填入方法名(moveZeroes)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read FUNC
if [[ -z $FUNC ]]; then
  FUNC="default"
fi
echo -e "\033[42;37m当前方法名: ${FUNC} \033[0m"

# 中文名
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m填入中文名(121. 买卖股票的最佳时机)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read INDEX CNNAME
if [[ -z $CNNAME || -z $INDEX ]]; then
  INDEX="0."
  CNNAME="默认题目"
fi
echo -e "\033[42;37m当前文件名: ${INDEX} ${CNNAME} \033[0m"


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
echo -e "\033[42;37m填入文件名(0118.pascals-triangle)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read FILENAME
if [[ -z $FILENAME ]]; then
  FILENAME="0000.default"
fi
echo -e "\033[42;37m当前文件名: ${FILENAME} \033[0m"

# test func
testFunc=$(echo ${FUNC} | awk '{for (i=1;i<=NF;i++)printf toupper(substr($i,0,1))substr($i,2,length($i))" ";printf "\n"}')

# 建立英文名code
if [ ! -d "./code/${FILENAME}" ]; then
  mkdir ./code/${FILENAME}
  echo "package ${FUNC}

func ${FUNC}() {

}
" > ./code/${FILENAME}/${FUNC}.go
# add *.test.go
  echo "package ${FUNC}

import (
  \"testing\"
)

func Test${testFunc}(t *testing.T) {

}
" > ./code/${FILENAME}/${FUNC}_test.go
fi


# 难度
echo -e "\033[42;37m---------------\033[0m"
echo -e "\033[42;37m难度(1: 简单/2: 中等/3: 困难)\033[0m"
echo -e "\033[42;37m---------------\033[0m"
read difficulty
if [[ -z $difficulty ]]; then
  difficulty="未知"
  echo -e "\033[41;36m难度设置为默认未知 \033[0m"
elif [[ $difficulty == "1" ]]; then
  difficulty="简单"
elif [[ $difficulty == "2" ]]; then
  difficulty="中等"
elif [[ $difficulty == "3" ]]; then
  difficulty="困难"
fi
echo -e "\033[42;37m难度设置为${difficulty} \033[0m"
echo ""

# 追加信息到_sidebar.md
echo >> _sidebar.md
echo "  * [${INDEX} ${CNNAME} (${difficulty})](docs/${FILEPATH}/${FILENAME}.md)" | awk '{printf $0}' >> _sidebar.md
echo -e "\033[42;37m已添加至侧边栏目录 \033[0m"

ENTITLE=`echo $FILENAME | tr '.' ' ' | awk '{print$2}'`

echo "## ${CNNAME}

LeetCode: [https://leetcode.cn/problems/${ENTITLE}/](https://leetcode.cn/problems/${ENTITLE}/)

## 题目说明



示例 1：
\`\`\`text

\`\`\`
示例 2：
\`\`\`text

\`\`\`

## 解题思路

- "> ./docs/${FILEPATH}/${FILENAME}.md

echo "Successfully created docs/${FILEPATH}/${FILENAME}.md"

exit 0