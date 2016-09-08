/**
 * Created by my9074 on 16/7/15.
 */
(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceConfigCtrl', ServiceConfigCtrl);

    /* @ngInject */
    function ServiceConfigCtrl(service) {
        var self = this;

        self.service = service;
    }
})();
