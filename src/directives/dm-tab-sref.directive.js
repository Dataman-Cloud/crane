(function () {
    'use strict';
    angular.module('glance')
        .directive('dmTabSref', dmTabSref);

    /* @ngInject */
    function dmTabSref($state, $mdUtil) {
        return {
            require: '^^mdTabs',
            restrict: 'A',
            link: link
        };

        function link(scope, elem, attrs, tabsCtrl) {
            var sref = attrs['dmTabSref'];
            var index = elem.index();
            if (!tabsCtrl.srefs) {
                tabsCtrl.srefs = [];
            }
            tabsCtrl.srefs[index] = sref;
            if ($state.includes(sref)) {
                tabsCtrl.selectedIndex = index;
            }
            if (!tabsCtrl.oldSelect) {
                tabsCtrl.oldSelect = tabsCtrl.select;
            }

            tabsCtrl.select = function (index, canSkipClick) {
                var sref = tabsCtrl.srefs[index];
                tabsCtrl.oldSelect(index, canSkipClick);
                $state.go(sref, null, {reload: true});
            };

//            scope.$on('$stateChangeSuccess',
//                    function (event, toState, toParams, fromState, fromParams) {
//                        console.log("my sref:", sref)
//                        if ($state.includes(sref)) {
//                            $mdUtil.nextTick(function () {
//                                tabsCtrl.oldSelect(index);
//                            });
//                        }
//                    });
        }
    }
})();