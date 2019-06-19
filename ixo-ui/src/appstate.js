export class Appstate {

  blockchainProviders = {
    metamask: {id: 0, doShow: false, windowKey: "web3", extension: "Metamask", provider: null},
    ixo_keysafe: {id: 1, doShow: true, windowKey: "ixoKs", extension: "IXO Keysafe", provider: null}
  };

  keySafeInfo = {
    username: undefined,    //username
    did: undefined,         //unique userid
    pubkey: undefined       //public key of user
  }

}
