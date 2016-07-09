(function () {
    'use strict';
    angular.module('glance')
        .directive('triggerPrevClick', triggerPrevClick);

    /* @ngInject */
    function triggerPrevClick() {
        return {
            restrict: 'A',
            link: function(scope, element) {
                element.bind('click', function(e) {
                    angular.element(e.currentTarget).prev().trigger('click');
                });
            }
        };
    }
})();