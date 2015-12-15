function runBlock($rootScope, Restangular, $state) {
	Restangular.addResponseInterceptor(function (data, operation, what, url, response, deferred, $rootState) {
		// redirect to login page when server return 401 unauthorized
		// if (response.status === 401) {
		// 	$rootState.go("login");
		// }
		console.log("````````````````");
		console.log("data: " + data);
		console.log("operation: " + operation);
		console.log("what: " + what);
		console.log("url: " + url);
		console.log("response: " + response);
		console.log("response status: " + response.status);
		console.log("deferred: " + deferred);
		console.log("````````````````");

		// The responseInterceptor must return the restangularized data element.
		var restElem = Restangular.restangularizeElement(null, data, what);
		restElem.fromServer = true;
		
		return restElem;
	});

	Restangular.setErrorInterceptor(function (response) {
		// redirect to login page when server return 401 unauthorized
		if (response.status === 401) {
			$state.go("login");

			return false;
		}

		return true;
	});

	// Restangular.addFullRequestInterceptor(function (headers, params, element, httpConfig) {
	Restangular.addFullRequestInterceptor(function (headers, params, element, httpConfig) {
		console.log("````````````````");
		console.log("headers: " + headers);
		console.log("params: " + params);
		console.log("element: " + element);
		console.log("httpConfig: " + httpConfig);
		console.log("````````````````");
		console.log("----------------");
		console.log("identifiant: " + $rootScope.identifiant);
		console.log("pwd: " + $rootScope.password);
		console.log("----------------");
		
		headers.Authorization = "Basic " + btoa($rootScope.identifiant + ":" + $rootScope.password);
		console.log("headers.Authorization: ", headers.Authorization);
		// http://stackoverflow.com/questions/24780067/angularjs-set-header-on-every-request
		return {
			element: element,
			params: params,
			headers: headers,
			httpConfig: httpConfig
		};
	});
}
