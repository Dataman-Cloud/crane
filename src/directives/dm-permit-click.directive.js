(function () {
    'use strict';
    angular.module('glance')
        .directive('dmPermitClick', dmPermitClick);

    dmPermitClick.$inject = ["alertModal"];

    function dmPermitClick(alertModal) {
        return {
            priority: -1,
            restrict: 'A',
            scope: {
                permitRoles: "@dmPermitClick",
                role: "=dmRole",
                errorMsg: "=dmPermitMsg"
            },
            link: link
        }
        
        function link(scope, elem, attrs) {
            elem.on('click', function(e) {
                var permitRoles = scope.permitRoles.split(",");
                permitRoles.push("", "-1");
                if (permitRoles.indexOf(scope.role+"") < 0 ) {
                    if (scope.errorMsg) {
                        alertModal.open(scope.errorMsg, e);
                    } else {
                        alertModal.open("抱歉，您没有权限进行此操作!", e);
                    }
                    e.preventDefault();
                    e.stopImmediatePropagation();
                    e.stopPropagation();
                    return false;
                }
            })
        }
    }
})();