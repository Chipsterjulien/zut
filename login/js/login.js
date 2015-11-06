// Script personnel
var app = angular.module("loginHabilitation", ["ui.router", "restangular", "ngMaterial"]).run(runBlock);

function runBlock($rootScope, Restangular) {
	Restangular.addResponseInterceptor(function (data, operation, what, url, response, deferred) {
		console.log("````````````````");
		console.log(data);
		console.log(operation);
		console.log(what);
		console.log(url);
		console.log(response);
		console.log(deferred);
		console.log("````````````````");
	});

	Restangular.addFullRequestInterceptor(function (headers, params, element, httpConfig) {
		console.log("````````````````");
		console.log(headers);
		console.log(params);
		console.log(element);
		console.log(httpConfig);
		console.log("````````````````");
		console.log("----------------");
    console.log($rootScope.identifiant);
    console.log($rootScope.password);
		console.log("----------------");
      headers.Authorization = btoa($rootScope.identifiant + ":" + $rootScope.password);
		console.log("header.Authorization: ", headers.Authorization);
	});
}

app.config(function ($stateProvider, $urlRouterProvider, RestangularProvider) {

	// Configuration de restangular
	RestangularProvider.setBaseUrl("localhost:8080/");

	// Configuration de uirouter
	$urlRouterProvider.otherwise("/login");
	$stateProvider
		.state("login", {
			url: "/login",
			templateUrl: "partials/login.html",
			controller: "loginHabilitationCtrl",
			controllerAs: "login"
			// controllerAs permet de raccourcir le code dans le template. Je peux écrire login au lieu de loginHabilitationCtrl
	});
});

// Fonction qui permet de se logger
app.controller("loginHabilitationCtrl", function ($scope, $rootScope, Restangular) {
	// Chipster1: tu fais ça avec retangular: Restangular.addResponseInterceptor (pour verif le code de retour du serveur), et Restangular.addFullRequestInterceptor pour set le header avec les credential a chaque requests

  var that = $scope;

	$scope.validateLogin = function () {
    $rootScope.identifiant = that.identifiant;
    $rootScope.password = that.password;
    Restangular.one("plop").get();
	};
});
