## $ glock

### Minimal blockchain implementation

> Plan && features:
> 1. block structure w/ *hash, prev. hash, timestamp, data, nonce*
> 2. **PoW** mining w/ adjustable difficulty
> 3. chain validation == integrity
> 4. add blocks
> 5. **SHA-256** hash calculation

#### CLI

```
glock init *initializes the blockchain w/ genesis*

glock add "data" *mines and adds a new block w/ custom data*

glock print *prints the entire blockchain*

glock validate *checks blockchain integrity*

glock stats *shows chain stats (length, total difficulty, avg. mining time)*
```

> This educational blockchain implementation uses PoW (Proof-of-Work) consensus to demonstrate fundamental concepts. While Ethereum transitioned to PoS (Proof-of-Stake) in 2022, understanding PoW remains a must for blockchain devs.
