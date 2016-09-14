(function () {
    'use strict';
    angular.module('app.stack')
        .controller('ServiceUpdateCtrl', ServiceUpdateCtrl);

    /* @ngInject */
    function ServiceUpdateCtrl(stackCurd, $stateParams, networkBackend, $scope, service, registryAuthBackend) {
        var self = this;

        self.service = service;

        self.modeChange = modeChange;
        self.addConfig = addConfig;
        self.deleteConfig = deleteConfig;
        self.initSelectNetworks = initSelectNetworks;
        self.listConfigByKey = listConfigByKey;
        self.create = create;

        activate();

        function activate() {
            ///
            self.form = formatServeToForm(service);
            loadNetworks();
            loadRegAuths();
        }

        function formatServeToForm(service) {
            var tempForm = {
                "Name": "",
                "RegistryAuth": "",
                "Labels": {},
                "TaskTemplate": {
                    "ContainerSpec": {
                        "Image": "", //镜像地址
                        "Labels": {},
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
                        "Window": null //重启间隔
                    },
                    "Placement": {
                        "Constraints": [] //限制条件
                    },
                    "LogDriver": {}
                },
                "Mode": {
                    //Replicated/Global 二选一
                    //"Replicated": {
                    //    "Replicas": 1 //任务数
                    //}
                    //"Global": {}
                },
                "UpdateConfig": {
                    "Parallelism": null,//更新并行任务数
                    "Delay": null, //更新延迟
                    "FailureAction": 'continue' //更新失败策略 pause/continue
                },
                "Networks": [], //网络
                "EndpointSpec": {
                    "Mode": "vip", //vip/dnsrr
                    "Ports": [] //端口
                },
                "defaultMode": 'Replicated'
            };

            var form = angular.merge(tempForm, service.Spec);

            form.formLabels = [];
            form.formContainerLabels = [];
            form.formPorts = [];
            form.formMounts = [];
            form.formConstraints = [];
            form.formEnv = [];
            form.formCmd = [];
            form.formArgs = [];
            form.defaultMode = 'Replicated';

            //Unit conversion
            form.TaskTemplate.Resources.Limits.NanoCPUs = form.TaskTemplate.Resources.Limits.NanoCPUs ? form.TaskTemplate.Resources.Limits.NanoCPUs / Math.pow(10, 9) : null;
            form.TaskTemplate.Resources.Limits.MemoryBytes = form.TaskTemplate.Resources.Limits.MemoryBytes ? form.TaskTemplate.Resources.Limits.MemoryBytes / (1024 * 1024) : null;
            form.TaskTemplate.Resources.Reservations.NanoCPUs = form.TaskTemplate.Resources.Reservations.NanoCPUs ? form.TaskTemplate.Resources.Reservations.NanoCPUs / Math.pow(10, 9) : null;
            form.TaskTemplate.Resources.Reservations.MemoryBytes = form.TaskTemplate.Resources.Reservations.MemoryBytes ? form.TaskTemplate.Resources.Reservations.MemoryBytes / (1024 * 1024) : null;
            form.TaskTemplate.RestartPolicy.Delay = form.TaskTemplate.RestartPolicy.Delay ? form.TaskTemplate.RestartPolicy.Delay / Math.pow(10, 9) : null;
            form.TaskTemplate.RestartPolicy.Window = form.TaskTemplate.RestartPolicy.Window ? form.TaskTemplate.RestartPolicy.Window / Math.pow(10, 9) : null;
            form.TaskTemplate.ContainerSpec.StopGracePeriod = form.TaskTemplate.ContainerSpec.StopGracePeriod ? form.TaskTemplate.ContainerSpec.StopGracePeriod / Math.pow(10, 9) : null;
            form.UpdateConfig.Delay = form.UpdateConfig.Delay ? form.UpdateConfig.Delay / Math.pow(10, 9) : null;


            angular.forEach(form.Labels, function (value, key) {
                var obj = {
                    key: key,
                    value: value
                };

                form.formLabels.push(obj)
            });

            angular.forEach(form.TaskTemplate.ContainerSpec.Labels, function (value, key) {
                var obj = {
                    key: key,
                    value: value
                };

                form.formContainerLabels.push(obj)
            });

            if (form.EndpointSpec.Ports) {
                form.formPorts = form.EndpointSpec.Ports
            }

            if (form.TaskTemplate.ContainerSpec.Mounts) {
                angular.forEach(form.TaskTemplate.ContainerSpec.Mounts, function (mount, index) {
                    if (!mount.ReadOnly) mount.ReadOnly = false;
                    form.formMounts.push(mount)
                });
            }

            if (form.TaskTemplate.ContainerSpec.Command) {
                angular.forEach(form.TaskTemplate.ContainerSpec.Command, function (item, index) {
                    var obj = {
                        command: item
                    };

                    form.formCmd.push(obj)
                });
            }

            if (form.TaskTemplate.ContainerSpec.Args) {
                angular.forEach(form.TaskTemplate.ContainerSpec.Args, function (item, index) {
                    var obj = {
                        arg: item
                    };

                    form.formArgs.push(obj)
                });
            }

            if (form.TaskTemplate.Placement.Constraints) {
                angular.forEach(form.TaskTemplate.Placement.Constraints, function (item, index) {
                    var obj = {
                        key: item.slice(0, item.indexOf('==')),
                        value: item.slice(item.indexOf('==') + 1)
                    };

                    form.formConstraints.push(obj)
                });
            }

            if (form.TaskTemplate.ContainerSpec.Env) {
                angular.forEach(form.TaskTemplate.ContainerSpec.Env, function (item, index) {
                    var obj = {
                        key: item.slice(0, item.indexOf('=')),
                        value: item.slice(item.indexOf('=') + 1)
                    };

                    form.formEnv.push(obj)
                });
            }

            self.showStartupParameter = form.TaskTemplate.ContainerSpec.Dir || form.formCmd.length || form.formArgs.length;
            self.showTag = true;
            self.showResourceLimit = form.TaskTemplate.Resources.Limits.NanoCPUs || form.TaskTemplate.Resources.Limits.MemoryBytes || form.TaskTemplate.Resources.Reservations.NanoCPUs || form.TaskTemplate.Resources.Reservations.MemoryBytes;
            self.showFaultTolerant = true;
            self.showUpdatePolicy = true;
            self.showSchedulingStrategy = !!form.formConstraints.length;
            self.showFileMount = !!form.formMounts.length;

            form.defaultMode = Object.keys(form.Mode)[0];
            self.serviceLabelLength = form.formLabels.length;

            return form
        }

        function formatFormToJson() {
            var form = angular.copy(self.form);

            //Unit conversion
            form.TaskTemplate.Resources.Limits.NanoCPUs = form.TaskTemplate.Resources.Limits.NanoCPUs ? form.TaskTemplate.Resources.Limits.NanoCPUs * Math.pow(10, 9) : null;
            form.TaskTemplate.Resources.Limits.MemoryBytes = form.TaskTemplate.Resources.Limits.MemoryBytes ? form.TaskTemplate.Resources.Limits.MemoryBytes * 1024 * 1024 : null;
            form.TaskTemplate.Resources.Reservations.NanoCPUs = form.TaskTemplate.Resources.Reservations.NanoCPUs ? form.TaskTemplate.Resources.Reservations.NanoCPUs * Math.pow(10, 9) : null;
            form.TaskTemplate.Resources.Reservations.MemoryBytes = form.TaskTemplate.Resources.Reservations.MemoryBytes ? form.TaskTemplate.Resources.Reservations.MemoryBytes * 1024 * 1024 : null;
            form.TaskTemplate.RestartPolicy.Delay = form.TaskTemplate.RestartPolicy.Delay ? form.TaskTemplate.RestartPolicy.Delay * Math.pow(10, 9) : null;
            form.TaskTemplate.RestartPolicy.Window = form.TaskTemplate.RestartPolicy.Window ? form.TaskTemplate.RestartPolicy.Window * Math.pow(10, 9) : null;
            form.TaskTemplate.ContainerSpec.StopGracePeriod = form.TaskTemplate.ContainerSpec.StopGracePeriod ? form.TaskTemplate.ContainerSpec.StopGracePeriod * Math.pow(10, 9) : null;
            form.UpdateConfig.Delay = form.UpdateConfig.Delay ? form.UpdateConfig.Delay * Math.pow(10, 9) : null;

            form.TaskTemplate.ContainerSpec.Env = [];
            form.TaskTemplate.Placement.Constraints = [];
            form.Labels = {};
            form.TaskTemplate.ContainerSpec.Labels = {};
            form.TaskTemplate.ContainerSpec.Command = [];
            form.TaskTemplate.ContainerSpec.Args = [];
            form.TaskTemplate.ContainerSpec.Mounts = [];
            form.EndpointSpec.Ports = [];

            if (form.formEnv.length) {
                angular.forEach(self.form.formEnv, function (env, index, array) {
                    form.TaskTemplate.ContainerSpec.Env[index] = env.key + '=' + env.value
                });
            }

            if (form.formConstraints.length) {
                angular.forEach(form.formConstraints, function (constraint, index, array) {
                    form.TaskTemplate.Placement.Constraints[index] = constraint.key + '==' + constraint.value
                });
            }

            if (form.formLabels.length) {
                angular.forEach(form.formLabels, function (label, index, array) {
                    form.Labels[label.key] = label.value
                });
            }

            if (form.formContainerLabels.length) {
                angular.forEach(form.formContainerLabels, function (label, index, array) {
                    form.TaskTemplate.ContainerSpec.Labels[label.key] = label.value
                });
            }

            if (form.formCmd.length) {
                angular.forEach(form.formCmd, function (cmd, index, array) {
                    form.TaskTemplate.ContainerSpec.Command[index] = cmd.command
                });
            }

            if (form.formArgs.length) {
                angular.forEach(form.formArgs, function (arg, index, array) {
                    form.TaskTemplate.ContainerSpec.Args[index] = arg.arg
                });
            }

            form.TaskTemplate.ContainerSpec.Mounts = form.formMounts;
            form.Networks = service.Spec.Networks;
            form.EndpointSpec.Ports = form.formPorts;

            delete form.formEnv;
            delete form.formConstraints;
            delete form.formLabels;
            delete form.formContainerLabels;
            delete form.formCmd;
            delete form.formArgs;
            delete form.formPorts;
            delete form.formNetworks;
            delete form.formMounts;
            delete form.defaultMode;

            return form
        }

        function loadNetworks() {
            return networkBackend.listNetwork()
                .then(function (data) {
                    self.networks = data;
                })
        }

        function loadRegAuths() {
            return registryAuthBackend.listRegAuth()
                .then(function (data) {
                    self.regAuths = data
                })
        }

        function initSelectNetworks(name) {
            if (service.Spec.Networks && service.Spec.Networks.length) {

                return service.Spec.Networks.indexOf(name) !== -1;
            }
        }

        function modeChange() {
            if (self.form.defaultMode === 'Replicated') {
                self.form.Mode = {
                    Replicated: {
                        Replicas: 1
                    }
                }
            } else {
                self.form.Mode = {
                    Global: {}
                }
            }
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

        function create() {
            var json = formatFormToJson();
            stackCurd.updateService(json, $scope.staticForm, $stateParams.stack_name, $stateParams.service_id)
        }


    }
})();