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
                url: '/list?open',
                templateUrl: '/src/registry/list/list.html',
                targetState: 'catalogs',
                ncyBreadcrumb: {
                    label: '镜像列表'
                }
            })
            .state('registry.list.catalogs', {
                url: '/catalogs',
                templateUrl: '/src/registry/list/catalog.html',
                controller: 'RepositorieListCatalogCtrl as repositorieListCatalogCtrl',
                ncyBreadcrumb: {
                    label: '应用目录'
                },
                resolve: {
                    catalogs: listCatalogs
                }
            })
            .state('registry.catalogDetail', {
                url: '/catalogDetail/:catalog_name',
                templateUrl: '/src/registry/catalog-detail/detail.html',
                controller: 'CatalogDetailCtrl as catalogDetailCtrl',
                ncyBreadcrumb: {
                    label: '应用部署'
                },
                resolve: {
                    catalog: getCatalog
                }
            })
            .state('registry.list.public', {
                url: '/public',
                templateUrl: '/src/registry/list/public.html',
                controller: 'RepositorieListPublicCtrl as repositorieListPublicCtrl',
                ncyBreadcrumb: {
                    label: '公共镜像'
                },
                resolve: {
                    repositories: listPublicRepositories
                }
            })
            .state('registry.list.mine', {
                url: '/mine',
                templateUrl: '/src/registry/list/mine.html',
                controller: 'RepositorieListMineCtrl as repositorieListMineCtrl',
                ncyBreadcrumb: {
                    label: '我的镜像'
                },
                resolve: {
                    repositories: listMineRepositories
                }
            })
            .state('registry.imageDetail', {
                url: '/imageDetail/:repository/:tag',
                templateUrl: '/src/registry/image-detail/detail.html',
                controller: 'RegistryImageCtrl as registryImageCtrl',
                targetState: 'history',
                ncyBreadcrumb: {
                    label: '镜像详情'
                },
                resolve: {
                    image: getImage
                }
            })
            .state('registry.imageDetail.history', {
                url: '/history',
                templateUrl: '/src/registry/image-detail/history.html',
                controller: 'RegistryImageHistoryCtrl as registryImageHistoryCtrl',
                ncyBreadcrumb: {
                    label: '镜像历史'
                }
            });

        /* @ngInject */
        function listMineRepositories(registryBackend) {
            return registryBackend.listMineRepositories();
        }
            
        /* @ngInject */
        function listPublicRepositories(registryBackend) {
            return registryBackend.listPublicRepositories();
        }
        
            
        /* @ngInject */
        function getImage(registryBackend, $stateParams) {
            return registryBackend.getImage($stateParams.repository, $stateParams.tag);
        }

        /* @ngInject */
        function listCatalogs(registryBackend) {
            return registryBackend.listCatalogs();
        }

        /* @ngInject */
        function getCatalog(registryBackend, $stateParams) {
            return registryBackend.getCatalog($stateParams.catalog_name);
        }
    }
    
})();
