MESSAGE_CODE = {
    success: 0,
    dataInvalid: 10001,
    noExist: 10009,
    needActive: 11005,
    needLicence: 11011,
    unknow: 10000
};

CODE_MESSAGE = {
    10000: "服务忙，请稍后再试",
    10001: "参数错误",
    10008: "参数错误",
    10010: "您没有权限进行此操作",
    10011: "请求参数per_page出错",
    10012: "请求参数page出错",
    11006: "已加入用户组",
    11007: "该用户组有其它用户存在或有集群存在",
    11008: "没有验证手机号",
    11009: "已发送，请等待",
    11010: "验证码错误",
    11011: "授权错误",
    11012: "用户已冻结",
    14001: "应用不存在",
    14002: "应用历史版本id不存在",
    14003: "应用名称冲突",
    14004: "应用端口冲突",
    14005: "应用版本冲突",
    14006: "应用被锁定",
    14007: "应用扩展失败",
    14008: "应用更新已完成，回滚失败",
    14009: "环境变量命名错误",
    14010: "集群不存在",
    14011: "请求错误",
    14012: "应用不存在",
    14013: "应用组件不可用",
    14014: "权限错误",
    14015: "系统保留端口",
    14016: "集群异常",
    14404: "实例不存在",
    14017: "灰度发布已存在",
    15001: "数据库操作错误",
    15002: "参数错误",
    15003: "参数错误",
    15004: "获取路径参数错误",
    15005: "参数错误",
    15006: "类型转换错误",
    15007: "构建entry错误",
    15008: "项目不存在",
    15009: "镜像不存在",
    15010: "drone激活错误",
    15011: "获取公钥错误",
    15012: "构建错误",
    15013: "参数错误",
    15014: "查询构建信息失败",
    15015: "查询日志失败",
    15016: "注册harbor失败",
    15017: "查询最新构建信息失败",
    15018: "更新drone失败",
    15019: "删除drone错误",
    15020: "查询日志失败",
    15021: "仓库地址不合法",
    15022: "拉取代码失败",
    16000: "repo配置文件未找到",
    16001: "应用部署有问题",
    16002: "应用名称冲突",
    16003: "不能删除此应用，有服务存在",
    16004: "问题模板解析错误",
    16005: "Docker Compose解析错误",
    16006: "Marathon Config解析错误",
    16007: "当前状态应用不能跟新",
    16008: "MarathonConfig和DockerCompose不匹配",
    16010: "删除失败",
    16011: "更新失败",
    16012: "获取告警列表失败",
    16020: "参数错误",
    16021: "获取告警事件错误",
    17001: "数据错误",
    17002: "主机状态异常，请修复后再查看",
    17003: "数据错误",
    17004: "数据错误",
    17005: "数据错误",
    17006: "数据错误",
    17007: "数据错误",
    17008: "数据错误",
    17009: "数据错误",
    17010: "数据错误",
    17011: "数据错误",
    17012: "数据错误",
    17013: "数据错误",
    17014: "数据错误",
    17018: "日志告警策略重复",
    18001: "告警策略参数解析失败",
    18002: "应用告警策略重复",
    18010: "删除告警策略失败",
    18011: "查询告警策略失败",
    18012: "查询告警策略失败",
    18020: "告警事件参数解析失败",
    18021: "查询告警事件失败"
};

WS_CODE = {
    token_invalide: 4051
};

FRONTEND_MSG = {
    no_group_admin: "对不起，您不是用户组管理员，不能进行此操作!"
};

APP_STATUS = {
    'undefined': "未知",
    '1': "部署中",
    '2': "运行中",
    '3': "已停止",
    '4': "停止中 ",
    '5': "删除中",
    '6': "扩展中",
    '7': "启动中",
    '8': "撤销中",
    '9': "失联",
    '10': "异常"
};

APP_FAIL_RESULT = {
    0: '正常',
    1: '镜像拉取失败',
    2: '容器启动失败'
};

APP_INS_STATUS = {
    '1': "运行中",
    '2': "部署中"
};

