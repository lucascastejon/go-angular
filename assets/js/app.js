(function() {
  var app = angular.module('gemStore', ["users-angular", "extends-html"]);


  app.controller('userController', function($http, $scope) {
	
	this.login = {};

	this.addUsers = function(product) {
		$http({
		  method  : 'POST',
		  url     : 'http://localhost:1357/user',
		  data    : {
		  				"username": this.login.username,
		  				"email": this.login.email,
		  				"passwd": this.login.passwd
		  			}
		});

	  this.login = {};
	};

	this.logo = function(product) {

	  this.image = logo_teeto;
	};

	var logo_teeto = "/images/teeto.png";


  });
})();