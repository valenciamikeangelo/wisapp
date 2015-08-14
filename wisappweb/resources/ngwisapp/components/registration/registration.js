(function(){

	angular.module('registration', ['ngResource']);

  

  angular.module('registration').factory('Profile', function($resource) {
    var resource = $resource('/api/profiles/:id');
    return resource;
});
  
angular.module('registration').controller('RegistrationController',function($scope,Profile){
     


  $scope.createUser=function(){
   var formuser=$scope.user;
   var profile = new Profile();
   profile.username=formuser.username;
   profile.email=formuser.email;
   profile.$save(function(){

    });

   console.log(profile);
   $scope.user=profile;
  }

  


})

   
})();