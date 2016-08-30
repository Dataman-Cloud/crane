(function () {
    'use strict';
    angular.module('app.stack')
        .factory('createSelectModal', createSelectModal);

    /* @ngInject */
    function createSelectModal($mdDialog) {

        return {
            open: open
        };

        function open(templateUrl, ev) {

            var dialog = $mdDialog.show({
                controller: FormModalCtrl,
                controllerAs: 'formCtrl',
                templateUrl: templateUrl,
                parent: angular.element(document.body),
                targetEvent: ev,
                clickOutsideToClose: true
            });
            return dialog;
        }

        /* @ngInject */
        function FormModalCtrl($mdDialog, $state) {
            var self = this;

            self.ok = function (state) {
                $state.go(state);
                $mdDialog.hide();
            };

            self.cancel = function () {
                $mdDialog.cancel();
            };

        }
    }
})();