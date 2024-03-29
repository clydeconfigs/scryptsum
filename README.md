# scryptsum

compute and check scrypt key-derivation function (KDF), in the command line

options: L, N, p, r  
`L` - how many bytes to generate as output, e.g. 32 bytes (256 bits)  
`N` – iterations count (affects memory and CPU usage), e.g. 19 or 20  
`r` – block size (affects memory and CPU usage), e.g. 8  
`p` – parallelism factor (threads to run in parallel - affects the memory, CPU usage), usually 1

# how to use

## install

run `make install` (do not use as sudo)

## examples

if you want to hash the string "banana"

`echo -n "banana" | scryptsum` 

if you want to use 2GB of RAM (-N 21 means 2^21, which is 2097152)

`echo -n "banana" | scryptsum -N 21`

and other options (-L for length, -p for whatever that meant, and -r for the other thing i forgot too)

`echo -n "banana" | scryptsum -L 32 -N 19 -p 8 -r 8`