IMAGE_STATUS = {
    'pending': "等待中",
    'running': "执行中",
    'success': "成功",
    'skipped': "已忽略",
    'failure': "失败",
    'killed': "已停止",
    'error': "失败"
};

APP_EVENTS_MSG = {
    ScaleApplication: "应用扩展操作",
    StartApplication: "应用部署操作",
    StopApplication: "应用停止操作",
    TASK_RUNNING: "实例正在运行",
    TASK_FINISHED: "实例运行成功",
    TASK_FAILED: "实例启动失败",
    TASK_KILLED: "实例已被杀死",
    TASK_STAGING: "实例启动中",
    TASK_LOST: "实例已经丢失",
    StartApp: "应用启动",
    StopApp: "应用停止",
    DeployApp: "应用部署",
    UpdateApp: "应用更新",
    UpdateAppNum: "应用扩展",
    CancelScale: "取消应用扩展",
    CancelDeployment: "取消应用部署",
    RestartApplication: "应用重启",
    Redeploy: "应用重新部署",
    PullImageError: "镜像拉取失败",
    UnknownError: "未知错误",
    SlaveRemoved: "主机节点已被删除",
    DockerRunError: "Docker 启动错误"
};

IMAGE_TRIGGER_TYPE = {
    SELECT_TAG: 1,
    SELECT_BRANCH: 2,
    SELECT_ALL: 3
};

APP_EVENTS_TYPE = {
    status_update_event: "实例状态更新",
    deployment_success: "部署成功",
    deployment_failed: "部署失败",
    deployment_step_success: "部署操作成功",
    deployment_step_failure: "部署操作失败",
    AppOperation: "应用操作",
    FailureMessage: "部署失败消息"
};

APP_PROTOCOL_TYPE = [
    {value: 1, name: 'TCP'},
    {value: 2, name: 'HTTP'}
];

SUB_INFOTYPE = {
    nodeStatus: "NodeStatus",
    nodeMetric: "NodeMetric",
    serviceStatus: "ServiceStatus",
    agentUpgradeFailed: "AgentUpgradeFailed",
    clusterStatus: "ClusterStatus"
};

NODE_STATUS = {
    running: "7_running",
    terminated: "0_terminated",
    failed: "2_failed",
    installing: "4_installing",
    initing: "6_initing",
    upgrading: "5_upgrading",
    repairing: "3_repairing",
    abnormal: "1_abnormal"
};

NODE_STATUS_NAME = {
    '7_running': '运行正常',
    '0_terminated': '主机失联',
    '2_failed': '主机异常',
    '6_initing': '主机初始化中',
    '5_upgrading': '主机升级中',
    '1_abnormal': '主机预警',
    '4_installing': '主机安装中',
    '3_repairing': '主机修复中'
};

SERVICE_NAME = {
    master: '主控组件',
    marathon: '应用调度组件',
    zookeeper: 'Zookeeper',
    exhibitor: 'ZK监控组件',
    slave: '节点组件',
    cadvisor: '监控组件',
    logcollection: '日志收集组件',
    bamboo: '服务发现监控组件',
    haproxy: '服务发现代理组件',
    chronos: '定时任务组件',
    docker: 'Docker',
    agent: 'Agent'
};

CLUSTER_STATUS = {
    'new': '新集群',
    'installing': '初始化中',
    'failed': '异常',
    'abnormal': '预警',
    'running': '运行正常',
    'upgrading': '升级中'
};


SERVICES_STATUS = {
    running: 'running',
    installing: 'installing',
    failed: 'failed',
    uninstalled: 'uninstalled',
    uninstalling: "uninstalling",
    pulling: "pulling",
    restarting: "restarting"
};

LOG = {
    logDownloadToplimit: 5000
};

SMS = {
    phoneCodeResendExpire: 60
};

IMAGE_BASE_URL = {
    dev: 'http://devstatic.dataman-inc.net/',
    demo: 'http://demostatic.dataman-inc.net/',
    prod: 'https://static.shurenyun.com/'
};

