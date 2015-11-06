function Route ($stateProvider, $urlRouterProvider, RestangularProvider) {

  // Configuration de uirouter
  $urlRouterProvider.otherwise("/login");
  $stateProvider
    .state("login", {
      url: "/login",
      templateUrl: "./app/login/login.html",
      controller: "loginHabilitationCtrl",
      controllerAs: "login"
      // controllerAs permet de raccourcir le code dans le template. Je peux écrire login au lieu de loginHabilitationCtrl
  })
    .state("exem", {
      url: "/exem",
      templateUrl: "./app/exem/exem.html",
      controller: "exemCtrl",
      controllerAs: "exem"
  });
}