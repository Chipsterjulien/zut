// Fonction qui permet de se logger
function LoginHabilitationCtrl($scope, $rootScope, $state, Restangular) {
	// Chipster1: tu fais ça avec retangular: Restangular.addResponseInterceptor (pour verif le code de retour du serveur), et Restangular.addFullRequestInterceptor pour set le header avec les credential a chaque requests

  var that = $scope;

	$scope.validateLogin = function () {
		$rootScope.identifiant = that.identifiant;
		$rootScope.password = that.password;
		$state.go("exem");
	};
}