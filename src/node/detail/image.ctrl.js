(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeImageCtrl', NodeImageCtrl);

    /* @ngInject */
    function NodeImageCtrl(images) {
        var self = this;

        self.images = images;

    }
})();
