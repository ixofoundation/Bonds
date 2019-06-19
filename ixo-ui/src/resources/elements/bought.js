import { inject } from 'aurelia-framework';
import { ChainService } from 'services/chain-service';

@inject( ChainService )
export class Bought {
  maxOffer = undefined;
  numOfTokens = 200;
  totalPaymentAmount = 52400;
  pricePerEDU = 262;

  constructor( svcChain ) {
    this.svcChain = svcChain;
  }

  submitOrder() {
    const order = {
      'numOfTokens': this.numOfTokens,
      'maxOffer': this.maxOffer,
      'totalPaymentAmount': this.totalPaymentAmount,
      'pricePerEDU': this.pricePerEDU
    };
    this.svcChain.orderNow( order );
  }

}
