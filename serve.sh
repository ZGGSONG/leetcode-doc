#!/bin/sh

msg=""
if [ "${1}" == "" ]; then
  msg="update"
else
  msg="${1}"
fi
git pull

git add .

git commit -m "${msg}"

echo "commit with msg: ${msg}"

git push origin main

echo "push to server"

git push github main

echo "push to github"
