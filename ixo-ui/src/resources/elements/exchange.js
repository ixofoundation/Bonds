import { inject } from 'aurelia-dependency-injection';

export class Exchange {
  configureRouter(config, router) {
    config.options.eagerLoadAll = true;
    config.map([
      { route: '', redirect: 'buy' },
      { route: 'buy', name: 'buy',  moduleId: PLATFORM.moduleName('resources/elements/buy'), nav: true, title: 'Buy' },
      { route: 'bought', name: 'bought',  moduleId: PLATFORM.moduleName('resources/elements/bought'), nav: false, title: 'Bought' },
      { route: 'sell', name: 'sell',  moduleId: PLATFORM.moduleName('resources/elements/sell'), nav: true, title: 'Sell' },
      { route: 'swap', name: 'swap',  moduleId: PLATFORM.moduleName('resources/elements/swap'), nav: true, title: 'Swap' },
    ]);
    this.router = router;
  }

}
