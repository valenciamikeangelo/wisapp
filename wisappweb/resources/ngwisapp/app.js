(function(){
	console.log('Initializing Angular Application');	

  


	angular.module('wisapp', ['ngRoute','shared','registration']);

  


	angular.module('wisapp').config(['$routeProvider', '$locationProvider', function ($routeProvider,$locationProvider) {
		$locationProvider.html5Mode(true);
		$routeProvider
  		.when('/register',{
  			templateUrl:'resources/ngwisapp/components/registration/register.html',
        controller:'RegistrationController',
        controllerAs:'regCont'
  		})
  		.when('/login',{
  			templateUrl:'resources/ngwisapp/components/basetemplate/logintemplate.html',
  	        controller:'SharedController',
  	        controllerAs:'sharedCont'
  		})
  		.when('/home',{
  			templateUrl:'resources/ngwisapp/components/basetemplate/hometemplate.html',
        controller:'SharedController',
        controllerAs:'sharedCont'
      })
  		.otherwise({redirectTo:'/home'});	
	}]);
	
  

})();
	







	