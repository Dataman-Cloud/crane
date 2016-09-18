(function () {
    'use strict';
    angular.module('app.utils')
        .config(i18nCn);

    /* @ngInject */
    function i18nCn($translateProvider) {
        $translateProvider.translations('zh-cn', {
            //general
            "Create Time": '创建时间',
            "Update Time": '更新时间',
            "Operation": '操作',
            "Update": '更新',

            //index module
            "LOGOUT": '登出',
            "Stack": '项目',
            "Image": '镜像',
            "Facility": '设施',
            "Node": '主机',
            "Network": '网络',
            "Warehouse certification": '仓库认证',
            "Information": '信息',
            "Activation": '激活',

            //stack module
            "Stack List": '项目列表',
            "Create Project": '增加项目',
            "Service Update": '服务更新',
            "Project Detail": '项目详情',
            "Service Detail": '服务详情',
            "Choose the way to create": '选择创建方式',
            "Create of DAB": 'DAB 创建',
            "Create of Form": '向导创建',
            "Create of Shortcut": '快捷创建',
            "Delete Project": '删除项目',
            "Add to template": '添加至模板',
            "Service List": '服务列表',
            "Service Name": '服务名称',
            "Task Number": '任务数',
            "Running/Total": '运行中/总数',
            "Service Scale": '修改任务数',
            "Modal title for delete stack": '项目删除后将无法恢复，确认要删除?',
        });
    }
})();
