#!/usr/bin/env bash
# shellcheck disable=SC1090
#while true; do
#  echo "> 0: HTTP" && echo "> 1: RPC"
#  read -rp "Please select server mode: [0/1] " slc
#  case ${slc} in
#    [0]*|'' ) echo -e "Selected HTTP.\n" && mode="http"; break;;
#    [1]* ) echo -e "Selected RPC.\n" && mode="rpc"; break;;
#    * ) echo "";;
#  esac
#done
mode="http"

while true; do
  echo "> 0: Go" && echo "> 1: Node.js" && echo "> 2: PHP"
  read -rp "Please select server language: [0/1/2] " slc
  case ${slc} in
    [0]*|'' ) echo "Running Go server..." && language=0; break;;
    [1]* ) echo "Running Node.js server..." && language=1; break;;
    [2]* ) echo "Running PHP server..." && language=2; break;;
    * ) echo "";;
  esac
done

case "$language" in
  0) cd "./go/$mode" || exit 1 ;;
  1) cd "./nodejs/$mode" || exit 1  ;;
  2) cd "./php/$mode" || exit 1  ;;
esac

./serve.sh
