(function () {
    'use strict';

    angular.module('app.layout').factory('layoutBackend', layoutBackend);

    /* @ngInject */
    function layoutBackend(gHttp) {

        return {
            globalSearch: globalSearch
        };

        function globalSearch(data) {
            var params = {
                keyword: data
            };
            return gHttp.Resource('layout.search').get({params: params, loading: ''});
        }
    }
})();