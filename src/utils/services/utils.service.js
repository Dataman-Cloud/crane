(function () {
    'use strict';
    angular.module('app.utils')
        .factory('utils', utils);

    /* @ngInject */
    function utils(Notification, $rootScope, $window, $cookies) {
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
            return base + $rootScope.BACKEND_URL[categoryKey][detailKey];
        }

        function convertProtocol2WS(url) {
            var urls = url.split(':');
            var protocol = 'ws';
            if (urls[0] === 'https' || urls[0] === 'wss') {
                protocol = 'wss';
            }
            urls[0] = protocol;
            return urls.join(':');
        }
        
        function buildFullURL(name, params, convertWS) {
            var url = getUrlTemplate(name);
            if (params) {
                $.each(params, function (key, val) {
                    url = url.replace("$" + key, val);
                });
            }
            if (convertWS) {
                url = convertProtocol2WS(url);
            }
            return url;
        }

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
        }

        function redirectLogin(isReturn) {
            $cookies.remove('token');
            var href = HOME_URL + "?timestamp=" + new Date().getTime();
            if (isReturn) {
                href += '&return_to=' + encodeURIComponent($window.location.href);
            }
            $window.location.href = href;
            $rootScope.$destroy();
        }

    }

})();