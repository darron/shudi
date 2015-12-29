
trap "make consul_kill" INT TERM EXIT

export SHUDI_DEBUG=1
export HOSTNAME=$(hostname)
export COMMAND="w"
export W_SHA="50e721e49c013f00c62cf59f2163542a9d8df02464efeb615d31051b0fddc326"
export CHECK_URL="shudi/$W_SHA/$HOSTNAME"

sleep 5

T_10runOnce() {
  once="$(bin/shudi run -e $COMMAND -d 3 -s 1 --once --verbose | grep $W_SHA)"
  [[ "$?" -eq "0" ]]
}

T_20blockCommand() {
  result="$(bin/shudi block -e $COMMAND --verbose | grep 'was blocked')"
  [[ "$?" -eq "0" ]]
}

T_30lookAtBlockKey() {
  reason="$(consul-cli kv-read $CHECK_URL | grep 'No reason given')"
  [[ "$?" -eq "0" ]]
}

T_35runOnceBlocked() {
  blocked="$(bin/shudi run -e $COMMAND -d 3 -s 1 --once --verbose | grep 'skip.*noexec')"
  [[ "$?" -eq "0" ]]
}

T_40unblockCommand() {
  unblocked="$(bin/shudi unblock -e $COMMAND --verbose | grep 'was unblocked')"
  [[ "$?" -eq "0" ]]
}

T_50lookAtBlankKey() {
  emptykey="$(consul-cli kv-read $CHECK_URL)"
  [[ "$emptykey" == "" ]]
}
