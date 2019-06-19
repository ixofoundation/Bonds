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

  async sendSignedMessage( message ) {
    return await this.web.fetch( config => {
      return config.withUrl( '/txs' ).withMethod( 'POST' ).withBody( message );
    } );
  }

  async orderNow( order ) {
    throw new Error( 'Not yet implemented' );
  }
}
