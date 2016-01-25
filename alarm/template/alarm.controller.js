function AlarmCtrl($scope, Restangular, $state) {
	var that = $scope;

	ipMap = {
		//salon: "localhost:8090",
		"Salon": "192.168.1.6:8090",
		"Balcon": "192.168.1.13:8090",
	};

	that.items = {};

	for (var name in ipMap) {
		// console.log("État avant promise de \"" + name + "\": " + that.name)
		Restangular.oneUrl(name, "http://" + ipMap[name] + "/api/v1/stateAlarm").get().then(function (data) {
			var location = data.location;
			that.items[location] = data;
			// console.log("Location: " + data.location);
			// console.log("État: " + data.state);
			// console.log("État enregistré: " + that.items[location].state);
			// console.log("État via scope: " + $scope.items[location].state);
		});
	}

	that.startCam = function (nomCam) {
		Restangular.oneUrl(nomCam, "http://" + ipMap[nomCam] + "/api/v1/startAlarm").put().then(function (data) {
			$state.reload();
		});
	};

	that.stopCam = function (nomCam) {
		Restangular.oneUrl(nomCam, "http://" + ipMap[nomCam] + "/api/v1/stopAlarm").put().then(function (data) {
			console.log("C'est bon, c'est arrêté !");
			$state.reload();
		});
	};

	that.takePicture = function (nomCam) {
	};
}

angular.module("appAlarm", ["ui.router", "restangular", "ngMaterial"])
	.config(function ($stateProvider, $urlRouterProvider, RestangularProvider) {
		$urlRouterProvider.otherwise("/alarm");
		$stateProvider
			.state("alarm", {
				url: "/alarm",
				templateUrl: "template/alarmTemplate.html",
				controller: "alarmCtrl",
				controllerAs: "alarm"
			});
	})
	.controller("alarmCtrl", AlarmCtrl);