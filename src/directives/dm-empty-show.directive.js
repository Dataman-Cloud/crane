(function () {
    'use strict';
    angular.module('app')
        .directive('dmEmptyShow', dmEmptyShow);

    /* @ngInject */
    function dmEmptyShow(utils) {
        return {
            restrict: 'A',
            link: link,
            scope: {
                value: '=dmEmptyShow'
            }
        };

        function link(scope, elem) {
            if (utils.isEmpty(scope.value)) {
                elem.replaceWith('<div class="alert-noinfo">暂无信息</div>');
            }
        }
    }
})();