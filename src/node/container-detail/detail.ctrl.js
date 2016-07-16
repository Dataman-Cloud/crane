/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeContainerDetailCtrl', NodeContainerDetailCtrl);

    /* @ngInject */
    function NodeContainerDetailCtrl(container) {
        var self = this;

        self.container = container;
    }
})();
