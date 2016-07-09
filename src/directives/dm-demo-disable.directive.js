(function () {
    'use strict';
    angular.module('glance')
        .directive('demoDisable', dmDemoDisable)
        .directive('dmDemoDisable', dmDemoDisable);

    /* @ngInject */
    function dmDemoDisable($mdDialog, $window) {
        return {
            priority: -1,
            restrict: 'A',
            link: link
            }
        
        function link(scope, elem, attrs, ctrl) {
            elem.addClass('demo-hide');
            scope.$watch('isDemo', function (value) {
                if (value != undefined) {
                    elem.removeClass('demo-hide');
                    if (value){
                        elem.removeAttr('onclick');
                        elem.attr('href', '#');
                        elem.off('click');
                        elem.unbind('click');
                        elem.bind('click', function (e) {
                            e.preventDefault();
                            e.stopImmediatePropagation();
                            e.stopPropagation();
                            alertDemoDisable(e);
                            return false;
                        });
                    }
                }
            });
        }
        
        function alertDemoDisable(ev) {
            $mdDialog.show(
                    $mdDialog.confirm()
                      .clickOutsideToClose(true)
                      .title('您想体验更多数人云功能吗？')
                      .textContent('')
                      .ariaLabel('confirm demo disable')
                      .ok('是，立即使用')
                      .cancel('不，再看看')
                      .targetEvent(ev)
                  ).then(function () {
                      var w = $window.open();
                      w.location = '/auth/register';
                  });
        }
    }
})();