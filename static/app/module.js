angular.module("loginHabilitation", ["ui.router", "restangular", "ngMaterial"])
  .run(runBlock)
  .config(Config)
  .config(Route)
  .controller("loginHabilitationCtrl", LoginHabilitationCtrl);
