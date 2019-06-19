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

    const strVal = '' + value;
    if ( strVal.includes( '.' ) ) {
      return numeral( value ).format('0,0.00');
    }
    return numeral( value ).format('0,0');
  }

  fromView( formattedValue ) {
    return formattedValue;
  }
}
