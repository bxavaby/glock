> [!WARNING]
> Have yet to fully test the cli. <br>
> No key elements missing, but it might not look that pretty.

#

<div align="center">

<br>

![Deps](https://img.shields.io/badge/1-deps?style=plastic&label=deps&labelColor=000000&color=00ADD8)
[![Last Commit](https://img.shields.io/github/last-commit/bxavaby/glock?style=plastic&labelColor=000000&color=00ADD8)](https://github.com/bxavaby/glock/commits/main)

───────────────────

### **glock** is a portmanteau that simply pairs _go_ with _block_,
### ergo consisting of a minimal blockchain implementation.

<br>

<img src="assets/help.png" width="643" alt="glock help" />

<br>

</div>

## Plan && features:

- block structure w/ *hash, prev. hash, timestamp, data, nonce*
- **PoW** mining w/ adjustable difficulty
- chain validation == integrity
- add blocks
- **SHA-256** hash calculation

<br>

## Command-line

```
glock init *initializes the blockchain w/ genesis*
```

```
glock add "data" *mines and adds a new block w/ custom data*
```

```
glock print *prints the entire blockchain*
```

```
glock validate *checks blockchain integrity*
```

```
glock stats *shows chain stats (length, total difficulty, avg. mining time)*
```

<br>

> [!NOTE]
> 
> This educational blockchain implementation uses PoW (Proof-of-Work) consensus to demonstrate fundamental concepts. While Ethereum transitioned to PoS (Proof-of-Stake) in 2022, understanding PoW remains a must for blockchain devs.

<br>

<div align="center">
  
───────────────────

*Hope this will serve you somehow.*

**[Report Bug](../../issues)** | **[Suggest Feature](../../issues)**

**MIT License © 2025 bxavaby**

</div>
