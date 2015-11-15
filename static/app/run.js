function runBlock($rootScope, Restangular) {
	Restangular.addResponseInterceptor(function (data, operation, what, url, response, deferred, $state) {
      // redirect to login page when server return 401 unauthorized
      if (response.status === 401) {
        $state.go("login");
      }
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
		// Ne faut-il pas mettre "basic " non encodé en base64 sur la ligne suivante ?
		headers.Authorization = btoa($rootScope.identifiant + ":" + $rootScope.password);
		console.log("header.Authorization: ", headers.Authorization);
	});
}
