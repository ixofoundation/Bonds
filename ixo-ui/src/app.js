import {inject} from 'aurelia-framework';
import {PLATFORM} from 'aurelia-pal';
import Web3 from 'web3';
import {Appstate} from 'appstate';


@inject( Appstate )
export class App {
  constructor( state ) {
    this.state = state;

    if ( this.state.blockchainProviders.metamask.doShow ) {
      this.initProvider( this.state.blockchainProviders.metamask );
      //TODO requestInfo();
    }
    if ( this.state.blockchainProviders.ixo_keysafe.doShow ) {
      this.initProvider( this.state.blockchainProviders.ixo_keysafe );
      this.requestInfo();
    }
  }


  configureRouter(config, router) {
    this.router = router;
    config.title = '';
    config.map([
      {
        route: ['', 'overview'],
        name: 'overview',
        moduleId: PLATFORM.moduleName('resources/elements/overview'),
        nav: true,
        title: 'Overview'
      },
      {
        route: 'exchange',
        name: 'exchange',
        moduleId: PLATFORM.moduleName('resources/elements/exchange'),
        nav: true,
        title: 'Exchange'
      },
      {
        route: 'orders',
        name: 'orders',
        moduleId: PLATFORM.moduleName('resources/elements/orders'),
        nav: true,
        title: 'My Orders'
      },
      {
        route: 'keysafe-tester',
        name: 'keysafe-tester',
        moduleId: PLATFORM.moduleName('resources/elements/keysafe-tester'),
        nav: true,
        title: 'KS Tester'
      }
    ]);
    this.router = router;
  }


  initProvider( blockchainProvider ) {
    if ( !window[blockchainProvider.windowKey] ) {
      blockchainProvider.doShow = false;
      window.alert(`Please install ${blockchainProvider.extension} first.`);
    } else {
      if ( !blockchainProvider.provider ) {
        if ( blockchainProvider.id === this.state.blockchainProviders.metamask.id ) {
          blockchainProvider.provider = new Web3( window[blockchainProvider.windowKey].currentProvider );
        } else if ( blockchainProvider.id === this.state.blockchainProviders.ixo_keysafe.id ) {
          const IxoKeysafeInpageProvider = window[blockchainProvider.windowKey]
          blockchainProvider.provider = new IxoKeysafeInpageProvider();
        }
      }
    }
  }


  requestInfo() {
    this.state.blockchainProviders.ixo_keysafe.provider.getInfo( (error, response) => {
      if ( error ) {
        console.log( `error: ${error}` );
        return;
      }
      // console.log(`info: ${JSON.stringify(response)}, error: ${JSON.stringify(error)}`)
      this.state.keySafeInfo.username = response.name;
      this.state.keySafeInfo.did = response.didDoc.did;
      this.state.keySafeInfo.pubkey = response.didDoc.pubKey;

      /* EXAMPLE
        {
          "didDoc": {
            "did": "did:sov:BhHF1yt33YVivywggsKZ4k",
            "pubKey": "6q5GvVbsarDupenM8hmJugjy3yqyRPAAT2ixoQ6XCBuL"
          },
          "name": "Your Account Name"
        }
      */
    });
  }
}
