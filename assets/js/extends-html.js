(function() {
  var app = angular.module("extends-html", []);

  app.directive("headerUp", function() {
    return {
      restrict: 'E',
      templateUrl: "header-up.html"
    };
  });


})();
