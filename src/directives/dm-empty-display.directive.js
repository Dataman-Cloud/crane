(function () {
    'use strict';
    angular.module('app')
        .directive('dmEmptyDisplay', dmEmptyDisplay);

    /* @ngInject */
    function dmEmptyDisplay(utils) {
        return {
            restrict: 'A',
            link: link,
            priority: 1001,
            scope: {
                value: '=dmEmptyDisplay'
            }
        };

        function link(scope, elem) {
            if (utils.isEmpty(scope.value)) {
                elem.replaceWith('<div class="alert-noinfo">暂无信息</div>');
            }
        }
    }
})();