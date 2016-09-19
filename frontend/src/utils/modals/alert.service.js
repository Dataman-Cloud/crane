(function () {
    'use strict';
    angular.module('app.utils')
        .factory('alertModal', alertModal);


    /* @ngInject */
    function alertModal($mdDialog, $filter) {
        
        return {
            open: open
        };
        
        function open(title, ev, content) {
            var alert = $mdDialog.alert()
                    .clickOutsideToClose(true)
                    .title(title)
                    .ok($filter('translate')('Confirm'))
                    .targetEvent(ev);
            if (content) {
                alert.htmlContent(content);
            }
            var dialog = $mdDialog.show(alert);
            return dialog;
        }
       
    }

})();