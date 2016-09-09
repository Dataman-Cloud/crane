/**
 * @description
 * if attrs['dmDisplayComponents'] value not includes $rootScope.componentList, the element will be hide.
 * eg. registry,search,logging and more ...
 * @example
 <md-autocomplete
 dm-display-components="search"
 md-no-cache="false"
 md-search-text="rootCtrl.searchText"
 md-items="item in rootCtrl.querySearch(rootCtrl.searchText)"
 md-item-text="item.Name ? item.Name : item.ID"
 md-min-length="1"
 placeholder="输入关键字查询"
 md-delay="500"
 md-selected-item-change="rootCtrl.searchJump(item)" class="search-header">
 </md-autocomplete>
 *
 * @example
 <li title="镜像仓库" dm-display-components="registry">
 <md-button data-ui-sref="registry" data-ui-sref-opts="{inherit: false}"><i
 class="fa fa-university"></i>镜像仓库
 </md-button>
 </li>
 */
(function () {
    'use strict';
    angular.module('app')
        .directive('dmDisplayComponents', dmDisplayComponents);

    /* @ngInject */
    function dmDisplayComponents($rootScope) {
        return {
            restrict: 'A',
            link: link
        };

        function link(scope, elem, attrs) {
            var component = attrs['dmDisplayComponents'];

            scope.$watch(function () {
                    return $rootScope.componentList
                },
                function (newValue, oldValue) {
                    if (newValue && angular.isArray(newValue)) {
                        newValue.indexOf(component) === -1 ? elem.hide() : elem.show();
                    }
                })
        }
    }
})();