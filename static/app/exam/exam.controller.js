function ExamCtrl($scope, Restangular) {
	var that = $scope;

	Restangular.one("authorized", "listOfUnfinishedExams").get().then(function (data) {
		that.listOfUnfinishedExams = data;
	});

	Restangular.one("authorized", "listOfFinishedExams").get().then(function (data) {
		that.listOfFinishedExams = data;
	});

	$scope.createNewExam = function () {
		console.log("niveau: " + $scope.niveau);
		Restangular.one("authorized", "createNewExam").one($scope.niveau).get().then(function () {
		});
	};

	// <Mo0O>	et post les reponse une fois validé par le user
	// POST /user/:id/exem/:id/response
}