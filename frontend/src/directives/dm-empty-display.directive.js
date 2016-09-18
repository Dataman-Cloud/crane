/**
 * @description
 * if dm-empty-display value is invalid, the html and child html will replaced dm-empty-text value or '暂无信息'
 * @example
 <div flex="70" dm-empty-display="serviceConfigCtrl.service.Spec.Networks" dm-empty-text="{/'Not configured' | translate/}">
     <ul>
        <li data-ng-repeat="network in serviceConfigCtrl.service.Spec.Networks">{/network/}</li>
     </ul>
 </div>
 */

(function () {
    'use strict';
    angular.module('app')
        .directive('dmEmptyDisplay', dmEmptyDisplay);

    /* @ngInject */
    function dmEmptyDisplay(utils, $compile) {
        return {
            restrict: 'A',
            link: link,
            priority: 1001,
            scope: {
                value: '=dmEmptyDisplay',
                text: '@dmEmptyText'
            }
        };

        function link(scope, elem) {
            var html = '<div class="no-info">{/text ? text : "No information" | translate/}</div>';

            if (utils.isEmpty(scope.value)) {
                var e = $compile(html)(scope);
                elem.replaceWith(e);
            }
        }
    }
})();