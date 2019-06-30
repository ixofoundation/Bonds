# Cosmos SDK Application

This Cosmos SDK application contains the Cosmic Bonding module under ```./x/pricing/```.

The application is based on the nameservice app provided in the [Cosmos sdk-application-tutorial repository](https://github.com/cosmos/sdk-application-tutorial). As a result, it has both the functionality of the nameservice app and the Cosmic Bonding module and uses the nameservice's ```nsd``` and ```nscli``` commands. It will be made independent of the nameservice app in the future.

**Note**: Requires [Go 1.12+](https://golang.org/dl/)

## Building and Running

To build and run the application:
```bash
make run
```

To run a demo:
```bash
make demo
```

The demo consists of:
- Cosmic Bond creation
- Cosmic Bond querying
- A mix of buys and sells