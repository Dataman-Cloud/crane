/*
 * 后台各个服务的地址（含协议，网络地址，端口）.
 * defaultBase为默认地址，如果某个服务的配置为null，则使用它作为后台地址.
 * ws是指websocket.
 */
BACKEND_URL_BASE = {
    defaultBase: "/",
    node: null,
    stack: null,
    network: null,
    misc: null
};

HOME_URL = '/auth';
SAMPLES_URL = '/stack/samples/';
MISC_TOOLS_URL = window.location.protocol + "//" + window.location.host + '/misc-tools/';
DOCKER_REGISTRY_URL = window.location.host + ':5000/';