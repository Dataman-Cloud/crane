(function () {
    'use strict';
    angular.module('glance')
        .directive('dynamicVali', dynamicVali);

    dynamicVali.$inject = ["$parse"];

    function dynamicVali($parse) {
        return {
            restrict: "A",
            require: 'ngModel',
            link: function (scope, ele, attrs, ngModelController) {

                ngModelController.$validators.dyRange = function (modelValue, viewValue) {
                    var fn = $parse(attrs['dynamicVali']);
                    var parsedValue = fn(scope);

                    if (ngModelController.$isEmpty(modelValue)) {
                        // consider empty models to be valid
                        return true;
                    }

                    if (rangeCheck(parsedValue, modelValue)) {
                        // it is valid
                        return true;
                    }

                    // it is invalid
                    return false;
                };

                ngModelController.$validators.dyPattern = function (modelValue, viewValue) {
                    var fn = $parse(attrs['dynamicVali']);
                    var parsedValue = fn(scope);

                    if (ngModelController.$isEmpty(modelValue)) {
                        // consider empty models to be valid
                        return true;
                    }

                    if (patternCheck(parsedValue, modelValue)) {
                        // it is valid
                        return true;
                    }

                    // it is invalid
                    return false;
                };

                ngModelController.$validators.dyLength = function (modelValue, viewValue) {
                    var fn = $parse(attrs['dynamicVali']);
                    var parsedValue = fn(scope);

                    if (ngModelController.$isEmpty(modelValue)) {
                        // consider empty models to be valid
                        return true;
                    }

                    if (lengthCheck(parsedValue, modelValue)) {
                        // it is valid
                        return true;
                    }

                    // it is invalid
                    return false;
                };

                ngModelController.$validators.dyEmail = function (modelValue, viewValue) {
                    var fn = $parse(attrs['dynamicVali']);
                    var parsedValue = fn(scope);

                    if (ngModelController.$isEmpty(modelValue)) {
                        // consider empty models to be valid
                        return true;
                    }

                    if (emailCheck(parsedValue, modelValue)) {
                        // it is valid
                        return true;
                    }

                    // it is invalid
                    return false;
                };
            }
        };

        function rangeCheck(parsedValue, modelValue) {
            if (parsedValue && parsedValue.length) {
                for (var i = 0; i < parsedValue.length; i++) {
                    if (parsedValue[i].schema === 'range') {
                        var temporaryArray = parsedValue[i].value.split(',');
                        return temporaryArray.some(function (item, index, array) {
                            var min = parseInt(item.split('-')[0]);
                            var max = parseInt(item.split('-')[1]);

                            return !!(modelValue >= min && modelValue <= max);
                        });
                    }
                }
            }

            return true
        }

        function patternCheck(parsedValue, modelValue) {
            if (parsedValue && parsedValue.length) {
                return parsedValue.every(function (item, index, array) {
                    if (item.schema === 'regexp') {
                        var patt = new RegExp(item.value);
                        return patt.test(modelValue)
                    }
                    return true
                })
            }

            return true
        }

        function lengthCheck(parsedValue, modelValue) {
            if (parsedValue && parsedValue.length) {
                for (var i = 0; i < parsedValue.length; i++) {
                    if (parsedValue[i].schema === 'length') {
                        var minLength = parseInt(parsedValue[i].value.split('-')[0]);
                        var maxLength = parseInt(parsedValue[i].value.split('-')[1]);

                        return (modelValue.length >= minLength && modelValue.length <= maxLength)
                    }
                }
            }

            return true
        }

        function emailCheck(parsedValue, modelValue) {
            if (parsedValue && parsedValue.length) {
                for (var i = 0; i < parsedValue.length; i++) {
                    if (parsedValue[i].schema === 'email') {
                        var EMAIL_REGEXP = /^[a-z0-9!#$%&'*+\/=?^_`{|}~.-]+@[a-z0-9]([a-z0-9-]*[a-z0-9])?(\.[a-z0-9]([a-z0-9-]*[a-z0-9])?)*$/i;

                        return EMAIL_REGEXP.test(modelValue)
                    }
                }
            }

            return true
        }
    }
})();
