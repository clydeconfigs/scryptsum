how to use

## install

run `make install` (do not use as sudo)

## run

`echo -n "password" | scryptsum` 

## examples

if you want to hash the string "banana"

`echo -n "banana" | scryptsum` 

if you want to use 2GB of RAM (-N 21 means 2^21, which is 2097152)

`echo -n "banana" | scryptsum -N 21`

and other options (-L for length, -p for whatever that meant, and -r for the other thing i forgot too)

`echo -n "banana" | scryptsum -L 32 -N 19 -p 8 -r 8`
