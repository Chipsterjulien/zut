angular.module("appHabilitation", ["ui.router", "restangular", "ngMaterial"])
  .run(runBlock)
  .config(Config)
  .config(Route)
  .controller("loginHabilitationCtrl", LoginHabilitationCtrl)
  .controller("exemCtrl", ExemCtrl);