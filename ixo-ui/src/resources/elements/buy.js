import { inject } from 'aurelia-framework';
import { ChainService } from 'services/chain-service';
import { Router } from 'aurelia-router';

@inject( ChainService, Router )
export class Buy {
  numOfTokens = undefined;
  maxOffer = undefined;
  totalPaymentAmount = 0;
  pricePerEDU = 0

  //TODO: this is just an example - read from IXO KeySafe
  signatureData = 'b89569f33f4bf7c48a73cdd2068fdd82f3549cf6e0106b441f8bb73d8d9ec8a72b4d33b4f6bed21b5692f7f7ea97ee7c5a1ae61aad76a5df3bed6b9879423903';
  pubkeyData = {
    type: 'tendermint/PubKeySecp256k1',
    value: 'D6mEF3EmiWGgBgr8UrveuataxrdRtDWfuhMTgmjHL7X4'
  };
  buyer = 'cosmos1metmk3f4xx6lq8swt7yze9hvgfeeww97ykc2dp';  //TODO where to get this info from?

  //TODO: this is just an example
  msgData = [{
    'type': 'pricing/Buy',
    'value': {
      'Moniker': 'bondtoken',
      'Buyer': this.buyer,
      'Amount': {
        'denom': 'reservetoken',
        'amount': this.numOfTokens
      },
      'MaxPrice': this.maxOffer
    }
  }]

  message = {
    tx: {
      msg: this.msgData,
      fee: {
        gas: '5000',  //TODO put slider into ui for adjusting this
        amount: [{
          denom: 'reservetoken',
          amount: '50'  //what is this?
        }]
      },
      memo: '',
      signature: {
        signature: this.signatureData,
        pub_key: this.pubkeyData,
        account_number: '0',
        sequence: '0'
      }
    },
    'mode': 'block'
  }

  constructor( svcChain, router ) {
    this.svcChain = svcChain;
    this.router = router;
  }

  async getQuote() {
    try {
      //TODO block ui

      //TODO call correct service
      // const result = await this.svcChain.quote( this.message );

      //TODO handover to 'bought' VM
      this.router.navigateToRoute( 'bought' );
    } catch ( err ) {
      console.log( err );
      throw err;
    }
  }
}
