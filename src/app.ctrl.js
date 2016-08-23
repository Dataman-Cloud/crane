(function () {
    'use strict';
    angular.module('app')
        .controller('RootCtrl', RootCtrl);

    /* @ngInject */
    function RootCtrl($state, $window, mdSideNav, gHttp, utils, userBackend, $rootScope, tty, stream, layoutBackend, miscBackend) {
        var self = this;

        $rootScope.accountId = null;
        $rootScope.licenseValidFlag = true;

        self.noticeNav = mdSideNav.createSideNav('noticeNav');
        self.goBack = goBack;
        self.simulateQuery = true;

        self.querySearch = querySearch;
        self.logout = logout;
        self.searchJump = searchJump;
        self.openSearch = openSearch;

        activate();

        function activate() {
            ///
            initUser();
            listComponent();
        }

        function initUser() {
            var token = $window.sessionStorage.getItem('token');
            if (token) {
                gHttp.setToken(token);
                tty.setToken(token);
                stream.setToken(token);

            } else {
                utils.redirectLogin(true)
            }

            userBackend.aboutMe()
                .then(function (data) {
                    $rootScope.accountId = data.Id
                })
        }

        function checkLicense() {
            userBackend.checkLicense()
                .then(function (data) {
                    $rootScope.licenseValidFlag = (Date.now() / 1000) < data.License;
                })
        }

        function listComponent() {
            miscBackend.rolexConfig()
                .then(function (data) {
                    $rootScope.componentList = data.FeatureFlags;

                    if(data.FeatureFlags.indexOf('license') !== -1){
                        checkLicense();
                    }
                });
        }

        function logout() {
            userBackend.logout()
                .then(function () {
                    utils.redirectLogin()
                });
        }

        function goBack(state) {
            if (state) {
                $state.go(state);
            } else {
                $window.history.length > 2 ? $window.history.back() : $state.go('dashboard.home');
            }
        }

        function querySearch(query) {
            return layoutBackend.globalSearch(query)
        }

        function searchJump(item) {
            if (item) {
                //hide search ui
                self.isSearch = false;

                switch (item.Type) {
                    case 'node':
                        //go to node detail
                        $state.go('node.detail', {node_id: item.Param.NodeId});
                        break;
                    case 'network':
                        //go to network detail
                        $state.go('node.networkDetail', {node_id: item.Param.NodeId, network_id: item.Param.NetworkID});
                        break;
                    case 'stack':
                        //go to service list
                        $state.go('stack.detail.service', {stack_name: item.Param.NameSpace});
                        break;
                    case 'service':
                        //go to service detail
                        $state.go('stack.serviceDetail', {
                            stack_name: item.Param.NameSpace,
                            service_id: item.Param.ServiceId
                        });
                        break;
                    case 'task':
                        //go to container detail
                        $state.go('node.containerDetail', {
                            node_id: item.Param.NodeId,
                            container_id: item.Param.ContainerId
                        });
                        break;
                    case 'volume':
                        //go to volume list
                        $state.go('node.detail.volume', {node_id: item.Param.NodeId});
                        break;
                }
            }
        }

        function openSearch() {
            self.isSearch = !self.isSearch;
            self.searchText = '';
            setTimeout(function () {
                document.querySelector('#autoCompleteSearch').focus();
            }, 0);
        }
    }
})();