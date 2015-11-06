angular.module("appHabilitation", ["ui.router", "restangular", "ngMaterial"])
  .run(runBlock)
  .config(Config)
  .config(Route)
  .controller("appHabilitationCtrl", LoginHabilitationCtrl)
  .controller("exemCtrl", ExemCtrl);