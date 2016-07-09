/*
 * 后台各个服务的地址（含协议，网络地址，端口）.
 * defaultBase为默认地址，如果某个服务的配置为null，则使用它作为后台地址.
 * ws是指websocket.
 */
BACKEND_URL_BASE = {
    defaultBase: "http://10.3.10.131:8000/",
    ws: null,
    auth: null,
    cluster: null,
    metrics: null,
    log: null,
    app: null,
    image: null,
    warning: null,
    billing: null,
    stack: null
};

//部署模式，可选dev，demo，prod.
RUNNING_ENV = "dev"; //dev, demo, prod

//webpage地址（含协议，网络地址，端口）.
USER_URL = "http://localhost:8001";

//是否为线下环境
IS_OFF_LINE = true;  //set true or false

//线下环境图片路径
OFF_LINE_IMAGE_URL = "www.www.com/";

/*
 * agent的配置
 * dmHost为streaming的地址（含协议，网络地址，端口）.
 * installUrl为agent的安装脚本路径
 */
AGENT_CONFIG = {
    dmHost: "DM_HOST=ws://10.3.10.131:8000/",
    installUrl: "AGENT_URL"
};

/*
 * group：共享集群中的应用，映射端口使用的地址
 * demo：demo用户的应用，映射端口使用的地址
 */
APP_CONFIG_SPE_URL = {
    group: "1111.111.11",
    demo: "1111.111.22"
};

//demo用户的邮箱地址
DEMO_EMAIL = "DEMO_USER";

//共享用户权限的域（保存cookies用）。
DOMAIN = '';

// licence是否开启
IS_LICENCE_ON = false; //set true or false;
