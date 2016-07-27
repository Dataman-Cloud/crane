(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateByFormCtrl', StackCreateByFormCtrl);

    /* @ngInject */
    function StackCreateByFormCtrl($state, stackCurd, $stateParams, networkBackend, $scope) {
        var self = this;

        self.form = {
            Namespace: $stateParams.stack_name || '',
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
                        "NanoCPUs": null, //CPU 限制
                        "MemoryBytes": null //内存限制
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
                    ////Replicated/GlobalService 二选一
                    //"Replicated": {
                    //    "Replicas": null //任务数
                    //},
                    "GlobalService": {}
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
                "defaultMode": 'GlobalService'
            }
        ];

        self.serveNameList = [];

        self.addServe = addServe;
        self.removeServe = removeServe;
        self.create = create;
        self.addConfig = addConfig;
        self.deleteConfig = deleteConfig;
        self.modeChange = modeChange;
        self.loadNetworks = loadNetworks;
        self.listNames = listNames;
        self.listEnv = listEnv;
        self.listServeLabel = listServeLabel;
        self.listContainerLabel = listContainerLabel;
        self.listConstraints = listConstraints;

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
                        "NanoCPUs": null, //CPU 限制
                        "MemoryBytes": null //内存限制
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
                    ////Replicated/GlobalService 二选一
                    //"Replicated": {
                    //    "Replicas": null //任务数
                    //},
                    "GlobalService": {}
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
                "defaultMode": 'GlobalService'
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

            stackCurd.createStack(self.form, $scope.staticForm, $stateParams.group_id);
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
                    GlobalService: {}
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
                angular.forEach(item.TaskTemplate.ContainerSpec.Env, function (env, index, array) {
                    array[index] = env.key + '=' + env.value
                });

                angular.forEach(item.TaskTemplate.Placement.Constraints, function (constraint, index, array) {
                    array[index] = constraint.key + '=' + constraint.value
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

        function listEnv(serve) {
            var env = serve.TaskTemplate.ContainerSpec.Env.map(function (item, index) {
                if (item.key) {
                    return item.key
                }
            });

            return env
        }

        function listServeLabel(serve) {
            var serveLabel = serve.Labels.map(function (item, index) {
                if (item.key) {
                    return item.key
                }
            });

            return serveLabel
        }

        function listContainerLabel(serve) {
            var containerLabel = serve.TaskTemplate.ContainerSpec.Labels.map(function (item, index) {
                if (item.key) {
                    return item.key
                }
            });

            return containerLabel
        }

        function listConstraints(serve) {
            var constraints = serve.TaskTemplate.Placement.Constraints.map(function (item, index) {
                if (item.key) {
                    return item.key
                }
            });

            return constraints
        }
    }
})();