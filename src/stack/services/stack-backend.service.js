(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackBackend', stackBackend);


    /* @ngInject */
    function stackBackend(gHttp) {
        return {
            createStack: createStack,
            listStacks: listStacks
        };
        
        function createStack(data, form) {
            return gHttp.Resource('stack.stacks').post(data, {form: form});
        }
        
        function listStacks() {
            return gHttp.Resource('stack.stacks').get();
        }
    }
})();