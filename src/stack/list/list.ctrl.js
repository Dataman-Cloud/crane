(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackListCtrl', StackListCtrl);


    /* @ngInject */
    function StackListCtrl(stackBackend, utils, stackCurd) {
        var self = this;

        activate();

        function activate() {

        }
    }
})();
