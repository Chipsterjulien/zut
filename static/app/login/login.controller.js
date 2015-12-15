// Fonction qui permet de se logger
function LoginHabilitationCtrl($scope, $rootScope, $state, Restangular) {
	var that = $scope;

	$scope.validateLogin = function () {
		$rootScope.identifiant = that.identifiant;
		$rootScope.password = that.password;
		
		Restangular.one("authorized").get().then(function() {
			$state.go("exam");
		});
	};
}