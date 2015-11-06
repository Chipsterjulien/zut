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
