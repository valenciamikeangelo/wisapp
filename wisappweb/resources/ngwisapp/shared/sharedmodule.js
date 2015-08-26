(function(){
	angular.module('shared', []);
	
	angular.module('shared').service('authService',function($http){
		var self = this;
		
		self.login=function(email,password){
			return $http.post('/auth/login',{email:email,password:password});
		}
	});

	
	angular.module('shared').controller('SharedController',function(authService){

		var self = this;
		
		self.isAuthenticated=function(){
			return false;
		}
		
		self.login=function(){
			console.log('authenticating user!');
			console.log(self.emailadd+' '+ self.password);
			authService.login(self.emailadd, self.password)
		}
		
	});

})();