WARNING_TYPE = {
    MemoryUsed: '内存',
    CpuUsedCores: 'CPU 使用',
    DiskIOReadBytesRate: '磁盘读取',
    DiskIOWriteBytesRate: '磁盘写入',
    NetworkReceviedByteRate: '网络接收',
    NetworkSentByteRate: '网络发送',
    reqrate: '每秒请求数'
};

WARNING_RULE = {
    '>': '大于',
    '==': '等于',
    '<': '小于'
};

WARNING_LEVEL = {
    info: '普通',
    warn: '重要',
    crit: '紧急'
};

STACK_STATUS = {
    pending: '尚未部署',
    deploying: '部署进行',
    deploy_fail: '部署失败',
    running: '运行正常',
    stopped: '应用停止',
    unhealthy: '应用故障'
};

STACK_DEFAULT = {
    DockerCompose: 'mysql:\n' +
    '  image:  catalog.shurenyun.com/library/mysql\n' +
    '  restart: always\n' +
    '  ports:\n' +
    '    - "3306:3306"\n' +
    '  environment:\n' +
    '    MYSQL_ROOT_PASSWORD: foobar\n' +
    'wordpress:\n' +
    '  image:  catalog.shurenyun.com/library/wordpress\n' +
    '  restart: always\n' +
    '  ports:\n' +
    '    - "80:80"\n' +
    '  environment:\n' +
    '    WORDPRESS_DB_HOST: mysql:3306\n' +
    '    WORDPRESS_DB_USER: root\n' +
    '    WORDPRESS_DB_PASSWORD: foobar\n' +
    '  links:\n' +
    '    - mysql:mysql\n',
    SryunCompose: 'mysql:\n' +
    '  cpu: 0.1\n' +
    '  mem: 168\n' +
    '  instances: 1\n' +
    '\n' +
    'wordpress:\n' +
    '  cpu: 0.1\n' +
    '  mem: 168\n' +
    '  instances: 1'
};

USER_STATUS = {
    running: "正常",
    frozen: "冻结"
}

USER_TYPE = {
    true: "管理员",
    false: "成员",
}


