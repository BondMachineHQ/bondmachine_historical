#!/bin/bash

BASE="test"

if [ "a$1" == "areset" ]
then
	for i in $BASE/sourc*go
	do
		echo
		echo
		echo "Resetting $i"
		TESTIN="$i"".asm"
		echo -n "   Creating $TESTIN"
		go run bondgo_main.go -input-file $i -save-assembly $TESTIN && echo -e "\033[32m\033[300C\033[10D[ Reset ]\033[0m"		
	done
	for i in $BASE/mpm_sourc*go
	do
		echo
		echo
		echo "Resetting $i"
		TESTIN="$i"".asm"
		go run bondgo_main.go -input-file $i -save-assembly $TESTIN -mpm
	done
elif [ "a$1" == "adiff" ]
then
	for i in $BASE/sourc*go
	do
		echo
		echo
		echo "Testing $i"
		TESTIN="$i"".asm"
		TESTOUT="$i"".out"
		go run bondgo_main.go -input-file $i  -save-assembly $TESTOUT
		echo -n "   Comparing $TESTIN and $TESTOUT "
		if [ "a$2" == "a+" ]
		then
			vimdiff $i $TESTIN $TESTOUT
		else
			vimdiff $TESTIN $TESTOUT
		fi
		rm -f "$TESTOUT"
	done
	for i in $BASE/mpm_sourc*go
	do
		echo
		echo
		echo "Testing $i"
		TESTIN="$i"".asm"
		TESTOUT="$i"".out"
		go run bondgo_main.go -input-file $i -save-assembly $TESTOUT -mpm
		for j in $TESTOUT*
		do
			CMP1=$TESTIN""_"`echo $j | cut -d_ -f 3`"
			CMP2=$j
			echo -n "   Comparing $CMP1 and $CMP2"
        	        if [ "a$2" == "a+" ]
        	        then
        	                vimdiff $i $CMP1 $CMP2
        	        else
        	                vimdiff $CMP1 $CMP2
	                fi
			rm -f "$j"
		done
	done
else
	for i in $BASE/sourc*go
	do
		echo
		echo
		echo "Testing $i"
		TESTIN="$i"".asm"
		TESTOUT="$i"".out"
		go run bondgo_main.go -input-file $i  -save-assembly $TESTOUT
		echo -n "   Comparing $TESTIN and $TESTOUT "
		( cmp $TESTIN $TESTOUT > /dev/null 2>&1 && echo -e "\033[32m\033[300C\033[11D[ Passed ]\033[0m" ) || (echo -e "\033[31m\033[300C\033[11D[ Failed ]\033[0m" ; exit 1)
		rm -f "$TESTOUT"
	done
	for i in $BASE/mpm_sourc*go
	do
		echo
		echo
		echo "Testing $i"
		TESTIN="$i"".asm"
		TESTOUT="$i"".out"
		go run bondgo_main.go -input-file $i -save-assembly $TESTOUT -mpm
		for j in $TESTOUT*
		do
			CMP1=$TESTIN""_"`echo $j | cut -d_ -f 3`"
			CMP2=$j
			echo -n "   Comparing $CMP1 and $CMP2"
			( cmp $CMP1 $CMP2 > /dev/null 2>&1 && echo -e "\033[32m\033[300C\033[11D[ Passed ]\033[0m" ) || (echo -e "\033[31m\033[300C\033[11D[ Failed ]\033[0m" ; exit 1)
			rm -f "$j"
		done
	done
fi
