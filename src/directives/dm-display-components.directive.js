(function () {
    'use strict';
    angular.module('app')
        .directive('dmDisplayComponents', dmDisplayComponents);

    /* @ngInject */
    function dmDisplayComponents() {
        return {
            restrict: 'A',
            link: link
        };

        function link(scope, elem, attrs) {
            var component = attrs['dmDisplayComponents'];

            scope.$watch('componentList', function (newValue, oldValue) {
                if (newValue && angular.isArray(newValue)) {
                    newValue.indexOf(component) === -1 ? elem.hide() : elem.show();
                }
            })
        }
    }
})();