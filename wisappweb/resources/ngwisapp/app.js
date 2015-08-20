(function(){
	console.log('Initializing Angular Application');	
	angular.module('wisapp', ['ngRoute','registration']);

	
	angular.module('wisapp').config(['$routeProvider', '$locationProvider', function ($routeProvider,$locationProvider) {
		$locationProvider.html5Mode(true);
		$routeProvider
  		.when('/register',{
  			templateUrl:'resources/ngwisapp/components/registration/register.html',
  			controller:'RegistrationController'
  		})
  		.when('/login',{
  			templateUrl:'resources/ngwisapp/components/basetemplate/logintemplate.html'
  		})
  		.when('/home',{
  			templateUrl:'resources/ngwisapp/components/basetemplate/hometemplate.html'
  		})
  		.otherwise({redirectTo:'/home'});	
	}]);
	

})();
	







	