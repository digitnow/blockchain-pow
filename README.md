Based on https://github.com/tensor-programming/golang-blockchain/tree/master

Jeg bruker guidene fra "tensor-programming" for å teste en lokal implementasjon av grunnleggende strukturer og metoder av blockchain. 

# version01: grunnleggende struktuerer (lenket liste; genesis block; generering av hash)
* (14:49) https://www.youtube.com/watch?v=mYlHT9bB6OE
* https://steemit.com/utopian-io/@tensor/building-a-blockchain-with-go---go-modules-and-a-basic-blockchain---part-1 
```
$ go run main.go
------------------------------------
Previous Hash: 
Data in Block: Genesis
Hash: 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5

------------------------------------
Previous Hash: 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5
Data in Block: Block #1
Hash: adea772f15a7a2ddf41265f943749960c7e9f59506cd212eb079569a633a08f7

------------------------------------
Previous Hash: adea772f15a7a2ddf41265f943749960c7e9f59506cd212eb079569a633a08f7
Data in Block: Block #2
Hash: 32e79dc8f2fc481c9bb58c875bef5a9f806ae2dfd7de90f13a08c89e5cf1216f

------------------------------------
Previous Hash: 32e79dc8f2fc481c9bb58c875bef5a9f806ae2dfd7de90f13a08c89e5cf1216f
Data in Block: Block #3
Hash: 7456a6eed0f3975371ef34eb60f6e3e7de719aa299dcdbc70346c359af6bf574
```
# version02 
