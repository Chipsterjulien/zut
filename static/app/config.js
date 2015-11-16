// function Config ($stateProvider, $urlRouterProvider, RestangularProvider) {
function Config (RestangularProvider) {

	// Configuration de restangular
	// L'adresse de base est http://localhost:8080/ (toujours bien pensé à mettre http:// devant l'adresse !)
	RestangularProvider.setBaseUrl("http://localhost:8080/");
}