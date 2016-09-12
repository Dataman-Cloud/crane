(function () {
    'use strict';
    angular.module('app.registry')
        .config(route);

    /* @ngInject */
    function route($stateProvider) {
        $stateProvider
            .state('registry', {
                url: '/registry',
                template: '<ui-view/>',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: '镜像'
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
                    skip: true
                },
                resolve: {
                    catalogs: listCatalogs
                }
            })
            .state('registry.catalogDetail', {
                url: '/catalogDetail/:catalog_id',
                templateUrl: '/src/registry/catalog-detail/detail.html',
                controller: 'CatalogDetailCtrl as catalogDetailCtrl',
                ncyBreadcrumb: {
                    label: '项目部署'
                },
                resolve: {
                    catalog: getCatalog
                }
            })
            .state('registry.createCatalog', {
                url: '/createCatalog/:stack_name',
                templateUrl: '/src/registry/create-catalog/create.html',
                controller: 'CreateCatalog as createCatalog',
                ncyBreadcrumb: {
                    label: '创建项目模板'
                },
                resolve: {
                    stack: getStack
                }
            })
            .state('registry.list.public', {
                url: '/public',
                templateUrl: '/src/registry/list/public.html',
                controller: 'RepositorieListPublicCtrl as repositorieListPublicCtrl',
                ncyBreadcrumb: {
                    skip: true
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
                    skip: true
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
                    skip: true
                }
            })
            .state('registry.createNote', {
                url: '/createNote',
                templateUrl: '/src/registry/create-image-note/note.html',
                ncyBreadcrumb: {
                    label: '如何创建镜像'
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
            return registryBackend.getCatalog($stateParams.catalog_id);
        }

        /*@ngInject*/
        function getStack(stackBackend, $stateParams) {
            return stackBackend.getStack($stateParams.stack_name);
        }
    }

})();
