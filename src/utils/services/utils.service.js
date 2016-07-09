(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('utils', utils);

    /* @ngInject */
    function utils(Notification, $rootScope) {
        return {
            buildFullURL: buildFullURL,
            encodeQueryParams: encodeQueryParams,
            redirectLogin: redirectLogin
        };

        function getUrlTemplate(name) {
            var confs = name.split('.');
            var categoryKey = confs[0];
            var detailKey = confs[1];
            var base;
            if (BACKEND_URL_BASE[categoryKey]) {
                base = BACKEND_URL_BASE[categoryKey];
            } else {
                base = BACKEND_URL_BASE.defaultBase;
            }
            return base + BACKEND_URL[categoryKey][detailKey];
        }

        function buildFullURL(name, params) {
            var url = getUrlTemplate(name);
            if (params) {
                $.each(params, function (key, val) {
                    url = url.replace("$" + key, val);
                });
            }
            return url;
        };
        
        function encodeQueryParams($stateParams) {
            var params = {
                    page: $stateParams.page, 
                    per_page: $stateParams.per_page,
                    keywords: $stateParams.keywords
                    };
            if ($stateParams.sort_by) {
                params.sort_by = $stateParams.sort_by;
                params.order = $stateParams.order;
            }
            return params;
        };
        
        function redirectLogin(isReturn) {
            var href = $rootScope.HOME_URL + "?timestamp=" + new Date().getTime(); 
            if (isReturn) {
                href += '&return_to=' + encodeURIComponent(window.location.href);
            }
            window.location.href = href;
            $rootScope.$destroy();
        }

    }

})();