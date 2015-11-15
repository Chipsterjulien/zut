function Config ($stateProvider, $urlRouterProvider, RestangularProvider) {

	// Configuration de restangular
	// L'adresse de base est localhost:8080/
	RestangularProvider.setBaseUrl("localhost:8080/");
}