(function () {
    'use strict';
    angular.module('app.stack')
        .controller('StackCreateByFormCtrl', StackCreateByFormCtrl);

    /* @ngInject */
    function StackCreateByFormCtrl($state, stackBackend, $stateParams) {
        var self = this;

        self.form = {
            Namespace: stackBackend.stack_name || '',
            Stack: {},
            Version: ''
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
                        "StopGracePeriod": "" //杀死容器前等待时间
                    },
                    "Resources": {
                        "NanoCPUs": "", //CPU 限制
                        "MemoryBytes": "" //内存限制
                    },
                    "RestartPolicy": {
                        "Condition": "none", //重启模式
                        "Delay": "", //重启延迟
                        "MaxAttempts": "",//重启最大尝试次数
                        "Window": ""
                    },
                    "Placement": {
                        "Constraints": [] //限制条件
                    },
                    "LogDriver": {}
                },
                "Mode": {
                    //Replicated/GlobalService 二选一
                    "Replicated": {
                        "Replicas": "" //任务数
                    },
                    "GlobalService": {}
                },
                "UpdateConfig": {
                    "Parallelism": "",//更新并行任务数
                    "Delay": "" //更新延迟
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

        function addServe(preIndex) {
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
                        "StopGracePeriod": "" //杀死容器前等待时间
                    },
                    "Resources": {
                        "NanoCPUs": "", //CPU 限制
                        "MemoryBytes": "" //内存限制
                    },
                    "RestartPolicy": {
                        "Condition": "none", //重启模式
                        "Delay": "", //重启延迟
                        "MaxAttempts": "",//重启最大尝试次数
                        "Window": "" //重启间隔
                    },
                    "Placement": {
                        "Constraints": [] //限制条件
                    },
                    "LogDriver": {}
                },
                "Mode": {
                    //Replicated/GlobalService 二选一
                    "Replicated": {
                        "Replicas": "" //任务数
                    },
                    "GlobalService": {}
                },
                "UpdateConfig": {
                    "Parallelism": "",//更新并行任务数
                    "Delay": "" //更新延迟
                },
                "Networks": [], //网络
                "EndpointSpec": {
                    "Mode": "vip", //vip/dnsrr
                    "Ports": [] //端口
                },
                "defaultMode": 'GlobalService'
            };

            self.serveFormArray.push(form);

            //Check same name
            self.serveNameList.push(self.serveFormArray[preIndex].Name);
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

            angular.forEach(serveArray, function (serve, index) {
                self.form.Stack[serve.Name] = serve
            });

            console.log(self.form)
        }

        function addConfig(serveForm, typeName) {

            var config = {
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
                }
            };

            switch (typeName) {
                case 'Labels':
                    serveForm.TaskTemplate.ContainerSpec.Labels.push(config[typeName]);
                    break;
                case 'ServeLabels':
                    serveForm.Labels.push(config[typeName]);
                    break;
                case 'Env':
                    serveForm.TaskTemplate.ContainerSpec.Env.push(config[typeName]);
                    break;
                case 'Mounts':
                    serveForm.TaskTemplate.ContainerSpec.Mounts.push(config[typeName]);
                    break;
                case 'Constraints':
                    serveForm.TaskTemplate.Placement.Constraints.push(config[typeName]);
                    break;
                case 'Ports':
                    serveForm.EndpointSpec.Ports.push(config[typeName]);
                    break;

            }
        }

        function deleteConfig(index, serveForm, typeName) {
            switch (typeName) {
                case 'Labels':
                    serveForm.TaskTemplate.ContainerSpec.Labels.splice(index, 1);
                    break;
                case 'Env':
                    serveForm.TaskTemplate.ContainerSpec.Env.splice(index, 1);
                    break;
                case 'Mounts':
                    serveForm.TaskTemplate.ContainerSpec.Mounts.splice(index, 1);
                    break;
                case 'Constraints':
                    serveForm.TaskTemplate.Placement.Constraints.splice(index, 1);
                    break;
                case 'Ports':
                    serveForm.EndpointSpec.Ports.splice(index, 1);
                    break;

            }
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

        function formatServeArray() {
            var serveTempArray = angular.copy(self.serveFormArray);
            var serveLabels = {};
            var containerLabels = {};

            angular.forEach(serveTempArray, function (item, index, array) {
                angular.forEach(item.TaskTemplate.ContainerSpec.Env, function (env, index, array) {
                    array[index] = env.key + '=' + env.value
                });

                angular.forEach(item.TaskTemplate.Placement.Constraints, function (env, index, array) {
                    array[index] = env.key + '=' + env.value
                });

                angular.forEach(item.Labels, function (env, index, array) {
                    serveLabels[env.key] = env.value
                });

                angular.forEach(item.TaskTemplate.ContainerSpec.Labels, function (env, index, array) {
                    containerLabels[env.key] = env.value
                });

                item.Labels = serveLabels;
                item.TaskTemplate.ContainerSpec.Labels = containerLabels;
                delete item.defaultMode;

            });

            return serveTempArray;
        }
    }
})();
