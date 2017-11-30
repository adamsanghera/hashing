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

# Example workflow – Securing an asset, storing the secure version

1.  Receive user registration info, including username and password.
2.  Run `salt, hash := hashing.WithNewSalt(password)`
3.  Store the salt and hash in your database however you like.  Popular options are to concatenate hash+salt, or to store them separately.

# Example workflow – Retrieving a secureAsset, unlocking it, comparing it to a challenge.

1.  Receive a user challenge, including `username` and `password`.
2.  Retrieve the salt and hash associated with `username` from your database.
3.  `hashing.IsValidChallenge(password, salt, hash)` will return `true` if the challenge matches the underlying asset.