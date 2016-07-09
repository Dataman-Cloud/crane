(function () {
    'use strict';
    angular.module('glance.utils')
        .factory('mdTable', mdTable);

    /* @ngInject */
    function mdTable($state, $stateParams) {

        return {createTable: createTable};
        function MdTable(stateName, defSort) {
            this.stateName = stateName;
            this.query = {
                order: initOrder(defSort),
                limit: $stateParams.per_page || 20,
                page: $stateParams.page || 1
            };

            /*
             * do not use prototype for md-table
             */
            this.getPage = function (page, perPage) {
                $state.go(stateName, {
                    page: page,
                    per_page: perPage,
                    order: $stateParams.order,
                    sort_by: $stateParams.sort_by,
                    keywords: $stateParams.keywords
                });
            };
            this.getOrder = function (order) {
                var orderObj = {
                    direction: 'asc',
                    order: order
                };

                if (order.charAt(0) === '-') {
                    orderObj.direction = 'desc';
                    orderObj.order = order.slice(1);
                }

                $state.go(stateName, {
                    per_page: $stateParams.per_page,
                    order: orderObj.direction,
                    sort_by: orderObj.order,
                    keywords: $stateParams.keywords
                });
            };

            this.doSearch = function (searchKey) {
                $state.go(stateName, {
                    page: 1,
                    per_page: $stateParams.per_page,
                    keywords: searchKey
                });
            }
        }

        function initOrder(defSort) {
            if (!$stateParams.order) {
                return defSort
            } else if ($stateParams.order === 'asc') {
                return $stateParams.sort_by
            } else {
                return '-' + $stateParams.sort_by
            }
        }

        function createTable(stateName, defSort) {
            return new MdTable(stateName, defSort)
        }

    }
})();