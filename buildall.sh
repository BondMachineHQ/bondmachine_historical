#!/bin/bash

PREFIX=`pwd`

BN=`basename $PREFIX`
DN=`dirname $PREFIX`

BUILD="false"

while [ 1 ] 
do
	if [ $BN == "bondmachine" ]
	then
		PREFIX=$DN
		BUILD="true"
		break
	fi
	if [ $BN == "/" ]
	then
		break
	fi
	PREFIX=$DN
	BN=`basename $PREFIX`
	DN=`dirname $PREFIX`
done

if [ $BUILD == "false" ]
then
	exit
fi

export GOPATH=$PREFIX/bondmachine/bondgo
cd $GOPATH/src
(
go build bondgo_main.go && (
sudo mv bondgo_main /usr/local/bin/bondgo
sudo chown root:root /usr/local/bin/bondgo
sudo chmod a+rx /usr/local/bin/bondgo
echo -n "Bondgo "
echo -e "\033[32m[ Ok ]\033[0m" ) || echo -e "\033[31m[ Failed ]\033[0m"
) &
cd - > /dev/null

export GOPATH=$PREFIX/bondmachine/procbuilder
cd $GOPATH/src
(
go build procbuilder_main.go && (

sudo mv procbuilder_main /usr/local/bin/procbuilder
sudo chown root:root /usr/local/bin/procbuilder
sudo chmod a+rx /usr/local/bin/procbuilder
echo -n "Procbuilder "
echo -e "\033[32m[ Ok ]\033[0m" ) || echo -e "\033[31m[ Failed ]\033[0m"
) &
cd - > /dev/null

export GOPATH=$PREFIX/bondmachine/procbuilder
cd $GOPATH/src
(
go build bondmachine_main.go && (

sudo mv bondmachine_main /usr/local/bin/bondmachine
sudo chown root:root /usr/local/bin/bondmachine
sudo chmod a+rx /usr/local/bin/bondmachine
echo -n "Bondmachine "
echo -e "\033[32m[ Ok ]\033[0m" ) || echo -e "\033[31m[ Failed ]\033[0m"
) &
cd - > /dev/null

for job in `jobs -p`
do
	wait $job
done
