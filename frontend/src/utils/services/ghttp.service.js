(function () {
    'use strict';

    angular.module('app.utils').factory('gHttp', gHttp);

    /* @ngInject */
    function gHttp(utils, $q, $rootScope, $http, Notification) {
        var token;

        var ResourceCls = buildResourceCls();

        return {
            setToken: setToken,
            Resource: Resource
        };

        function setToken(val) {
            token = val;
        }

        function Resource(urlName, params) {
            return new ResourceCls(urlName, params);
        }

        function buildResourceCls() {
            function Resource(urlName, params) {
                this.url = utils.buildFullURL(urlName, params);
                this.options = {
                    isAuth: true,
                    loading: 'default',
                    ignoreCodes: [] //忽略错误码对应通知的集合
                }
            }

            Resource.prototype.get = function (options) {
                return this.req('get', options);
            };

            Resource.prototype.post = function (data, options) {
                return this.dataReq('post', data, options);
            };

            Resource.prototype.put = function (data, options) {
                return this.dataReq('put', data, options);
            };

            Resource.prototype.patch = function (data, options) {
                return this.dataReq('patch', data, options);
            };

            Resource.prototype.dataReq = function (method, data, options) {
                if (!options) {
                    options = {};
                }
                options.data = data;
                if (options.form) {
                    options.form.$setPristine();
                    options.form.message_error_info = null;

                    options.ignoreCodes = options.ignoreCodes || [];
                    options.ignoreCodes.push(MESSAGE_CODE.dataInvalid);
                    options.form.$setValidity("submit", false);
                }
                var promise = this.req(method, options);
                if (options.form) {
                    promise.catch(function (data) {
                        if (data.code === MESSAGE_CODE.dataInvalid) {
                            if (data.data) {
                                options.form.message_error_info = data.data;
                            } else {
                                Notification.error(CODE_MESSAGE[data.code]);
                            }
                        }
                        options.form.$setValidity("submit", true);
                    });
                }
                return promise;
            };

            Resource.prototype.delete = function (options) {
                return this.req('delete', options);
            };

            Resource.prototype.req = function (method, options) {
                angular.forEach(options, function (value, key) {
                    if (value !== undefined) {
                        this.options[key] = value;
                    }
                }.bind(this));
                var headers = {
                    'Content-Type': (function () {
                        return this.options.upload ? undefined : 'application/json; charset=UTF-8'
                    }).bind(this)()
                };
                if (this.options.isAuth) {
                    headers["Authorization"] = token;
                }
                var req = {
                    method: method,
                    url: this.url,
                    headers: headers,
                    cache: false,
                    data: this.options.data,
                    params: this.options.params
                };

                if (!this.options.loading) {
                    req.ignoreLoadingBar = true;
                }
                var deferred = $q.defer();
                $http(req).success(function (data) {
                    if (data.code === MESSAGE_CODE.success) {
                        deferred.resolve(data.data);
                    } else {
                        this._handleErrors(status, data, deferred);
                    }
                }.bind(this)).error(function (data, status) {
                    this._handleErrors(status, data, deferred);
                }.bind(this));

                return deferred.promise;

            };

            Resource.prototype._handleErrors = function (status, data, deferred) {
                if (data && angular.isObject(data) && data.code && CODE_MESSAGE[data.code]) {
                    if (SHOW_ERROR_MESSAGE.indexOf(data.code) !== -1) {
                        //notification data.message if SHOW_ERROR_MESSAGE includes code
                        Notification.error(data.message);
                    } else if (this.options.ignoreCodes.indexOf(data.code) === -1) {
                        Notification.error(CODE_MESSAGE[data.code]);
                    }
                } else if (status == 401) {
                    utils.redirectLogin(true);
                } else if (status == 404) {
                    //$state.go('404');
                } else {
                    data = {code: MESSAGE_CODE.unknow};
                    Notification.error(CODE_MESSAGE[data.code]);
                }

                deferred.reject(data)

            };

            return Resource;
        }
    }
})();
