(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateByFormCtrl', StackCreateByFormCtrl);

    /* @ngInject */
    function StackCreateByFormCtrl($state, stackCurd, $stateParams, networkBackend, $scope, userBackend, $rootScope) {
        var self = this;

        self.step = 1;
        self.serveNameList = [];
        self.groups = [];
        self.form = {
            Namespace: '',
            Stack: {
                Version: '',
                Services: {}
            }
        };

        self.serveFormArray = [
            {
                "Name": "",
                "Labels": [],
                "TaskTemplate": {
                    "ContainerSpec": {
                        "Image": "", //镜像地址
                        "Labels": [],//标签,需要转换为 {}
                        "Command": [], //命令行
                        "Env": [],//环境变量
                        "Dir": "",//目录
                        "User": "",
                        "Mounts": [],//挂载
                        "StopGracePeriod": null //杀死容器前等待时间
                    },
                    "Resources": {
                        "Limits": {
                            "NanoCPUs": null, //CPU 限制
                            "MemoryBytes": null //内存限制
                        },
                        "Reservations": {
                            "NanoCPUs": null, //CPU 预留
                            "MemoryBytes": null //内存预留
                        }
                    },
                    "RestartPolicy": {
                        "Condition": "none", //重启模式
                        "Delay": null, //重启延迟
                        "MaxAttempts": null,//重启最大尝试次数
                        "Window": null
                    },
                    "Placement": {
                        "Constraints": [] //限制条件
                    },
                    "LogDriver": {}
                },
                "Mode": {
                    ////Replicated/Global 二选一
                    //"Replicated": {
                    //    "Replicas": null //任务数
                    //},
                    "Global": {}
                },
                "UpdateConfig": {
                    "Parallelism": null,//更新并行任务数
                    "Delay": null //更新延迟
                },
                "Networks": [], //网络
                "EndpointSpec": {
                    "Mode": "vip", //vip/dnsrr
                    "Ports": [] //端口
                },
                "defaultMode": 'Global'
            }
        ];

        self.loadGroups = loadGroups;
        self.addServe = addServe;
        self.removeServe = removeServe;
        self.create = create;
        self.addConfig = addConfig;
        self.deleteConfig = deleteConfig;
        self.modeChange = modeChange;
        self.loadNetworks = loadNetworks;
        self.listNames = listNames;
        self.listConfigByKey = listConfigByKey;

        function loadGroups(){
            userBackend.listGroup($rootScope.accountId)
                .then(function(data){
                    self.groups = data
                })
        }

        function addServe() {
            var form = {
                "Name": "",
                "Labels": [],
                "TaskTemplate": {
                    "ContainerSpec": {
                        "Image": "", //镜像地址
                        "Labels": [],//标签,需要转换为 {}
                        "Command": [], //命令行
                        "Env": [],//环境变量
                        "Dir": "",//目录
                        "User": "",
                        "Mounts": [],//挂载
                        "StopGracePeriod": null //杀死容器前等待时间
                    },
                    "Resources": {
                        "Limits": {
                            "NanoCPUs": null, //CPU 限制
                            "MemoryBytes": null //内存限制
                        },
                        "Reservations": {
                            "NanoCPUs": null, //CPU 预留
                            "MemoryBytes": null //内存预留
                        }
                    },
                    "RestartPolicy": {
                        "Condition": "none", //重启模式
                        "Delay": null, //重启延迟
                        "MaxAttempts": null,//重启最大尝试次数
                        "Window": null //重启间隔
                    },
                    "Placement": {
                        "Constraints": [] //限制条件
                    },
                    "LogDriver": {}
                },
                "Mode": {
                    ////Replicated/Global 二选一
                    //"Replicated": {
                    //    "Replicas": null //任务数
                    //},
                    "Global": {}
                },
                "UpdateConfig": {
                    "Parallelism": null,//更新并行任务数
                    "Delay": null //更新延迟
                },
                "Networks": [], //网络
                "EndpointSpec": {
                    "Mode": "vip", //vip/dnsrr
                    "Ports": [] //端口
                },
                "defaultMode": 'Global'
            };

            self.serveFormArray.push(form);
        }

        function removeServe(index) {
            self.serveFormArray.splice(index, 1);
            if (!self.serveFormArray.length) {
                $state.go('stack.create', {stack_name: $stateParams.stack_name})
            }
        }

        function create() {
            ///
            var serveArray = formatServeArray();
            self.form.Stack.Services = {};

            angular.forEach(serveArray, function (serve, index) {
                self.form.Stack.Services[serve.Name] = serve
            });

            stackCurd.createStack(self.form, $scope.staticForm, self.selectGroupId);
        }

        function addConfig(configs, typeName) {

            var configTemplate = {
                Env: {
                    key: '',
                    value: ''
                },
                Constraints: {
                    key: '',
                    value: ''
                },
                ServeLabels: {
                    key: '',
                    value: ''
                },
                Labels: {
                    key: '',
                    value: ''
                },
                Ports: {
                    Name: '',
                    Protocol: 'tcp',
                    TargetPort: '',
                    PublishedPort: ''
                },
                Mounts: {
                    Source: '',
                    Target: '',
                    ReadOnly: true
                },
                Cmd: {
                    command: ''
                }
            };

            configs.push(configTemplate[typeName]);
        }

        function deleteConfig(index, configs) {
            configs.splice(index, 1);
        }

        function modeChange(serveForm) {
            if (serveForm.defaultMode === 'Replicated') {
                serveForm.Mode = {
                    Replicated: {
                        Replicas: ""
                    }
                }
            } else {
                serveForm.Mode = {
                    Global: {}
                }
            }
        }

        function loadNetworks() {
            networkBackend.listNetwork()
                .then(function (data) {
                    self.networks = data
                })
        }

        function formatServeArray() {
            var serveTempArray = angular.copy(self.serveFormArray);
            var serveLabels = {};
            var containerLabels = {};

            angular.forEach(serveTempArray, function (item, index, array) {

                item.TaskTemplate.Resources.Limits.NanoCPUs = item.TaskTemplate.Resources.Limits.NanoCPUs ? item.TaskTemplate.Resources.Limits.NanoCPUs * Math.pow(10, 9) : null;
                item.TaskTemplate.Resources.Limits.MemoryBytes = item.TaskTemplate.Resources.Limits.MemoryBytes ? item.TaskTemplate.Resources.Limits.MemoryBytes * 1024 * 1024 : null;
                item.TaskTemplate.Resources.Reservations.NanoCPUs = item.TaskTemplate.Resources.Reservations.NanoCPUs ? item.TaskTemplate.Resources.Reservations.NanoCPUs * Math.pow(10, 9) : null;
                item.TaskTemplate.Resources.Reservations.MemoryBytes = item.TaskTemplate.Resources.Reservations.MemoryBytes ? item.TaskTemplate.Resources.Reservations.MemoryBytes * 1024 * 1024 : null;

                angular.forEach(item.TaskTemplate.ContainerSpec.Env, function (env, index, array) {
                    array[index] = env.key + '=' + env.value
                });

                angular.forEach(item.TaskTemplate.Placement.Constraints, function (constraint, index, array) {
                    array[index] = constraint.key + '==' + constraint.value
                });

                angular.forEach(item.Labels, function (label, index, array) {
                    serveLabels[label.key] = label.value
                });

                angular.forEach(item.TaskTemplate.ContainerSpec.Labels, function (label, index, array) {
                    containerLabels[label.key] = label.value
                });

                angular.forEach(item.TaskTemplate.ContainerSpec.Command, function (cmd, index, array) {
                    array[index] = cmd.command
                });

                item.Labels = serveLabels;
                item.TaskTemplate.ContainerSpec.Labels = containerLabels;

                delete item.defaultMode;

                serveLabels = {};
                containerLabels = {}
            });

            return serveTempArray;
        }

        /*
         Check same name
         */
        function listNames() {
            var nameList = [];
            angular.forEach(self.serveFormArray, function (item, index) {
                nameList.push(item.Name)
            });

            return nameList
        }

        function listConfigByKey(config, curIndex) {
            var configs = [];
            if (angular.isArray(config)) {
                configs = config.map(function (item, index) {
                    if (item.key && index != curIndex) {
                        return item.key
                    }
                });
            }

            return configs
        }
    }
})();