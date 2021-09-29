package blockchain



type Block struct {
    Hash        []byte // Derived from Data and PrevHash
    Data        []byte // Can be passed around in a dist db
    PrevHash    []byte // Last block's hash; links blocks together
}