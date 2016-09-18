(function () {
    'use strict';
    angular.module('app.utils')
        .factory('confirmModal', confirmModal);

    /* @ngInject */
    function confirmModal($mdDialog, $filter) {
        
        return {
            open: open
        };
        
        function open(title, ev, content) {
            var confirm = $mdDialog.confirm()
            .clickOutsideToClose(true)
            .title($filter('translate')(title))
            .targetEvent(ev)
            .ok('确定')
            .cancel('取消');
            if (content) {
                confirm.htmlContent(content);
            }
            var dialog = $mdDialog.show(confirm);
            return dialog
        }
    }

})();