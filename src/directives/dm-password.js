/**
 * Created by my9074 on 16/2/23.
 */
(function () {
    'use strict';
    angular.module('glance')
        .directive('dmPassword', dmPassword);

    function dmPassword() {
        return {
            restrict: 'A',
            require: 'ngModel',
            link: function(scope, ele, attrs, ctrl) {
                var regex = /([A-Za-z0-9\?\,\.\:\;\'\"\!\(\)])*[A-Z]/;
                function valueLength(value) {
                    var length = value.length;
                    return Boolean(length <= 0 || (length >= 8 && length <= 16));
                }

                ctrl.$parsers.unshift(function(value) {
                    var valid = true;
                    if (value) {
                        var len = valueLength(value);
                        var reg = regex.test(value);
                        valid = Boolean(reg && len);
                    }
                    ctrl.$setValidity('dmPassword', valid);
                    return valid ? value : undefined;
                });

                ctrl.$formatters.unshift(function(value) {
                    var valid = true;
                    if (value) {
                        var reg = regex.test(value);
                        var len = valueLength(value);
                        valid = Boolean(reg && len);
                    }

                    ctrl.$setValidity('dmPassword', valid);
                    return value;
                });
            }
        };
    }
})();