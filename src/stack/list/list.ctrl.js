(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackListCtrl', StackListCtrl);


    /* @ngInject */
    function StackListCtrl(stacks, stackBackend) {
        var self = this;
        
        self.stacks = stacks
    }
})();
