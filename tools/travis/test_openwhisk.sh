#!/usr/bin/env bash

set -e

export OPENWHISK_HOME="$(dirname "$TRAVIS_BUILD_DIR")/incubator-openwhisk";
HOMEDIR="$(dirname "$TRAVIS_BUILD_DIR")"
cd $HOMEDIR

# Clone the OpenWhisk code
git clone --depth 3 https://github.com/apache/incubator-openwhisk.git

# Clone the OpenWhisk CLI code
git clone --depth 3 https://github.com/apache/incubator-openwhisk-cli.git

# Build script for Travis-CI.
WHISKDIR="$HOMEDIR/incubator-openwhisk"
OPENWHISK_CLI_BUILD_DIR="$HOMEDIR/incubator-openwhisk-cli"

cd $WHISKDIR
./tools/travis/setup.sh

ANSIBLE_CMD="ansible-playbook -i environments/local -e docker_image_prefix=openwhisk"

cd $WHISKDIR/ansible
$ANSIBLE_CMD setup.yml
$ANSIBLE_CMD prereq.yml
$ANSIBLE_CMD couchdb.yml
$ANSIBLE_CMD initdb.yml
$ANSIBLE_CMD apigateway.yml

cd $OPENWHISK_CLI_BUILD_DIR
TERM=dumb ./gradlew buildBinaries

cd $WHISKDIR/ansible
$ANSIBLE_CMD wipe.yml
$ANSIBLE_CMD openwhisk.yml -e openwhisk_cli_home=$OPENWHISK_CLI_BUILD_DIR

# Install the dependencies for openwhisk cli and build the binary based on the current changes.
cd $OPENWHISK_CLI_BUILD_DIR
go get -d -t ./...
go build -ldflags "-X main.CLI_BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%S%:z'`" -o wsk

# Copy the binary generated into the OPENWHISK_HOME/bin, so that the test cases will run based on it.
mkdir -p $WHISKDIR/bin
cp $OPENWHISK_CLI_BUILD_DIR/bin/wsk $WHISKDIR/bin

# Run the test cases under openwhisk to ensure the quality of the binary.
cd $OPENWHISK_CLI_BUILD_DIR
./gradlew :tests:test -Dtest.single=Wsk*Tests*
sleep 30
./gradlew tests:test -Dtest.single=*ApiGwRoutemgmtActionTests*
sleep 30
./gradlew tests:test -Dtest.single=*ApiGwTests*
sleep 30
./gradlew tests:test -Dtest.single=*ApiGwEndToEndTests*
