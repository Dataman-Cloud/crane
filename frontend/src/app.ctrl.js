(function () {
    'use strict';
    angular.module('app')
        .controller('RootCtrl', RootCtrl);

    /* @ngInject */
    function RootCtrl($state, $window, mdSideNav, gHttp, utils, userBackend, $rootScope, tty, stream, layoutBackend, miscBackend, userCurd, $scope, $translate) {
        var self = this;
        var token = $window.localStorage.getItem('token');

        $rootScope.accountId = null;
        $rootScope.licenseValidFlag = true;

        self.noticeNav = mdSideNav.createSideNav('noticeNav');
        self.simulateQuery = true;
        self.language = $window.localStorage.getItem('language') || 'en';

        self.goBack = goBack;
        self.querySearch = querySearch;
        self.logout = logout;
        self.searchJump = searchJump;
        self.openSearch = openSearch;
        self.regLicense = regLicense;

        activate();

        function activate() {
            ///
            initUser();
            listComponent();
        }

        function initUser() {
            if (token) {
                gHttp.setToken(token);
                tty.setToken(token);
                stream.setToken(token);

            } else {
                utils.redirectLogin(true)
            }

            userBackend.aboutMe()
                .then(function (data) {
                    $rootScope.accountId = data.Id;
                    $rootScope.userName = data.Email;
                })
        }

        function checkLicense() {
            userBackend.checkLicense()
                .then(function (data) {
                    $rootScope.licenseValidFlag = (Date.now() / 1000) < data.License;
                }, function (res) {
                    $rootScope.licenseValidFlag = false;
                })
        }

        function regLicense(ev) {
            userCurd.registerLicense(ev);
        }

        function listComponent() {
            miscBackend.craneConfig()
                .then(function (data) {
                    $rootScope.componentList = data.FeatureFlags;

                    if (data.FeatureFlags.indexOf('license') !== -1) {
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
                        $state.go('node.volumeDetail', {
                            node_id: item.Param.NodeId,
                            volume_name: item.Param.VolumeName
                        });
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

        $scope.$watch(function () {
            return self.language
        }, function (newVal) {
            if (newVal) {
                $window.localStorage.setItem('language', newVal);
                $translate.use(self.language)
            }
        })
    }
})();