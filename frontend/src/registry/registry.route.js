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
                controller: 'RegistryCtrl as registryCtrl',
                targetState: 'list',
                ncyBreadcrumb: {
                    label: "{/'Image' | translate/}"
                }
            })
            .state('registry.list', {
                url: '/list?open',
                templateUrl: '/src/registry/list/list.html',
                targetState: 'catalogs',
                ncyBreadcrumb: {
                    label: "{/'Images' | translate/}"
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
                    label: "{/'Stack Deploy' | translate/}"
                },
                resolve: {
                    catalog: getCatalog
                }
            })
            .state('registry.createCatalog', {
                url: '/createCatalog?stack_name',
                templateUrl: '/src/registry/create-update-catalog/create-update.html',
                controller: 'CreateUpdateCatalog as createUpdateCatalog',
                ncyBreadcrumb: {
                    label: "{/'Create Stack Template' | translate/}"
                },
                resolve: {
                    stack: getStack,
                    target: function () {
                        return 'create'
                    }
                }
            })
            .state('registry.updateCatalog', {
                url: '/updateCatalog/:catalog_id',
                templateUrl: '/src/registry/create-update-catalog/create-update.html',
                controller: 'CreateUpdateCatalog as createUpdateCatalog',
                ncyBreadcrumb: {
                    label: "{/'Update Stack Template' | translate/}"
                },
                resolve: {
                    stack: getCatalog,
                    target: function () {
                        return 'update'
                    }
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
                    label: "{/'Image Detail' | translate/}"
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
                    label: "{/'How to create a docker image' | translate/}"
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
            if ($stateParams.stack_name) {
                return stackBackend.getStack($stateParams.stack_name);
            } else {
                return ""
            }
        }
    }

})();
