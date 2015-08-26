(function(){
	angular.module('shared', []);

	angular.module('shared').controller('SharedController',function(){

		var self = this;
		
		self.isAuthenticated=function(){
			return false;
		}
	});

})();