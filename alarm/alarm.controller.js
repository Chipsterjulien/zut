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

	ipMap = {
		//salon: "localhost:8090",
		balcon: "192.168.1.13:8090"
	};

	that.items = {};
	for (var name in ipMap) {
		Restangular.oneUrl(name, "http://" + ipMap[name] + "/api/v1/stateAlarm").get().then(function (data) {
			that.items[name] = data;

			console.log("data state: " + data.state);
			console.log("data error: " + data.error);
			console.log("Nom: " + name);
			console.log("État: " + (that.items[name]).state);
			console.log("Erreur: " + (that.items[name]).error);
		});
	}

	console.log("Suite");

	for (var n in that.items) {
		console.log("Nom second: " + n);
	}

	console.log("Fin")
	console.log("Nom fin: " + that.items["balcon"].state);
}