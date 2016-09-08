(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateByFormCtrl', StackCreateByFormCtrl);

    /* @ngInject */
    function StackCreateByFormCtrl(stackCurd, networkBackend, $scope, userBackend, FileSaver, $rootScope, registryAuthBackend) {
        var self = this;
        var formTemplate = {
            "Name": "",
            "RegistryAuth": "",
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
                    "StopGracePeriod": null, //杀死容器前等待时间
                    "Args": []
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
                    "Condition": "any", //重启模式
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
                //Replicated/Global 二选一
                "Replicated": {
                    "Replicas": 1 //任务数
                }
                //"Global": {}
            },
            "UpdateConfig": {
                "Parallelism": null,//更新并行任务数
                "Delay": null,//更新延迟
                "FailureAction": 'continue' //更新失败策略 pause/continue
            },
            "Networks": [], //网络
            "EndpointSpec": {
                "Mode": "vip", //vip/dnsrr
                "Ports": [] //端口
            },
            "defaultMode": 'Replicated'
        };

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

        self.serveFormArray = [angular.copy(formTemplate)];

        self.loadGroups = loadGroups;
        self.addServe = addServe;
        self.removeServe = removeServe;
        self.create = create;
        self.createAndDownload = createAndDownload;
        self.addConfig = addConfig;
        self.deleteConfig = deleteConfig;
        self.modeChange = modeChange;
        self.loadNetworks = loadNetworks;
        self.loadRegAuths = loadRegAuths;
        self.listNames = listNames;
        self.listConfigByKey = listConfigByKey;

        function loadGroups() {
            userBackend.listGroup($rootScope.accountId)
                .then(function (data) {
                    self.groups = data
                })
        }

        function addServe() {
            var form = angular.copy(formTemplate);
            self.serveFormArray.push(form);
        }

        function removeServe(index) {
            self.serveFormArray.splice(index, 1);
        }

        function create() {
            ///
            var serveArray = formatServeArray();
            self.form.Stack.Services = {};

            angular.forEach(serveArray, function (serve, index) {
                self.form.Stack.Services[serve.Name] = serve
            });

            return stackCurd.createStack(self.form, $scope.staticForm, self.selectGroupId);
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
                },
                Args: {
                    arg: ''
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
                        Replicas: 1
                    }
                }
            } else {
                serveForm.Mode = {
                    Global: {}
                }
            }
        }

        function loadNetworks() {
            return networkBackend.listNetwork()
                .then(function (data) {
                    self.networks = data
                })
        }

        function loadRegAuths() {
            return registryAuthBackend.listRegAuth()
                .then(function (data) {
                    self.regAuths = data
                })
        }

        function formatServeArray() {
            var serveTempArray = angular.copy(self.serveFormArray);
            var serveLabels = {};
            var containerLabels = {};

            angular.forEach(serveTempArray, function (item, index, array) {

                //Unit conversion
                item.TaskTemplate.Resources.Limits.NanoCPUs = item.TaskTemplate.Resources.Limits.NanoCPUs ? item.TaskTemplate.Resources.Limits.NanoCPUs * Math.pow(10, 9) : null;
                item.TaskTemplate.Resources.Limits.MemoryBytes = item.TaskTemplate.Resources.Limits.MemoryBytes ? item.TaskTemplate.Resources.Limits.MemoryBytes * 1024 * 1024 : null;
                item.TaskTemplate.Resources.Reservations.NanoCPUs = item.TaskTemplate.Resources.Reservations.NanoCPUs ? item.TaskTemplate.Resources.Reservations.NanoCPUs * Math.pow(10, 9) : null;
                item.TaskTemplate.Resources.Reservations.MemoryBytes = item.TaskTemplate.Resources.Reservations.MemoryBytes ? item.TaskTemplate.Resources.Reservations.MemoryBytes * 1024 * 1024 : null;
                item.TaskTemplate.RestartPolicy.Delay = item.TaskTemplate.RestartPolicy.Delay ? item.TaskTemplate.RestartPolicy.Delay * Math.pow(10, 9) : null;
                item.TaskTemplate.RestartPolicy.Window = item.TaskTemplate.RestartPolicy.Window ? item.TaskTemplate.RestartPolicy.Window * Math.pow(10, 9) : null;
                item.TaskTemplate.ContainerSpec.StopGracePeriod = item.TaskTemplate.ContainerSpec.StopGracePeriod ? item.TaskTemplate.ContainerSpec.StopGracePeriod * Math.pow(10, 9) : null;
                item.UpdateConfig.Delay = item.UpdateConfig.Delay ? item.UpdateConfig.Delay * Math.pow(10, 9) : null;

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

                angular.forEach(item.TaskTemplate.ContainerSpec.Args, function (arg, index, array) {
                    array[index] = arg.arg
                });

                item.Labels = serveLabels;
                item.TaskTemplate.ContainerSpec.Labels = containerLabels;

                delete item.defaultMode;
                delete item.showStartupParameter;
                delete item.showTag;
                delete item.showResourceLimit;
                delete item.showFaultTolerant;
                delete item.showSchedulingStrategy;
                delete item.showFileMount;
                delete item.showUpdatePolicy;

                serveLabels = {};
                containerLabels = {}
            });

            return serveTempArray;
        }

        /*
         Check same name
         */
        function listNames(curIndex) {
            var nameList = [];
            angular.forEach(self.serveFormArray, function (item, index) {
                if (curIndex != index) {
                    nameList.push(item.Name)
                }
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

	function createAndDownload() {
            create().then(function(data) {
                var blob = new Blob([angular.toJson(self.form.Stack)], { type: 'text/plain;charset=utf-8' });
                FileSaver.saveAs(blob, self.form.Namespace + '.json');
            })
	}
    }
})();
