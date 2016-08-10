(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceCdCtrl', ServiceCdCtrl);

    /* @ngInject */
    function ServiceCdCtrl(serviceCdUrl, $stateParams, utils) {
        var self = this;

        var rolling_update_url = utils.buildFullURL('stack.serviceRollingUpdate', {stack_name: $stateParams.stack_name, service_id:  serviceCdUrl});
        self.serviceCdCommand = "curl -XPUT -H 'Content-Type: application/json' " +  rolling_update_url + "?image=<your-new-image>";
    }
})();
