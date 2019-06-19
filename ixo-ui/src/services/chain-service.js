import { inject } from 'aurelia-framework';
import { Web } from 'lib/web';

@inject( Web )
export class ChainService {
  constructor( web ) {
    this.web = web;
  }

  async getBondInfo() {
    return await this.web.fetch( config => {
      return config.withUrl( '/pricing/cosmic-bond/edu' ).withMethod( 'GET' );
    } );
  }

  async quote( message ) {
    //might be incorrect
    return await this.web.fetch( config => {
      return config.withUrl( '/txs' ).withMethod( 'POST' ).withBody( message );
    } );
  }

  async sendSignedMessage( message ) {
    //for testing purpose only
    return await this.web.fetch( config => {
      return config.withUrl( '/txs' ).withMethod( 'POST' ).withBody( message );
    } );
  }

  async orderNow( order ) {
    throw new Error( 'Not yet implemented' );
  }
}
