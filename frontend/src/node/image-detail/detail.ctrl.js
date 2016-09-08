/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.node')
        .controller('NodeImageDetailCtrl', NodeImageDetailCtrl);

    /* @ngInject */
    function NodeImageDetailCtrl(image) {
        var self = this;

        self.image = image;
    }
})();
