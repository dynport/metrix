#!/bin/bash
set -e

# to be called with
# GOOS=linux GOARCH=amd64 bash ./scripts/release.sh

BIN_NAME=metrix

# revision and version
PROJECT_ROOT=$(grealpath $(dirname $0)/../)
GIT_COMMIT=$(git rev-parse --short HEAD)
GIT_STATUS=$(test -n "`git status --porcelain`" && echo "+CHANGES" || echo "")
GIT_REV=$GIT_COMMIT$GIT_STATUS
VERSION=$(grep VERSION $PROJECT_ROOT/constants.go | cut -d '"' -f 2)

# dirs and paths
RELEASE_PATH=$PROJECT_ROOT/releases/$GIT_REV
NAME=$BIN_NAME-v$VERSION.$GOOS.$GOARCH
RELEASE_TMP_DIR=$RELEASE_PATH/$NAME
RELEASE_BIN=$RELEASE_TMP_DIR/$BIN_NAME

echo "building in $RELEASE_PATH"
mkdir -p $RELEASE_TMP_DIR
go build -a -ldflags "-X main.GITCOMMIT $GIT_REV" -o $RELEASE_BIN
chmod a+x $RELEASE_BIN
cd $(dirname $RELEASE_BIN) && tar cfz $RELEASE_PATH/$NAME.tar.gz $BIN_NAME
rm -Rf $RELEASE_TMP_DIR
