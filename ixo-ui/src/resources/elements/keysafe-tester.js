import {inject} from 'aurelia-framework';
import {Appstate} from 'appstate';
import { Web } from '../../lib/web';

@inject( Appstate, Web )
export class KeysafeTester {
  msgBody = undefined;

  constructor( state, web ) {
    this.state = state;
    this.web = web;
  }


  signContent() {
    console.log( 'about to sign content')

    this.state.blockchainProviders.ixo_keysafe.provider.requestSigning( this.msgBody, (error, response) => {
      if( error ) {
        console.log( `error: ${error}` );
        return;
      }
      
      this.signingResponse = response;
      /* EXAMPLE
      {
        "type": "ed25519-sha-256",
        "created": "2018-06-07T14:51:37Z",
        "creator": "did:sov:BhHF1yt33YVivywggsKZ4k",
        "publicKey": "52PTt1eA5gGSiXBuoNwtGrN3p52XKTHb4ayer48MCahR",
        "signatureValue": "B59D2CA3B084C1DE38E08627815AE62EE7DC03E466688267BCACA04B61040DDF8DCDB9CFC713D4B9694B5499281F9ACFE734C663A91E17CA48335F9CC8B58704"
      }
      */
    });
  }


  async buy() {
    await this.web.doGetOne();
    // let result = await this.web.doPost();
  }
}
