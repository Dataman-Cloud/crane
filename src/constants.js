MESSAGE_CODE = {
    success: 0,
    dataInvalid: 10001,
    noExist: 10009,
    needActive: 11005,
    needLicence: 11011,
    unknow: 10000
};

STACK_DEFAULT = {
    JsonCompose: '{\n' +
    '  "name": 2048,\n' +
    '  "cpu": 0.1,\n' +
    '  "mem": 64\n' +
    '}\n',
    JsonObj: {
        a: 1,
        b: 2,
        c: 3,
        d: {
            a: 1
        }
    }
}
;

BACKEND_URL = {
    node: {
        nodes: 'api/v1/nodes'
    },
    service: {
        services: 'api/v1/services'
    }
};
