(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('confirmModal', confirmModal);

    /* @ngInject */
    function confirmModal($mdDialog) {
        
        return {
            open: open
        }
        
        function open(title, ev, content) {
            var confirm = $mdDialog.confirm()
            .clickOutsideToClose(true)
            .title(title)
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