(function () {
    'use strict';
    angular.module('app.registry')
        .config(route);

    /* @ngInject */
    function route($stateProvider, $locationProvider, $interpolateProvider) {
        $stateProvider
            .state('registry', {
                url: '/registry',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: '镜像仓库'
                }
            })
            .state('registry.list', {
                url: '/list',
                templateUrl: '/src/registry/list/list.html',
                targetState: 'my',
                ncyBreadcrumb: {
                    label: '镜像列表'
                }
            })
            .state('registry.list.my', {
                url: '/my',
                templateUrl: '/src/registry/list/content.html',
                controller: 'RepositorieListContentCtrl as repositorieListContentCtrl',
                ncyBreadcrumb: {
                    label: '我的镜像'
                },
                resolve: {
                    repositories: listRepositories,
                    type: function () {return 'my';}
                }
            })
            .state('registry.list.public', {
                url: '/public',
                templateUrl: '/src/registry/list/content.html',
                controller: 'RepositorieListContentCtrl as repositorieListContentCtrl',
                ncyBreadcrumb: {
                    label: '公共镜像'
                },
                resolve: {
                    repositories: listRepositories,
                    type: function () {return 'public';}
                }
            })
            
        /* @ngInject */
        function listRepositories(registryBackend) {
            return registryBackend.listRepositories();
        }
            
    }
    
})();
