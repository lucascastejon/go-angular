(function() {
  var app = angular.module("users-angular", []);

  app.directive("usersAngular", function() {
    return {
      restrict: 'E',
      templateUrl: "user-list.html"
    };
  });

  app.directive("formUser", function() {
    return {
      restrict: 'E',
      templateUrl: "form-user.html"
    };
  });

  app.directive("loginUser", function() {
    return {
      restrict: 'E',
      templateUrl: "login-user.html"
    };
  });

  app.directive("alertMessage", function() {
    return {
      restrict: 'E',
      templateUrl: "alert-message.html"
    };
  });

})();
