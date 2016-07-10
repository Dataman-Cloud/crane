MESSAGE_CODE = {
    success: 0,
    dataInvalid: 10001,
    noExist: 10009,
    needActive: 11005,
    needLicence: 11011,
    unknow: 10000
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
    '    - mysql:mysql\n'
};

BACKEND_URL = {
    node: {
        nodes: 'api/v1/nodes'
    },
    service: {
        services: 'api/v1/services'
    }
};
