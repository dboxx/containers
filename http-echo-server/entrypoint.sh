#!/bin/sh

APP_NAME=${APP_NAME:-http-echo-server}
APP_ARCH=$(arch)

${APP_NAME}-${APP_ARCH}
