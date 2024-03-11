#!/bin/bash

RUN_NAME=monitor_coin
DEPLOY_PATH=/opt/www/monitor_coin.deploy/
LOG_PATH=/opt/logs/monitor_coin/

function build() {
  export GO15VENDOREXPERIMENT=1
  export GO111MODULE="on"

  check_code

  rm -rf src.tar.gz && tar -czf src.tar.gz *
  [ ! -d "output/bin" ] && mkdir -p output/bin
  [ ! -d "output/config" ] && mkdir -p output/config
  [ ! -d "output/script" ] && mkdir -p output/script
  [ ! -d "output/src" ] && mkdir -p output/src
  mv -f ./src.tar.gz output/src
  cp -r config/* output/config
  cp -r script/* output/script/

  go build -ldflags=-s -o output/bin/$RUN_NAME
}

function local_build() {
  export GO15VENDOREXPERIMENT=1
  export GO111MODULE="on"

  check_code

  rm -rf src.tar.gz && tar -czf src.tar.gz *
  [ ! -d "output/bin" ] && mkdir -p output/bin
  [ ! -d "output/config" ] && mkdir -p output/config
  [ ! -d "output/script" ] && mkdir -p output/script
  [ ! -d "output/src" ] && mkdir -p output/src
  mv -f ./src.tar.gz output/src
  cp -r config/* output/config
  cp -r script/* output/script/

  go build -ldflags=-s -race -o output/bin/$RUN_NAME
}

function initenv() {
  [ ! -d "$DEPLOY_PATH/bin" ] && mkdir -p $DEPLOY_PATH/bin
  [ ! -d "$DEPLOY_PATH/config" ] && mkdir -p $DEPLOY_PATH/config
  [ ! -d "$DEPLOY_PATH/script" ] && mkdir -p $DEPLOY_PATH/script
  [ ! -d "$DEPLOY_PATH/src" ] && mkdir -p $DEPLOY_PATH/src
  [ ! -d "$LOG_PATH" ] && mkdir -p $LOG_PATH
}

function check_code() {
  go fmt
  go vet main.go
  check_code_by_dir .
}

function check_code_by_dir() {
  for fn in `ls $1`
  do
    if [ -d $1"/"$fn ];then
      if [ "$fn" != "output" ] && [ "$fn" != "script" ] && [ "$fn" != "config" ] && [ "$fn" != "vendor" ];then
        go fmt $1"/"$fn
        go vet $1"/"$fn
        check_code_by_dir $1"/"$fn
      fi
    fi
  done
}

function show_msg() {
  PID=$(ps -ef | grep $RUN_NAME | grep -v grep | awk '{print $2}')
  if [ -n "$PID" ]; then
      if ps -p $PID >/dev/null; then
          echo "$(date +"%Y-%m-%d %T") $RUN_NAME: running, PID=$PID"
          return 0
      else
          echo "$(date +"%Y-%m-%d %T") $RUN_NAME: not running"
      fi
  else
      echo "$(date +"%Y-%m-%d %T") $RUN_NAME: not exist"
  fi
  return 1
}

function start() {
  cp -r output/* $DEPLOY_PATH
  supervisorctl start "$RUN_NAME"

  sleep 2
  show_msg
}

function stop() {
  show_msg

  bakpath=$DEPLOY_PATH`date +%Y%m%d%H%M`/
  if [ ! -d "$bakpath" ]; then
    mkdir -p $bakpath
    cp -r $DEPLOY_PATH/bin $bakpath
    cp -r $DEPLOY_PATH/config $bakpath
    cp -r $DEPLOY_PATH/script $bakpath
    cp -r $DEPLOY_PATH/src $bakpath
    echo "backup last version to dir ( $bakpath ) done."
  else
    echo "$bakpath existed, cannot stop service."
    exit 1
  fi

  supervisorctl stop $RUN_NAME

  sleep 2
  show_msg
}

function restart() {
  stop
  start
}

function usage() {
  echo "Usage: sh $0 {start|stop|restart|build|local_start}"
  exit 1
}

function local_start() {
  local_build
  if [ $? -ne 0 ]; then
    echo "== build fail =="
  else
    echo "== build OK =="
    ./output/bin/$RUN_NAME -env=local
  fi
}

function kill() {
  pkill -9 $RUN_NAME
}

function rollback() {
  show_msg

  supervisorctl stop $RUN_NAME

  sleep 2
  show_msg

  cp -r $DEPLOY_PATH$1/* $DEPLOY_PATH
  supervisorctl start "$RUN_NAME"

  sleep 2
  show_msg

  echo "rollback to version $1 done."
}

if [ $# != 1 ] && [ $# != 2 ]; then
  usage
fi

case "$1" in
    start|stop|restart|build|local_start|kill|initenv|local_build)
      $1
      ;;
    rollback)
      $1 $2
      ;;
    *)
    usage
esac
