angular.module("appAlarm", ["ui.router", "restangular", "ngMaterial"])
	.config(function ($stateProvider, $urlRouterProvider, RestangularProvider) {
		$urlRouterProvider.otherwise("/alarm");
		$stateProvider
			.state("alarm", {
				url: "/alarm",
				//templateUrl: "index.html",   <---- commenter cette ligne pour éviter d'avoir 2x la page sur la même
				controller: "alarmCtrl"
			});
	})
	.controller("alarmCtrl", AlarmCtrl);

console.log("test---------------------------------");

function AlarmCtrl($scope, Restangular) {
	var that = $scope;

	console.log("----");
	console.log("coin");
	console.log("----");

	Restangular.oneUrl("balcon", "http://192.168.1.13:8090/api/v1/stateAlarm").get().then(function (data) {
		that.balcon = data;
	});
}