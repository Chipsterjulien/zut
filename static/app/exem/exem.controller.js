function ExemCtrl($scope, Restangular) {
	var that = $scope;

	Restangular.one("authorized", "listOfUnfinishedExams").get().then(function (data) {
		that.liste = data;
	});

	$scope.getTimes = function(number) {
		return new Array(number);
	};

	// <Mo0O>	et post les reponse une fois validé par le user
	// POST /user/:id/exem/:id/response
}