(function(){

	angular.module('registration', ['ngResource']);

  angular.module('registration').factory('Profile', function($resource) {
    var resource = $resource('/api/profiles/:id');
    return resource;
});
  
angular.module('registration').controller('RegistrationController',function(Profile){
     var self=this;

     self.registerUSer=function(){
       console.log('Registering new User');
       var profile = new Profile();
       profile.username=self.username;
       profile.email=self.email;
       profile.password=self.password;
       profile.$save(function(){

        });
       console.log(profile);

      }
 
});

   
})();