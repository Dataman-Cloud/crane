(function () {
    'use strict';
    angular.module('app.utils')
        .factory('formModal', formModal);

    /* @ngInject */
    function formModal($mdDialog) {
        
        return {
            open: open
        };

        /*
           dataName: Module 传出的属性名称, 默认为 form
           initData: open Module 时传入的初始值, 默认为 form
           initDataName: open Module 时传入的数据名称
         */
        function open(templateUrl, ev, options) {

            if (!options) {
                options = {};
            }
            if (!options.dataName) {
                options.dataName = 'form';
            }
            if (!options.initDataName) {
                options.initDataName = options.dataName;
            }
            if (!options.ctrlName) {
                options.ctrlName = 'formCtrl';
            }
            
            var dialog = $mdDialog.show({
                controller: FormModalCtrl,
                controllerAs: options.ctrlName,
                templateUrl: templateUrl,
                parent: angular.element(document.body),
                targetEvent: ev,
                clickOutsideToClose:true,
                locals: {dataName: options.dataName, initData: options.initData, initDataName: options.initDataName}
            });
            return dialog;
        }
        
        /* @ngInject */
        function FormModalCtrl($mdDialog, dataName, initData, initDataName) {
            var self = this;
            self[initDataName] = initData;
            
            self.ok = function () {
                $mdDialog.hide(self[dataName]);
            };
            
            self.cancel = function () {
                $mdDialog.cancel();
            };
        
        }
    }
})();