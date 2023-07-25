#!/bin/bash

for file in `ls *.abi`; do
    abigen --abi $file --pkg contract --type `basename $file .abi` --out `basename $file .abi`.go
done
