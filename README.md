![](https://i.imgur.com/7YZMVTd.jpg)

# Cosmic Bonding
Cosmic Bonding is a custom [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) module that provides universal token bonding curve functions to mint, burn or swap any token in a Cosmos blockchain.

In future, once the Cosmos Inter-blockchain Communication protocol is available, this should enable cross-network exchanges of tokens at algormically determined prices.

The Cosmic Bonding module can be deployed through Cosmos Hubs and Zones to deliver applications such as:
* Automated market-makers (like [Uniswap](https://uniswap.io))
* Decentralised exchanges (like [Bancor](https://bancor.network))
* Curation markets (like [Relevant](https://github.com/relevant-community/contracts/tree/bondingCurves/contracts))
* Development Impact Bonds (like ixo alpha-Bonds)
* Continuous organisations (like [Moloch DAO](https://molochdao.com/))

> Hayek famously said that "...prices are an instrument of communication and guidance which embody more information than we directly have".

## Module functions

Any Cosmos application chain that implements the Cosmic Bonding module is able to perform functions such as:
* Issue a new token with custom parameters.
* Pool liquidity for reserves.
* Provide continuous funding.
* Automatically mint and burn tokens at deterministic prices.
* Swap tokens atomically within the same network.
* Exchange tokens across networks, with the IBC protocol.
* (Batch token transactions to prevent front-running)
* Launch a decentralised autonomous initial coin offerings ([DAICO](https://ethresear.ch/t/explanation-of-daicos/465))
* ...*other **DeFi**ant* innovations.
## Pricing algorithm libraries
The Cosmic Bonding module framework supports libraries for all types of pricing algorithms, such as:
* Exponential
* Logarithmic
* Negative exponential
* Constant product
* Positive initial price
* Quasi-polynomial
* Reserved Supply (Augmented)

Each formula is specified within the module library. 
This includes:
* Derived Mint equation
* Derived Burn equation

Updates to the module pricing functions must pass through a network governance process to update the module on all nodes, for changes to be made.
## Parameters
Each Cosmic Bond has an initial set of constant state (invariant) parameters that cannot be updated once these have been initialised, which include:
* Pricing function (the algorithm that will be used)
* Issuer
* Token name
* Token symbol
* Reserve wallet address
* Collateral wallet address
* Supply
* Initial reserve
* Exponent constant
* Slope constant
* Transaction fee rate
* Exit tax rate
* Coupon value (percentage)

When a Cosmic Bond transaction (such as buy, sell, swap) is submitted, this includes the variable parameters:
* Order quantity
* Maximum price
* Wallet address

The module returns the following parameters:
* Integral price
...

## Cloning this Repository

Run ```go get -u github.com/ixofoundation/cosmic```

## Running the App

For instructions on how to run the Cosmos SDK application containing the Cosmic Bonding module, click [here](./ixo-sdk/README.md). We also built a user interface, but it is currently not wired up to the SDK application. If you still wish to check it out, refer to the instructions [here](./ixo-ui/README.md).

- The code for the SDK application is under ```./ixo-sdk/```
- The code for the user interface is under ```./ixo-ui/```