BACKEND_URL = {
    auth: {
        auth: "api/v3/auth",
        user: "api/v3/user",
        customerservice: "api/v3/customerservice_url",
        password: 'api/v3/user/password',
        notice: 'api/v3/notice',
        phoneCode: 'api/v3/auth/phone/code',
        verifyPhone: 'api/v3/auth/phone/verification',
        register: 'api/v3/auth/user/registration',
        active: 'api/v3/auth/user/activation/$active_code',
        sendActiveMail: 'api/v3/auth/user/activation',
        forgotPassword: 'api/v3/auth/password/reseturl',
        resetPassword: 'api/v3/auth/password/$reset_code',
    },

    cluster: {
        clusters: "api/v3/clusters",
        versions: "api/v3/clusters/versions",
        cluster: "api/v4/clusters/$cluster_id",
        nodeId: "api/v3/clusters/$cluster_id/new_node_identifier",
        nodes: "api/v4/clusters/$cluster_id/nodes",
        node: "api/v3/clusters/$cluster_id/nodes/$node_id",
        nodeMonitor: "api/v3/clusters/$cluster_id/nodes/$node_id/metrics",
        service: "api/v3/clusters/$cluster_id/nodes/$node_id/services/$service_name",
        labels: "api/v3/labels",
        nodesLabels: "api/v3/clusters/$cluster_id/nodes/labels",
        oldversion: "api/v3/clusters/$cluster_id/oldversion_num"
    },
    metrics: {
        getClusterMonitor: "api/v3/clusters/$cluster_id/metrics",
        appmonit: "api/v3/clusters/$clusterID/apps/$aliase/metrics",
        reqRate: "api/v3/clusters/$cluster_id/apps/$aliase/session",
        monitor: "api/v3/clusters/$clusterID/apps/$aliase/monitor",
        nodeAppList: "api/v3/clusters/$cluster_id/ip/$node_ip/instance"
    },
    ws: {
        subscribe: "streaming/glance/$token"
    },
    log: {
        search: "api/v3/es/index",
        downloadSearch: "api/v3/es/download/index",
        searchContext: "api/v3/es/context",
        downloadContext: "api/v3/es/download/context",
        logPolicy: "api/v3/alarm/$log_id",
        logPolicies: "api/v3/alarm",
        logPolicyEvents: "api/v3/alarms"
    },
    app: {
        userApps: 'api/v3/apps',
        clusterApps: 'api/v3/clusters/$cluster_id/apps',
        clusterAllApps: "api/v3/clusters/$cluster_id/allapps",
        clusterApp: 'api/v3/clusters/$cluster_id/apps/$app_id',
        appEvent: 'api/v3/clusters/$cluster_id/apps/$app_id/events',
        appVersions: 'api/v3/clusters/$cluster_id/apps/$app_id/versions',
        appVersion: 'api/v3/clusters/$cluster_id/apps/$app_id/versions/$version_id',
        appsStatus: "api/v3/apps/status",
        appStatus: "api/v3/clusters/$cluster_id/apps/$app_id/status",
        appTask: "api/v3/clusters/$cluster_id/apps/$app_id/tasks",
        ports: "api/v3/clusters/$cluster_id/ports",
        logPaths: "api/v3/clusters/$cluster_id/apps/$app_id/logpaths",
        appNodes: "api/v3/clusters/$cluster_id/apps/$app_id/appnodes",
        scale: 'api/v3/clusters/$cluster_id/apps/$app_id/scale',
        crons: 'api/v3/crons',
        cron: 'api/v3/crons/$scale_id',
        scaleDetail: 'api/v3/clusters/$cluster_id/apps/$app_id/scale/$scale_id',
        changeWeight: 'api/v3/clusters/$cluster_id/apps/$app_id/weight',
        taskappExtend: 'api/v3/scales',
        canary: "api/v3/clusters/$cluster_id/apps/$app_id/canary/$version_id",
        canarys: "api/v3/clusters/$cluster_id/apps/$app_id/canary",
        canaryStatus: "api/v3/clusters/$cluster_id/apps/$app_id/canary/status"
    },
    user: {
        groups: 'api/v3/groups',
        group: 'api/v3/groups/$group_id',
        groupMemberships: 'api/v3/groups/$group_id/memberships',
        groupMyMemberships: 'api/v3/groups/$group_id/mymemberships',
        groupDemo: 'api/v3/groups/demo/mymemberships',
        users: "api/v3/users",
        user: "api/v3/users/$user_id",
    },

    billing: {
        billings: 'api/v3/billing'
    },

    image: {
        projects: 'api/v3/projects',
        project: 'api/v3/projects/$project_id',
        projectImages: 'api/v3/projects/$project_id/builds',
        projectImage: 'api/v3/projects/$project_id/builds/$build_num/1',
        projectApps: 'api/v3/projects/$project_id/apps',
        deleteImage: 'api/v3/projects/$project_id/images/$image_id',
        imageLog: 'api/v3/projects/$project_id/builds/$build_number/1/logs',
        manualBuild: 'api/v3/projects/$project_id/hook',
        streamLog: 'api/v3/projects/$project_id/builds/$build_number/1/stream'
    },

    repo: {
        repositories: 'api/v3/repositories',
        repository: 'api/v3/repositories/$project_name/$repository_name',
        repositoryTags: 'api/v3/repositories/$project_name/$repository_name/tags',
        repositoryCategories: 'api/v3/repositories/categories',
        deployRepo: 'api/v3/repositories/$project_name/$repository_name/apps'
    },

    warning: {
        tasks: 'api/v3/alert/tasks',
        task: 'api/v3/alert/tasks/$task_id',
        tasksEvent: 'api/v3/alert/events'
    },

    stack: {
        listStack: 'api/v3/stacks',
        stacks: 'api/v3/clusters/$cluster_id/stacks',
        stack: 'api/v3/clusters/$cluster_id/stacks/$stack_id',
        deploy: 'api/v3/clusters/$cluster_id/stacks/$stack_id/deploy',
        deployment: 'api/v3/clusters/$cluster_id/stacks/$stack_id/deployment/$key',
        sse: 'api/v3/stacks/deployment_process'
    },

    licence: {
        licence: 'api/v3/licence'
    }
};
