import { inject } from 'aurelia-framework';
import numeral from 'numeral';

@inject( numeral )
export class NumberFormatValueConverter {

  constructor() {
    this.numeral = numeral;
  }

  toView( value ) {
    if ( value === '' ) {
      return '';
    }

    let result = numeral( value ).format('0,0.00');
    return result;
  }

  fromView( formattedValue ) {
    return formattedValue;
  }
}
