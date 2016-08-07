(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackListCtrl', StackListCtrl);


    /* @ngInject */
    function StackListCtrl(stacks, stackCurd) {
        var self = this;
        
        self.stacks = stacks
    }
})();
