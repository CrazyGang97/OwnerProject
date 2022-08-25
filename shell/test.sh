#!/bin/bash

set -e

CURDIR=$(cd $(dirname $0); pwd)
echo $CURDIR
sysname=$(uname -s)
echo $sysname
