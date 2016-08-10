(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceCdCtrl', ServiceCdCtrl);

    /* @ngInject */
    function ServiceCdCtrl(serviceCdUrl, $stateParams) {
        var self = this;
        self.serviceCdCommand = "curl -XPUT -H 'Content-Type: application/json' " +  BACKEND_URL_BASE.defaultBase + "api/v1/stacks/" + $stateParams.stack_name + "/services/" + serviceCdUrl + "/rolling_update?image=<your-new-image>";
    }
})();
