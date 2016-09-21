(function () {
    'use strict';
    angular.module('app.utils')
        .factory('utils', utils);

    /* @ngInject */
    function utils(Notification, $rootScope, $window, $location) {
        return {
            buildFullURL: buildFullURL,
            encodeQueryParams: encodeQueryParams,
            redirectLogin: redirectLogin,
            isEmpty: isEmpty,
            convert2Mapping: convert2Mapping,
            startWith: startWith
        };

        function getUrlTemplate(name, isWS) {
            var confs = name.split('.');
            var categoryKey = confs[0];
            var detailKey = confs[1];
            var base;
            if (BACKEND_URL_BASE[categoryKey]) {
                base = BACKEND_URL_BASE[categoryKey];
            } else {
                base = BACKEND_URL_BASE.defaultBase;
            }
            base = buildBase(base, isWS);
            return urlJoin(base, BACKEND_URL[categoryKey][detailKey]);
        }
        
        function urlJoin(domain, uri) {
            domain = domain.replace(/(\/*$)/g,"");
            uri = uri.replace(/(^\/*)/g,"");
            return domain + "/" + uri;
        }

        function buildBase(base, isWS) {
            if (!base || base === '/') {
                base = $location.protocol() + "://" + $location.host() + ":" + $location.port();
            }
            if (isWS) {
                var bases = base.split(':');
                var protocol = 'ws';
                if (bases[0] === 'https' || bases[0] === 'wss') {
                    protocol = 'wss';
                }
                bases[0] = protocol;
                base = bases.join(':');
            }
            return base;
        }
        
        function buildFullURL(name, params, isWS) {
            var url = getUrlTemplate(name, isWS);
            if (params) {
                for(var key in params){
                    if(params.hasOwnProperty(key))
                        url = url.replace("$" + key, params[key]);
                }
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
            $window.localStorage.removeItem('token');
            var href = HOME_URL + "?timestamp=" + new Date().getTime();
            if (isReturn) {
                href += '&return_to=' + encodeURIComponent($window.location.href);
            }
            $window.location.href = href;
            $rootScope.$destroy();
        }
        
        function isEmpty(value) {
            var empty = false;
            if (!value) {
                empty = true;
            } else if (angular.isArray(value)) {
                if (value.length <= 0) {
                    empty = true;
                }
            } else if (angular.isObject(value)) {
                if (Object.keys(value).length <= 0) {
                    empty = true;
                }
            }
            return empty;
        }
        
        function convert2Mapping(values, key) {
            if (!key) {
                key = 'ID';
            }
            var mapping = {};
            angular.forEach(values, function (value) {
                mapping[value[key]] = value;
            });
            return mapping;
        }
        
        function startWith(str, prefix) {
            return str.substr(0, prefix.length) === prefix;
        }

    }

})();
