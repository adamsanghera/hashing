# Motivation

The motivation of this package is to make it brainless to store and use salted hashes in golang.

## Implications

Thus, this package will only accept as input hashes, salts, and assets that have been encoded in strings.

Furhter, this package will only output hashes, salts, and assets that are encoded in strings.

# Terminology

Asset = thing you want to hash.

Salt = random string you combine with a hash, to make it even better.

Challenge = string that you want to check the validity of.

HashedAsset = your asset in a locked box.  To open this box, you need to have a valid challenge and the Salt associated with it.